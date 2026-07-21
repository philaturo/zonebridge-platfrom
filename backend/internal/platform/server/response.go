// backend/internal/platform/server/response.go
package server

import (
	"encoding/json"
	"net/http"
)

// JSON encodes the provided data as JSON. It writes to a buffer first
// to ensure headers are only written on successful encoding.
// It returns an error if marshaling fails, allowing the caller to handle it.
func JSON(w http.ResponseWriter, status int, data any) error {
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	
	_, writeErr := w.Write(buf)
	return writeErr
}

// ErrorResponse is a standard envelope for error responses.
type ErrorResponse struct {
	Error string `json:"error"`
}

// Error writes a standardized JSON error response.
func Error(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	
	resp := ErrorResponse{Error: message}
	// We ignore the error here because if this fails, there's nothing more we can do
	_ = json.NewEncoder(w).Encode(resp)
}