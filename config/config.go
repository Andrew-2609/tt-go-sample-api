package config

import (
	"context"
	"fmt"
	"tt-go-sample-api/util"
)

// APIConfig stores the application general
// configuration and environment variables.
type APIConfig struct {
	Environment   string `json:"ENV" mapstructure:"ENV"`
	WebServerPort string `json:"WEB_SERVER_PORT" mapstructure:"WEB_SERVER_PORT"`
	DBDriver      string `json:"DB_DRIVER" mapstructure:"DB_DRIVER"`
	DBName        string `json:"DB_NAME" mapstructure:"DB_NAME"`
	DBUser        string `json:"DB_USER" mapstructure:"DB_USER"`
	DBPassword    string `json:"DB_PASSWORD" mapstructure:"DB_PASSWORD"`
	DBHost        string `json:"DB_HOST" mapstructure:"DB_HOST"`
	DBPort        string `json:"DB_PORT" mapstructure:"DB_PORT"`
	AWSRegion     string `json:"AWS_REGION" mapstructure:"AWS_REGION"`
	AWSEndpoint   string `json:"AWS_ENDPOINT" mapstructure:"AWS_ENDPOINT"`
}

// LoadAPIConfigBasedOnEnvironment will load the application's
// configuration according to the current environment.
//
// Since this is a sample project, I've only implemented a local
// environment configuration. In a work environment, some struct
// could be implemented to load configuration from Cloud (e.g.
// AWS Secrets Manager).
func LoadAPIConfigBasedOnEnvironment(ctx context.Context) (*APIConfig, error) {
	if !util.IsProductionEnv() {
		return NewLocalConfig(util.GetEnvFilepathBasedOnEnvironment()).LoadConfig(ctx)
	}

	return nil, fmt.Errorf("no configuration set for given environment (%s)", util.GetEnv())
}

// GetPostgreSQLSource formats the application's database
// variables to a PostgreSQL source string.
func (cfg *APIConfig) GetPostgreSQLSource() string {
	dbSource := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s",
		cfg.DBDriver,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	if !cfg.IsRunningOnProduction() {
		dbSource += "?sslmode=disable"
	}

	return dbSource
}

// IsRunningOnProduction checks if the application is running
// on production mode (i.e., not locally nor on test environment).
func (cfg *APIConfig) IsRunningOnProduction() bool {
	switch cfg.Environment {
	case "local", "test":
		return false
	default:
		return true
	}
}
