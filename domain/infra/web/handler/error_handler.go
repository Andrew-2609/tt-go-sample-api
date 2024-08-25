package handler

import (
	"errors"
	"fmt"
	"tt-go-sample-api/domain/apierr"
	"tt-go-sample-api/domain/usecase/dto"
	"tt-go-sample-api/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

// handleError centralizes the handling of application errors.
func handleError(ctx *fiber.Ctx, err error) error {
	var apiErr apierr.CustomAPIErrorInterface

	if errors.As(err, &apiErr) {
		return ctx.Status(apiErr.GetStatusCode()).JSON(dto.APIErrorOutputDTO{
			Error: apiErr,
		})
	}

	logger.APILoggerSingleton.Error(ctx.Context(), logger.LogInput{
		Message: "An unmapped error occurred in the application",
		Data:    map[string]any{"error": fmt.Sprintf("%v", err)},
	})

	internalServerError := apierr.NewInternalServerError(apierr.CodeUnknownServerError)

	return ctx.Status(internalServerError.GetStatusCode()).JSON(dto.APIErrorOutputDTO{
		Error: internalServerError,
	})
}
