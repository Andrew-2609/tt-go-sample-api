package repository_test

import (
	"context"
	"os"
	"testing"
	"tt-go-sample-api/config"

	"github.com/stretchr/testify/require"
)

var apiConfigTestSingleton *config.APIConfig

func setup(t *testing.T, ctx context.Context) {
	apiConfig, err := config.NewLocalConfig("../../../../.env.test").LoadConfig(ctx)
	require.NoError(t, err)
	apiConfigTestSingleton = apiConfig
}

func TestMain(m *testing.M) {
	setup(new(testing.T), context.Background())

	code := m.Run()

	os.Exit(code)
}
