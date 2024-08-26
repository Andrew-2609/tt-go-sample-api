package handler

import (
	"net/http"
	"tt-go-sample-api/domain/usecase"
	"tt-go-sample-api/domain/usecase/dto"
	"tt-go-sample-api/pkg/apivalidator"

	"github.com/gofiber/fiber/v2"
)

// RequireEmployeeFromHRWebHandler is a WebHandler for
// requiring employees from the HR team.
type RequireEmployeeFromHRWebHandler struct {
	requireEmployeeFromHRUseCase usecase.RequireEmployeeFromHRUseCaseInterface
}

// NewRequireEmployeeFromHRWebHandler returns a pointer of
// RequireEmployeeFromHRWebHandler with the given use case.
func NewRequireEmployeeFromHRWebHandler(requireEmployeeFromHRUseCase usecase.RequireEmployeeFromHRUseCaseInterface) *RequireEmployeeFromHRWebHandler {
	return &RequireEmployeeFromHRWebHandler{requireEmployeeFromHRUseCase: requireEmployeeFromHRUseCase}
}

// Handle tries to require an employee from the HR, with all
// needed validations.
func (h *RequireEmployeeFromHRWebHandler) Handle(ctx *fiber.Ctx) error {
	var inputDTO dto.RequireEmployeeFromHRInputDTO

	if err := ctx.BodyParser(&inputDTO); err != nil {
		return handleRequestValidationError(ctx, err)
	}

	if err := apivalidator.APIValidatorSingleton.Validate(ctx.Context(), inputDTO); err != nil {
		return handleRequestValidationError(ctx, err)
	}

	outputDTO, err := h.requireEmployeeFromHRUseCase.Execute(ctx.Context(), inputDTO)

	if err != nil {
		return handleError(ctx, err)
	}

	return ctx.Status(http.StatusAccepted).JSON(outputDTO)
}
