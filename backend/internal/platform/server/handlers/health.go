// Package handlers provides HTTP handlers for the ZoneBridge platform.
package handlers

import (
	"net/http"
	"time"

	"github.com/philaturo/zonebridge-platform/internal/platform/clock"
	"github.com/philaturo/zonebridge-platform/internal/platform/server"
	"github.com/philaturo/zonebridge-platform/internal/platform/version"
)

// HealthResponse is the typed structure for the health check endpoint.
// Note: No "data" envelope is used for infrastructure endpoints.
type HealthResponse struct {
	Status    string    `json:"status"`
	Service   string    `json:"service"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
}

// Health handles GET /health requests.
// It is a lightweight process-health check and must not perform database or external dependency checks.
func Health(c clock.Clock, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := HealthResponse{
			Status:    "healthy",
			Service:   "zonebridge",
			Version:   version.Get(),
			Timestamp: c.Now().UTC(),
		}

		server.JSON(w, logger, http.StatusOK, resp)
	}
}