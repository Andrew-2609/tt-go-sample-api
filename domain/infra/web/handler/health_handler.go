package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// HealthWebHandler is a WebHandler to check
// the application's health.
type HealthWebHandler struct {
}

// NewHealthWebHandler returns a pointer of
// HealthWebHandler.
func NewHealthWebHandler() *HealthWebHandler {
	return &HealthWebHandler{}
}

// Handle checks the application's health.
func (h *HealthWebHandler) Handle(ctx *fiber.Ctx) error {
	return ctx.SendStatus(http.StatusOK)
}
