package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/philaturo/zonebridge-platform/internal/platform/clock"
	"github.com/philaturo/zonebridge-platform/internal/platform/server"
)

// mockClock implements clock.Clock for deterministic testing.
type mockClock struct {
	now time.Time
}

func (m *mockClock) Now() time.Time {
	return m.now
}

func TestHealth(t *testing.T) {
	t.Parallel()

	expectedTime := time.Date(2026, 7, 21, 12, 0, 0, 0, time.UTC)
	mc := &mockClock{now: expectedTime}
	logger := slog.Default()

	handler := Health(mc, logger)
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	if rr.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", rr.Header().Get("Content-Type"))
	}

	// Deserialize directly into the typed struct, no generic maps
	var resp server.HealthResponse // Note: HealthResponse is in handlers package, but let's use local type or export it. 
	// Correction: HealthResponse is defined in handlers package. Let's use it directly.
	var healthResp HealthResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &healthResp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if healthResp.Status != "healthy" {
		t.Errorf("expected status 'healthy', got %q", healthResp.Status)
	}
	if healthResp.Service != "zonebridge" {
		t.Errorf("expected service 'zonebridge', got %q", healthResp.Service)
	}
	if healthResp.Version == "" {
		t.Error("expected version to be present")
	}
	if !healthResp.Timestamp.Equal(expectedTime) {
		t.Errorf("expected timestamp %v, got %v", expectedTime, healthResp.Timestamp)
	}
}