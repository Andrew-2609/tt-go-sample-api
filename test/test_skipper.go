package test

import "testing"

// SkipTestIfShortFlagWasPassed checks if the tests
// are running with -short flag. If so, the given
// test will be skipped.
func SkipTestIfShortFlagWasPassed(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test (-short flag provided)")
	}
}
