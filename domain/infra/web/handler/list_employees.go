package handler

import (
	"fmt"
	"net/http"
	"tt-go-sample-api/domain/usecase"
	"tt-go-sample-api/domain/usecase/dto"
	"tt-go-sample-api/pkg/logger"

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
		logger.APILoggerSingleton.Warn(ctx.Context(), logger.LogInput{
			Message: "Request validation error",
			Data:    map[string]any{"validationError": fmt.Sprintf("%v", err)},
		})

		return ctx.
			Status(http.StatusBadRequest).
			JSON(map[string]any{"error": "fields validation error"})
	}

	outputDTO, err := h.listEmployeesUseCase.Execute(ctx.Context(), inputDTO)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(map[string]any{
			"error": fmt.Sprintf("could not list employees: %v", err),
		})
	}

	return ctx.JSON(outputDTO)
}
