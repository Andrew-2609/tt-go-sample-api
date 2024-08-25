package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsZeroValue(t *testing.T) {
	nonZeroValues := []any{"1", 1, true}

	for _, value := range nonZeroValues {
		result := IsZeroValue(value)
		require.False(t, result)
	}

	zeroValues := []any{"", 0, false}

	for _, value := range zeroValues {
		result := IsZeroValue(value)
		require.True(t, result)
	}

	invalidValue := map[string]any{"invalid value": nil}

	result := IsZeroValue(invalidValue["invalid value"])

	require.True(t, result)
}
