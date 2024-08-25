package handler

import (
	"fmt"
	"net/http"
	"tt-go-sample-api/domain/usecase"
	"tt-go-sample-api/domain/usecase/dto"
	"tt-go-sample-api/pkg/apivalidator"

	"github.com/gofiber/fiber/v2"
)

// RegisterEmployeeWebHandler is a WebHandler for
// an employee's registration.
type RegisterEmployeeWebHandler struct {
	registerEmployeeUseCase usecase.RegisterEmployeeUseCaseInterface
}

// NewRegisterEmployeeWebHandler returns a pointer of
// RegisterEmployeeWebHandler with the given use case.
func NewRegisterEmployeeWebHandler(registerEmployeeUseCase usecase.RegisterEmployeeUseCaseInterface) *RegisterEmployeeWebHandler {
	return &RegisterEmployeeWebHandler{registerEmployeeUseCase: registerEmployeeUseCase}
}

// Handle tries to register an employee, with all needed
// validations.
func (h *RegisterEmployeeWebHandler) Handle(ctx *fiber.Ctx) error {
	var inputDTO dto.RegisterEmployeeInputDTO

	if err := ctx.BodyParser(&inputDTO); err != nil {
		return handleRequestValidationError(ctx, err)
	}

	if err := apivalidator.APIValidatorSingleton.Validate(ctx.Context(), inputDTO); err != nil {
		return handleRequestValidationError(ctx, err)
	}

	outputDTO, err := h.registerEmployeeUseCase.Execute(ctx.Context(), inputDTO)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(map[string]any{
			"error": fmt.Sprintf("could not list employees: %v", err),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(outputDTO)
}
