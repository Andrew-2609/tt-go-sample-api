package util

import (
	"os"
	"strings"
)

// GetEnv returns the value of ENV environment variable
// from the Operating System. If empty, it returns "local".
// The returned value is always lowercased.
func GetEnv() string {
	env := os.Getenv("ENV")

	if env == "" {
		return "local"
	}

	return strings.ToLower(env)
}

// GetAPIVersion returns the API Version from the
// API_VERSION environment variable.
//
// If empty, it returns "N/A"
// because this should NEVER happen.
func GetAPIVersion() string {
	apiVersion := os.Getenv("API_VERSION")

	if apiVersion == "" {
		return "N/A"
	}

	return apiVersion
}
