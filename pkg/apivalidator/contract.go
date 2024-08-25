package apivalidator

import "context"

// ValidatorInterface defines an interface for all possible
// custom validators of the application.
type ValidatorInterface interface {
	// Validate must take a context and any data, and return an error
	// if the data didn't passed the expected validations.
	Validate(ctx context.Context, data interface{}) error
}
