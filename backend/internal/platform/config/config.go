// Package config provides configuration loading and validation for the ZoneBridge platform.
package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/philaturo/zonebridge-platform/internal/platform/errors"
)

// Config holds all configuration values for the platform.
type Config struct {
	ServerHost            string
	ServerPort            int
	ServerReadTimeout     time.Duration
	ServerWriteTimeout    time.Duration
	ServerShutdownTimeout time.Duration
	LogLevel              string
	LogFormat             string
	Environment           string
}

// Load reads configuration from environment variables and returns a validated Config.
// It fails fast if an environment variable is present but contains an invalid value.
// Defaults are only applied when an environment variable is entirely absent.
func Load() (*Config, error) {
	cfg := &Config{
		ServerHost: getEnv("ZONEBRIDGE_SERVER_HOST", ""),
		LogLevel:   getEnv("ZONEBRIDGE_LOG_LEVEL", "info"),
		LogFormat:  getEnv("ZONEBRIDGE_LOG_FORMAT", "json"),
		Environment: getEnv("ZONEBRIDGE_ENVIRONMENT", "production"),
	}

	var err error

	cfg.ServerPort, err = getEnvInt("ZONEBRIDGE_SERVER_PORT", 8080)
	if err != nil {
		return nil, fmt.Errorf("load server port: %w", err)
	}

	cfg.ServerReadTimeout, err = getEnvDuration("ZONEBRIDGE_SERVER_READ_TIMEOUT", 10*time.Second)
	if err != nil {
		return nil, fmt.Errorf("load server read timeout: %w", err)
	}

	cfg.ServerWriteTimeout, err = getEnvDuration("ZONEBRIDGE_SERVER_WRITE_TIMEOUT", 10*time.Second)
	if err != nil {
		return nil, fmt.Errorf("load server write timeout: %w", err)
	}

	cfg.ServerShutdownTimeout, err = getEnvDuration("ZONEBRIDGE_SERVER_SHUTDOWN_TIMEOUT", 30*time.Second)
	if err != nil {
		return nil, fmt.Errorf("load server shutdown timeout: %w", err)
	}

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("validate configuration: %w", err)
	}

	return cfg, nil
}

// validate checks that all configuration values are within acceptable ranges.
func (c *Config) validate() error {
	if c.ServerPort < 1 || c.ServerPort > 65535 {
		return fmt.Errorf("%w: server port must be between 1 and 65535, got %d", errors.ErrInvalidConfiguration, c.ServerPort)
	}

	if c.ServerReadTimeout <= 0 {
		return fmt.Errorf("%w: server read timeout must be positive", errors.ErrInvalidConfiguration)
	}

	if c.ServerWriteTimeout <= 0 {
		return fmt.Errorf("%w: server write timeout must be positive", errors.ErrInvalidConfiguration)
	}

	if c.ServerShutdownTimeout <= 0 {
		return fmt.Errorf("%w: server shutdown timeout must be positive", errors.ErrInvalidConfiguration)
	}

	validLogLevels := map[string]bool{"debug": true, "info": true, "warn": true, "error": true}
	if !validLogLevels[c.LogLevel] {
		return fmt.Errorf("%w: log level must be one of [debug, info, warn, error], got %q", errors.ErrInvalidConfiguration, c.LogLevel)
	}

	validLogFormats := map[string]bool{"json": true, "text": true}
	if !validLogFormats[c.LogFormat] {
		return fmt.Errorf("%w: log format must be one of [json, text], got %q", errors.ErrInvalidConfiguration, c.LogFormat)
	}

	validEnvironments := map[string]bool{"development": true, "production": true, "test": true}
	if !validEnvironments[c.Environment] {
		return fmt.Errorf("%w: environment must be one of [development, production, test], got %q", errors.ErrInvalidConfiguration, c.Environment)
	}

	return nil
}

// getEnv reads an environment variable or returns a default value if absent.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvInt reads an integer environment variable.
// It returns an error if the variable exists but is not a valid integer.
// It returns the default value only if the variable is absent.
func getEnvInt(key string, defaultValue int) (int, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue, nil
	}
	
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("environment variable %q must be an integer, got %q", key, value)
	}
	
	return intValue, nil
}

// getEnvDuration reads a duration environment variable.
// It returns an error if the variable exists but is not a valid duration.
// It returns the default value only if the variable is absent.
func getEnvDuration(key string, defaultValue time.Duration) (time.Duration, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue, nil
	}
	
	duration, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf("environment variable %q must be a valid duration, got %q", key, value)
	}
	
	return duration, nil
}