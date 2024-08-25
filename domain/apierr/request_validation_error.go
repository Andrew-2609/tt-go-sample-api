package apierr

import (
	"net/http"
)

// RequestValidationError is a custom API Error used whenever
// a request validation fails, for whatever reason.
//
// It represents a Bad Request (Status Code 400) error.
type RequestValidationError struct {
	APIError
}

// NewRequestValidationError returns a pointer of
// RequestValidationError. The message and code
// don't change.
func NewRequestValidationError() *RequestValidationError {
	return &RequestValidationError{
		APIError: APIError{
			Code:    string(CodeRequestValidationError),
			Message: "Error in fields validation, please make the necessary adjustments and try again.",
		},
	}
}

// GetStatusCode returns 400 (Bad Request) HTTP Status Code.
func (r *RequestValidationError) GetStatusCode() int {
	return http.StatusBadRequest
}
