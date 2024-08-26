package main

import (
	"context"
	"fmt"
	"os"
	"tt-go-sample-api/config"
	"tt-go-sample-api/external/aws/sqs"
	"tt-go-sample-api/external/rdb"
	"tt-go-sample-api/external/rdb/postgresql"
	db "tt-go-sample-api/external/rdb/sqlc"
	"tt-go-sample-api/pkg/logger"
	"tt-go-sample-api/server"
)

func main() {
	mainCtx := context.Background()

	discoverAPIVersion(mainCtx)

	config := setupConfig(mainCtx)
	store := setupRelationalDatabaseConnection(mainCtx, config)
	db.SQLStoreSingleton = store

	logger.APILoggerSingleton = logger.NewWithLogrus("tt-go-sample-api")

	app := server.NewApp(config)
	setupAWSSQS(mainCtx, config)

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

// setupConfig sets up the application's configuration and
// environment variables based on the current running
// environment.
//
// It returns a pointer of config.APIConfig.
func setupConfig(ctx context.Context) *config.APIConfig {
	config, err := config.LoadAPIConfigBasedOnEnvironment(ctx)

	if err != nil {
		logger.NewWithLogrus("tt-go-sample-api").Fatal(ctx, logger.LogInput{
			Message: "Could not set up application configuration",
			Data:    map[string]any{"error": fmt.Sprintf("%v", err)},
		})
	}

	return config
}

// setupRelationalDatabaseConnection configures the application's
// relational database connection and returns its Store.
func setupRelationalDatabaseConnection(ctx context.Context, config *config.APIConfig) db.Store {
	dbDriver, dbName := config.DBDriver, config.DBName

	store, err := postgresql.NewPostgreSQLConnection(config.GetPostgreSQLSource())

	if err != nil {
		logger.APILoggerSingleton.Fatal(ctx, logger.LogInput{
			Message: "Could not establish relational database connection",
			Data:    map[string]any{"error": fmt.Sprintf("%v", err)},
		})
	}

	migrationRunner, err := rdb.NewMigrationRunner().WithMigrate(dbDriver, dbName, store.GetDB())

	if err != nil {
		logger.APILoggerSingleton.Fatal(ctx, logger.LogInput{
			Message: "Could not set application's migration runner",
			Data:    map[string]any{"error": fmt.Sprintf("%v", err)},
		})
	}

	if err = migrationRunner.Run(ctx); err != nil {
		logger.APILoggerSingleton.Fatal(ctx, logger.LogInput{
			Message: "Failed to run database migrations",
			Data:    map[string]any{"error": fmt.Sprintf("%v", err)},
		})
	}

	return store
}

// setupAWSSQS sets up the application's AWS SQS connection.
func setupAWSSQS(ctx context.Context, config *config.APIConfig) {
	apiSQS := sqs.GetAPISQSSingletonSingleton()

	if err := apiSQS.Connect(ctx, config); err != nil {
		logger.APILoggerSingleton.Fatal(ctx, logger.LogInput{
			Message: "Could not establish connection to AWS SQS",
			Data:    map[string]any{"error": fmt.Sprintf("%v", err)},
		})
	}

	logger.APILoggerSingleton.Info(ctx, logger.LogInput{
		Message: "SQS connection successfully established",
	})
}
