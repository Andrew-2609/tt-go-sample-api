package server

import (
	"net/http"
	"tt-go-sample-api/domain/infra/web/webserver"

	"github.com/gofiber/fiber/v2"
)

// setupRoutes sets up all of the API HTTPs routes for
// the given webserver.WebServer.
func setupRoutes(ws *webserver.WebServer) {
	setupBaseRoutes(ws)

	apiRouter := ws.Engine.Group("/api")
	v1 := apiRouter.Group("/v1")

	setupEmployeeManagementRoutes(v1)

	ws.Engine.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusNotFound).JSON(map[string]any{
			"error": "route not found",
		})
	})
}

// setupBaseRoutes sets up routes that are available at
// the base path of the application (e.g. `/health`).
func setupBaseRoutes(ws *webserver.WebServer) {
	// Check API Health
	ws.Engine.Get("/health", func(ctx *fiber.Ctx) error {
		ctx.Status(http.StatusOK)
		return nil
	})
}

// setupEmployeeManagementRoutes sets up all routes regarding
// employees management.
func setupEmployeeManagementRoutes(router fiber.Router) {
	employeesGroup := router.Group("/employees")

	// List Employees
	employeesGroup.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON([]struct{}{})
	})
}
