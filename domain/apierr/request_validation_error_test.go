package apierr

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRequestValidationError(t *testing.T) {
	err := NewRequestValidationError()

	require.NotEmpty(t, err)
	require.IsType(t, &RequestValidationError{}, err)
	require.Equal(t, string(CodeRequestValidationError), err.Code)
}

func TestRequestValidationError_GetStatusCode(t *testing.T) {
	err := NewRequestValidationError()
	errStatusCode := err.GetStatusCode()
	require.Equal(t, http.StatusBadRequest, errStatusCode)
}

func TestRequestValidationError_Error(t *testing.T) {
	err := NewRequestValidationError()
	errMessage := err.Error()
	require.Equal(t, "Error in fields validation, please make the necessary adjustments and try again.", errMessage)
}
