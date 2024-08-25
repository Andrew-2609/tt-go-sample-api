package apivalidator

import "context"

// APIValidatorSingleton is a singleton of *APIValidator. If a
// handler need a request validation, this singleton shall be
// called to perform the necessary validations.
//
// This object is reinitialized in `main.go`, and must not be
// used before that. This is the best way I thought of
// implementing this.
var APIValidatorSingleton *APIValidator = NewAPIValidator(context.Background())
