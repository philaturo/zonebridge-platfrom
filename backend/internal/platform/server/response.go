// Package server provides HTTP transport utilities for the ZoneBridge platform.
package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// JSONResponse is a standard envelope for successful JSON responses.
type JSONResponse struct {
	Data any `json:"data"`
}

// JSON encodes the provided data as JSON and writes it to the ResponseWriter.
// It sets the Content-Type header and handles encoding errors gracefully.
func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(JSONResponse{Data: data}); err != nil {
		// If encoding fails after headers are written, we can only log it.
		// We use slog.Default() here as a fallback, though in practice 
		// this should be injected if possible. For now, it prevents silent failures.
		slog.Error("failed to encode json response", "error", err)
	}
}

// ErrorResponse is a standard envelope for error responses.
type ErrorResponse struct {
	Error string `json:"error"`
}

// Error writes a standardized JSON error response with the given status code and message.
func Error(w http.ResponseWriter, logger *slog.Logger, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := ErrorResponse{Error: message}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Error("failed to encode error response", "error", err, "status", status)
	}
}