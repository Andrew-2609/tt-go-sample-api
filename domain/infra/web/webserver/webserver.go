package webserver

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

// WebServerEngine is an interface for
// an engine that will run underneath a
// WebServer wrapper.

// It defines the base methods necessary
// to run the HTTPs API and allow and
// handle requests.
type WebServerEngine interface {
	Get(path string, handlers ...fiber.Handler) fiber.Router
	Post(path string, handlers ...fiber.Handler) fiber.Router
	Use(args ...interface{}) fiber.Router
	Group(prefix string, handlers ...fiber.Handler) fiber.Router
	Listen(addr string) error
}

// WebServer is a wrapper for a certain
// WebServerEngine, and it's used to
// perform HTTPs requests.
//
// The WebServer must be used as the
// entrypoint of the API, through its
// (blocking) `Start` method.
type WebServer struct {
	Engine WebServerEngine
	Port   string
}

// NewWebServer returns a pointer to a
// WebServer with a given engine and port.
func NewWebServer(engine WebServerEngine, port string) *WebServer {
	return &WebServer{
		Engine: engine,
		Port:   port,
	}
}

// Start runs the application at the
// WebServer port, thus allowing the
// application to receive and handle
// HTTPs requests.
//
// This method will block its routine!
func (ws *WebServer) Start(ctx context.Context) {
	fmt.Printf("Server running at port %s. API Version: %s", ws.Port, "1.0.0")

	if err := ws.Engine.Listen(fmt.Sprintf(":%s", ws.Port)); err != nil {
		log.Fatalf("could not initialize web server: %v", err)
	}
}
