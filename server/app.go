package server

import (
	"tt-go-sample-api/config"
	"tt-go-sample-api/domain/infra/web/webserver"

	"github.com/gofiber/fiber/v2"
)

// NewApp returns a pointer to webserver.WebServer,
// which can be used to perform HTTPs calls.
func NewApp(config *config.APIConfig) *webserver.WebServer {
	ws := webserver.NewWebServer(fiber.New(), config.WebServerPort)

	setupRoutes(ws, config)

	return ws
}
