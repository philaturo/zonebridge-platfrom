// Package version provides application version management.
package version

import (
	"os"
	"strings"
)

// Version is the application version.
// It should be set at build time using ldflags:
//
//	go build -ldflags "-X github.com/philaturo/zonebridge-platform/internal/platform/version.Version=2.0.0"
var Version = "dev"

func init() {
	// If Version was already set via ldflags, do nothing.
	if Version != "dev" {
		return
	}

	// Fallback for local development: attempt to read from the VERSION file.
	// This avoids fragile absolute path assumptions by relying on the 
	// standard practice of running `go run` or the binary from the repository root.
	data, err := os.ReadFile("VERSION")
	if err == nil {
		v := strings.TrimSpace(string(data))
		if v != "" {
			Version = v
		}
	}
}

// Get returns the application version.
func Get() string {
	return Version
}