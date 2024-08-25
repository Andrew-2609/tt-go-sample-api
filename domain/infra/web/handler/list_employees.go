package handler

import (
	"tt-go-sample-api/domain/usecase"
	"tt-go-sample-api/domain/usecase/dto"
	"tt-go-sample-api/pkg/apivalidator"

	"github.com/gofiber/fiber/v2"
)

// ListEmployeesWebHandler is a WebHandler for
// employees listing.
type ListEmployeesWebHandler struct {
	listEmployeesUseCase usecase.ListEmployeesUseCaseInterface
}

// NewListEmployeesWebHandler returns a pointer of
// ListEmployeesWebHandler with the given use case.
func NewListEmployeesWebHandler(listEmployeesUseCase usecase.ListEmployeesUseCaseInterface) *ListEmployeesWebHandler {
	return &ListEmployeesWebHandler{listEmployeesUseCase: listEmployeesUseCase}
}

// Handle tries to retrieve a paginated list of
// employees, with all needed validations.
func (h *ListEmployeesWebHandler) Handle(ctx *fiber.Ctx) error {
	var inputDTO dto.ListEmployeesInputDTO

	if err := ctx.QueryParser(&inputDTO); err != nil {
		return handleRequestValidationError(ctx, err)
	}

	if err := apivalidator.APIValidatorSingleton.Validate(ctx.Context(), inputDTO); err != nil {
		return handleRequestValidationError(ctx, err)
	}

	outputDTO, err := h.listEmployeesUseCase.Execute(ctx.Context(), inputDTO)

	if err != nil {
		return handleError(ctx, err)
	}

	return ctx.JSON(outputDTO)
}
