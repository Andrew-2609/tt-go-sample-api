package handler

import (
	"fmt"
	"net/http"
	"tt-go-sample-api/domain/apierr"
	"tt-go-sample-api/domain/usecase/dto"
	"tt-go-sample-api/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

// handleRequestValidationError centralizes the handling of
// request validation errors.
//
// It logs the error cause and returns a generic response to
// the client for security reasons.
func handleRequestValidationError(ctx *fiber.Ctx, err error) error {
	logger.APILoggerSingleton.Warn(ctx.Context(), logger.LogInput{
		Message: "Request validation error",
		Data:    map[string]any{"error": fmt.Sprintf("%v", err)},
	})

	return ctx.Status(http.StatusBadRequest).JSON(dto.APIErrorOutputDTO{
		Error: apierr.NewRequestValidationError(),
	})
}
