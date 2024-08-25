package server

import (
	"tt-go-sample-api/domain/infra/web/webserver"

	"github.com/gofiber/fiber/v2"
)

// NewApp returns a pointer to webserver.WebServer,
// which can be used to perform HTTPs calls.
func NewApp(port string) *webserver.WebServer {
	ws := webserver.NewWebServer(fiber.New(), port)

	setupRoutes(ws)

	return ws
}
