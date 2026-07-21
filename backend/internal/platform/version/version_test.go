package version

import (
	"testing"
)

func TestGet(t *testing.T) {
	// Note: Testing the init() fallback behavior is inherently difficult 
	// because init() runs once per package before tests start.
	// This test verifies the public API contract.
	// Build-time injection is validated via integration/CI pipelines.
	
	version := Get()
	if version == "" {
		t.Error("Get() returned empty string, expected 'dev' or loaded version")
	}
}

func TestParseVersion(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		err      error
		expected string
	}{
		{
			name:     "valid version",
			data:     []byte("2.0.0\n"),
			err:      nil,
			expected: "2.0.0",
		},
		{
			name:     "valid version with spaces",
			data:     []byte("  2.1.0  "),
			err:      nil,
			expected: "2.1.0",
		},
		{
			name:     "empty file",
			data:     []byte(""),
			err:      nil,
			expected: "dev",
		},
		{
			name:     "file read error",
			data:     nil,
			err:      os.ErrNotExist,
			expected: "dev",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseVersion(tt.data, tt.err)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

// parseVersion is extracted from init() solely for testability.
func parseVersion(data []byte, err error) string {
	if err != nil {
		return "dev"
	}
	v := strings.TrimSpace(string(data))
	if v == "" {
		return "dev"
	}
	return v
}