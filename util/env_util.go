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

// IsLocalEnv checks if the application is running
// locally.
func IsLocalEnv() bool {
	if env := GetEnv(); env == "local" || env == "" {
		return true
	}

	return false
}

// IsTestEnv checks if the application is running
// on testing mode.
func IsTestEnv() bool {
	return GetEnv() == "test"
}

// IsProductionEnv checks if the application is running
// on production mode, that is, the environment is
// neither local nor test.
func IsProductionEnv() bool {
	return !IsLocalEnv() && !IsTestEnv()
}

// GetEnvFilepathBasedOnEnvironment returns the
// filepath of the env file, which can vary
// depending on the environment.
//
//   - When running locally, it'll point to ".env".
//   - When running during tests, it'll point to ".env.test"
func GetEnvFilepathBasedOnEnvironment() string {
	if IsTestEnv() {
		return ".env.test"
	}

	return ".env"
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
