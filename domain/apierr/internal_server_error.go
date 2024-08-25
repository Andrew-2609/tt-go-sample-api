package apierr

import "net/http"

// InternalServerError is a custom API Error used whenever
// the application returns an unknown server error.
//
// It represets an Internal Server Error (Status Code 500).
type InternalServerError struct {
	APIError
}

// NewInternalServerError returns a pointer of
// InternalServerError. The message doesn't change,
// but the code is the one given, since the application
// can fail unexpectedly in many workflows.
func NewInternalServerError[C Code](code C) *InternalServerError {
	return &InternalServerError{APIError{
		Code:    string(code),
		Message: "An unexpected error occurred. Please, try again later or contact our support.",
	}}
}

// GetStatusCode returns 500 (Internal Server Error)
// Status Code.
func (e *InternalServerError) GetStatusCode() int {
	return http.StatusInternalServerError
}
