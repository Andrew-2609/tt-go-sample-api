package apivalidator

import (
	"context"
	"errors"
	"fmt"
	"tt-go-sample-api/pkg/logger"

	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// APIValidator is the custom API validator of HTTPs requests.
// It wraps a pointer of validator.Validate.
//
// It has its own custom validations and implements the
// ValidatorInterface by having a Validate method.
type APIValidator struct {
	validator *validator.Validate
}

// NewAPIValidator returns a pointer of APIValidator with
// the given context.
func NewAPIValidator(ctx context.Context) *APIValidator {
	apiValidator := &APIValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
	apiValidator.loadCustomValidations(ctx)
	return apiValidator
}

// loadCustomValidations loads custom validations into the inner
// validator of APIValidator.
func (v *APIValidator) loadCustomValidations(ctx context.Context) {
	var err error

	errors.Join(err, v.validator.RegisterValidation("notblank", validators.NotBlank))

	if err != nil {
		logger.APILoggerSingleton.Warn(ctx, logger.LogInput{
			Message: "Could not register custom validations to API validator.",
			Data:    map[string]any{"error": fmt.Sprintf("%v", err)},
		})
	}
}

// Validate implements the ValidatorInterface, and takes a context and any
// data as its parameters.
func (v *APIValidator) Validate(ctx context.Context, data interface{}) error {
	err := v.validator.Struct(data)

	if err != nil {
		logger.APILoggerSingleton.Warn(ctx, logger.LogInput{
			Message: "Request fields validation error",
			Data:    map[string]any{"validationError": fmt.Sprintf("%v", err)},
		})

		return fmt.Errorf("error in fields validation, please make the necessary adjustments and try again")
	}

	return nil
}
