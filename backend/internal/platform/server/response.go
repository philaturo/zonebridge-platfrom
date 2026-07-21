// Package server provides HTTP transport utilities for the ZoneBridge platform.
package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// JSON encodes the provided data as JSON and writes it to the ResponseWriter.
// It sets the Content-Type header. If encoding fails, it logs the error using
// the provided logger and attempts to send a 500 Internal Server Error.
func JSON(w http.ResponseWriter, logger *slog.Logger, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Error("failed to encode json response", "error", err)
		// Attempt to send a generic error response since headers are already written
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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