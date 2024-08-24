package config

import (
	"context"

	"github.com/spf13/viper"
)

// LocalConfig is the configuration loader based on
// local environment (i.e. a developer's machine).
type LocalConfig struct {
	envFilepath string
}

// NewLocalConfig returns a pointer of LocalConfig.
func NewLocalConfig(envFilepath string) *LocalConfig {
	return &LocalConfig{envFilepath: envFilepath}
}

// LoadConfig loads the API configuration from local .env file.
func (c *LocalConfig) LoadConfig(ctx context.Context) (*APIConfig, error) {
	var apiConfig APIConfig

	viper.AddConfigPath(".")
	viper.SetConfigName(c.envFilepath)
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&apiConfig); err != nil {
		return nil, err
	}

	return &apiConfig, nil
}
