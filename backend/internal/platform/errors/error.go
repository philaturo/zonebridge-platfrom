// Package errors provides sentinel errors for platform-level concerns.
package errors

import "errors"

// Platform-level sentinel errors.
var (
	// ErrInvalidConfiguration indicates that the provided configuration is invalid.
	ErrInvalidConfiguration = errors.New("invalid configuration")

	// ErrVersionNotFound indicates that the application version could not be determined.
	ErrVersionNotFound = errors.New("version information not found")
)