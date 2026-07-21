// Package middleware provides HTTP middleware for the ZoneBridge platform.
package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log/slog"
	"net/http"
	"time"

	chimiddleware "github.com/go-chi/chi/v5/middleware"
	platformserver "github.com/philaturo/zonebridge-platform/internal/platform/server"
)

// ctxKey is an unexported type for context keys to prevent collisions.
type ctxKey int

const (
	requestIDKey ctxKey = iota
	realIPKey
)

// RequestID generates a cryptographically secure 32-character hex string
// and injects it into the request context and response headers.
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := generateRequestID()
		w.Header().Set("X-Request-Id", id)
		ctx := context.WithValue(r.Context(), requestIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func generateRequestID() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic(fmt.Errorf("critical failure: cannot generate secure request ID: %w", err))
	}
	return hex.EncodeToString(b)
}

// RealIP extracts the true client IP from X-Forwarded-For or X-Real-IP headers,
// falling back to r.RemoteAddr, and injects it into the request context.
func RealIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = r.Header.Get("X-Real-IP")
		}
		if ip == "" {
			ip = r.RemoteAddr
		}
		ctx := context.WithValue(r.Context(), realIPKey, ip)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Logger logs structured request information using the provided slog.Logger.
func Logger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Wrap the ResponseWriter to capture the status code
			ww := chimiddleware.NewWrapResponseWriter(w, r.ProtoMajor)

			next.ServeHTTP(ww, r)

			duration := time.Since(start)
			reqID, _ := r.Context().Value(requestIDKey).(string)
			realIP, _ := r.Context().Value(realIPKey).(string)

			logger.Info("http request",
				"method", r.Method,
				"path", r.URL.Path,
				"status", ww.Status(),
				"duration_ms", duration.Milliseconds(),
				"request_id", reqID,
				"remote_ip", realIP,
				"user_agent", r.UserAgent(),
			)
		})
	}
}

// Recoverer catches panics, logs them with full context using the platform logger,
// and returns a standardized 500 Internal Server Error response.
func Recoverer(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					reqID, _ := r.Context().Value(requestIDKey).(string)

					logger.Error("panic recovered",
						"request_id", reqID,
						"method", r.Method,
						"path", r.URL.Path,
						"panic_value", err,
					)

					// Use standardized error response helper
					platformserver.Error(w, logger, http.StatusInternalServerError, "Internal Server Error")
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
