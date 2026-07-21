package config

import (
	"testing"
	"time"

	"github.com/philaturo/zonebridge-platform/internal/platform/errors"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name        string
		envVars     map[string]string
		wantErr     bool
		errContains string
		validate    func(t *testing.T, cfg *Config)
	}{
		{
			name:    "default values when env vars are absent",
			envVars: map[string]string{},
			wantErr: false,
			validate: func(t *testing.T, cfg *Config) {
				if cfg.ServerPort != 8080 {
					t.Errorf("expected default port 8080, got %d", cfg.ServerPort)
				}
				if cfg.LogLevel != "info" {
					t.Errorf("expected default log level 'info', got %q", cfg.LogLevel)
				}
				if cfg.ServerReadTimeout != 10*time.Second {
					t.Errorf("expected default read timeout 10s, got %v", cfg.ServerReadTimeout)
				}
			},
		},
		{
			name: "custom valid values",
			envVars: map[string]string{
				"ZONEBRIDGE_SERVER_PORT":             "9090",
				"ZONEBRIDGE_LOG_LEVEL":               "debug",
				"ZONEBRIDGE_LOG_FORMAT":              "text",
				"ZONEBRIDGE_ENVIRONMENT":             "development",
				"ZONEBRIDGE_SERVER_READ_TIMEOUT":     "5s",
				"ZONEBRIDGE_SERVER_WRITE_TIMEOUT":    "15s",
				"ZONEBRIDGE_SERVER_SHUTDOWN_TIMEOUT": "60s",
			},
			wantErr: false,
			validate: func(t *testing.T, cfg *Config) {
				if cfg.ServerPort != 9090 {
					t.Errorf("expected port 9090, got %d", cfg.ServerPort)
				}
				if cfg.LogLevel != "debug" {
					t.Errorf("expected log level 'debug', got %q", cfg.LogLevel)
				}
				if cfg.ServerReadTimeout != 5*time.Second {
					t.Errorf("expected read timeout 5s, got %v", cfg.ServerReadTimeout)
				}
			},
		},
		{
			name: "fail fast on malformed port",
			envVars: map[string]string{
				"ZONEBRIDGE_SERVER_PORT": "abc",
			},
			wantErr:     true,
			errContains: "must be an integer",
		},
		{
			name: "fail fast on malformed duration",
			envVars: map[string]string{
				"ZONEBRIDGE_SERVER_READ_TIMEOUT": "invalid",
			},
			wantErr:     true,
			errContains: "must be a valid duration",
		},
		{
			name: "fail on invalid port range (zero)",
			envVars: map[string]string{
				"ZONEBRIDGE_SERVER_PORT": "0",
			},
			wantErr:     true,
			errContains: "server port must be between 1 and 65535",
		},
		{
			name: "fail on invalid log level",
			envVars: map[string]string{
				"ZONEBRIDGE_LOG_LEVEL": "verbose",
			},
			wantErr:     true,
			errContains: "log level must be one of",
		},
		{
			name: "fail on invalid environment",
			envVars: map[string]string{
				"ZONEBRIDGE_ENVIRONMENT": "staging",
			},
			wantErr:     true,
			errContains: "environment must be one of",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variables for this test
			for key, value := range tt.envVars {
				t.Setenv(key, value)
			}

			cfg, err := Load()

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				if tt.errContains != "" && !containsString(err.Error(), tt.errContains) {
					t.Errorf("expected error to contain %q, got %q", tt.errContains, err.Error())
				}
				// Ensure error is properly wrapped
				if !isInvalidConfigError(err) {
					t.Errorf("expected error to wrap ErrInvalidConfiguration, got %v", err)
			 }
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tt.validate != nil {
				tt.validate(t, cfg)
			}
		})
	}
}

func containsString(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func isInvalidConfigError(err error) bool {
	for err != nil {
		if err == errors.ErrInvalidConfiguration {
			return true
		}
		err = err.(interface{ Unwrap() error }).Unwrap()
	}
	return false
}