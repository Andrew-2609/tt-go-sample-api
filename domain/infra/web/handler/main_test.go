package handler_test

import (
	"context"
	"os"
	"testing"
	"tt-go-sample-api/config"
	"tt-go-sample-api/server"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

var (
	apiApp                 *fiber.App
	apiConfigTestSingleton *config.APIConfig
)

func setupRouter(t *testing.T, ctx context.Context) {
	apiConfig, err := config.NewLocalConfig("../../../../.env.test").LoadConfig(ctx)
	require.NoError(t, err)

	webServer := server.NewApp(apiConfig)

	fiberApp, ok := webServer.Engine.(*fiber.App)

	if !ok {
		panic("could not convert webserver engine to *fiber.App")
	}

	apiApp = fiberApp
	apiConfigTestSingleton = apiConfig
}

func TestMain(m *testing.M) {
	setupRouter(new(testing.T), context.Background())

	code := m.Run()

	os.Exit(code)
}
