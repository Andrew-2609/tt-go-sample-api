package server

import (
	"net/http"
	"tt-go-sample-api/config"
	"tt-go-sample-api/domain/infra/postgresql/repository"
	"tt-go-sample-api/domain/infra/web/handler"
	"tt-go-sample-api/domain/infra/web/webserver"
	"tt-go-sample-api/domain/usecase"
	"tt-go-sample-api/external/messaging"

	"github.com/gofiber/fiber/v2"
)

// setupRoutes sets up all of the API HTTPs routes for
// the given webserver.WebServer.
func setupRoutes(ws *webserver.WebServer, config *config.APIConfig) {
	setupBaseRoutes(ws)

	apiRouter := ws.Engine.Group("/api")
	v1 := apiRouter.Group("/v1")

	setupEmployeeManagementRoutes(v1, config)

	ws.Engine.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusNotFound).JSON(map[string]any{
			"error": "route not found",
		})
	})
}

// setupBaseRoutes sets up routes that are available at
// the base path of the application (e.g. `/health`).
func setupBaseRoutes(ws *webserver.WebServer) {
	// Check API Health (may be used by Kubernetes).
	ws.Engine.Get("/health", handler.NewHealthWebHandler().Handle)
}

// setupEmployeeManagementRoutes sets up all routes regarding
// employees management.
func setupEmployeeManagementRoutes(router fiber.Router, config *config.APIConfig) {
	employeesGroup := router.Group("/employees")

	// List Employees
	getEmployeesRepository := repository.NewGetEmployeesPostgreSQLRepository()
	listEmployeesUseCase := usecase.NewListEmployeesUseCase(getEmployeesRepository)

	employeesGroup.Get("/", handler.NewListEmployeesWebHandler(listEmployeesUseCase).Handle)

	// Register Employees
	writeEmployeeRepository := repository.NewWriteEmployeePostgreSQLRepository()
	registerEmployeeUseCase := usecase.NewRegisterEmployeeUseCase(writeEmployeeRepository)

	employeesGroup.Post("/", handler.NewRegisterEmployeeWebHandler(registerEmployeeUseCase).Handle)

	// Require Employees from HR
	messageProducer := messaging.NewAWSSQSProducer(config.AWSHRQueueURL)
	requireEmployeeFromHRUseCase := usecase.NewRequireEmployeeFromHRUseCase(messageProducer)

	employeesGroup.Post("/hr", handler.NewRequireEmployeeFromHRWebHandler(requireEmployeeFromHRUseCase).Handle)
}
