package handler_test

import (
	"net/http"
	"testing"
	"tt-go-sample-api/test"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type HealthWebHandlerTestSuite struct {
	suite.Suite
	Method   string
	BasePath string
}

func (suite *HealthWebHandlerTestSuite) SetupSuite() {
	suite.Method = http.MethodGet
	suite.BasePath = "/health"
}

func (suite *HealthWebHandlerTestSuite) TestHandle() {
	req, err := test.NewRequest(suite.Method, suite.BasePath, nil)
	require.NoError(suite.T(), err)

	resp, err := apiApp.Test(req)
	require.NoError(suite.T(), err)

	require.Equal(suite.T(), http.StatusOK, resp.StatusCode)
}

func TestHealthWebHandlerTestSuiteIntegration(t *testing.T) {
	test.SkipTestIfShortFlagWasPassed(t)
	suite.Run(t, new(HealthWebHandlerTestSuite))
}
