package main

import (
	"context"
	"fmt"
	"os"
	"tt-go-sample-api/pkg/logger"
	"tt-go-sample-api/server"
)

func main() {
	mainCtx := context.Background()

	discoverAPIVersion(mainCtx)
	logger.APILoggerSingleton = logger.NewWithLogrus("tt-go-sample-api")

	app := server.NewApp()

	go app.Start(mainCtx)

	select {}
}

// discoverAPIVersion tries to read the VERSION
// file at the project's root. The value from
// VERSION will be set in a "API_VERSION"
// environment variable.
//
// If any part of this process goes wrong, a log
// entry will be created.
func discoverAPIVersion(ctx context.Context) {
	versionFile, err := os.ReadFile("VERSION")

	if err != nil {
		logger.APILoggerSingleton.Fatal(ctx, logger.LogInput{
			Message: "Could not read project version in VERSION file",
			Data:    map[string]any{"error": fmt.Sprintf("%v", err)},
		})
	}

	if err := os.Setenv("API_VERSION", string(versionFile)); err != nil {
		logger.APILoggerSingleton.Fatal(ctx, logger.LogInput{
			Message: "Could not set API VERSION in environment variables",
			Data:    map[string]any{"error": fmt.Sprintf("%v", err)},
		})
	}
}
