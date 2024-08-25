package apierr

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewInternalServerError(t *testing.T) {
	code := CodeUnknownServerError
	err := NewInternalServerError(code)

	require.NotEmpty(t, err)
	require.IsType(t, &InternalServerError{}, err)
	require.Equal(t, string(code), err.Code)
}

func TestInternalServerError_GetStatusCode(t *testing.T) {
	err := NewInternalServerError(CodeUnknownServerError)
	errStatusCode := err.GetStatusCode()
	require.Equal(t, http.StatusInternalServerError, errStatusCode)
}

func TestInternalServerError_Error(t *testing.T) {
	err := NewInternalServerError(CodeUnknownServerError)
	errMessage := err.Error()
	require.Equal(t, "An unexpected error occurred. Please, try again later or contact our support.", errMessage)
}
