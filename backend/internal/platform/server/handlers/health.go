// backend/internal/platform/server/handlers/health.go
package handlers

import (
	"net/http"

	"github.com/philaturo/zonebridge-platform/internal/platform/clock"
	"github.com/philaturo/zonebridge-platform/internal/platform/response"
	"github.com/philaturo/zonebridge-platform/internal/platform/version"
)

type HealthResponse struct {
	Status    string `json:"status"`
	Service   string `json:"service"`
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"` // Using string for clean JSON serialization
}

func Health(c clock.Clock) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := HealthResponse{
			Status:    "healthy",
			Service:   "zonebridge",
			Version:   version.Get(),
			Timestamp: c.Now().UTC().Format(http.TimeFormat),
		}

		if err := server.JSON(w, http.StatusOK, resp); err != nil {
			// Fallback to standard library if our helper fails catastrophically
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}