package apierr

// CustomAPIErrorInterface defines an interface that must
// be implemented by any custom error within the application.
//
// It has methods useful for HTTPs handling and logging.
type CustomAPIErrorInterface interface {
	// Error returns the Custom Error's message.
	Error() string

	// GetStatusCode returns the HTTP Status Code of the
	// Custom Error.
	GetStatusCode() int
}

// APIError implements the CustomAPIErrorInterface, and
// has an inner unique code to identify a custom error
// exclusively.
type APIError struct {
	// Code is a unique identifier for an error.
	Code string `json:"code"`

	// Message describes why, when, and/or how the error happened.
	Message string `json:"message"`
}

// Error returns the APIError's inner message.
func (a *APIError) Error() string {
	return a.Message
}
