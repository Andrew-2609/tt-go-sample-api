package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetEnv(t *testing.T) {
	testCases := map[string]struct {
		env            string
		validateResult func(t *testing.T, env string)
	}{
		"Empty": {
			env: "",
			validateResult: func(t *testing.T, env string) {
				require.Equal(t, "local", env)
			},
		},
		"Local": {
			env: "local",
			validateResult: func(t *testing.T, env string) {
				require.Equal(t, "local", env)
			},
		},
		"Another value": {
			env: "DEV",
			validateResult: func(t *testing.T, env string) {
				require.NotEmpty(t, env)
				require.NotEqual(t, "local", env)
			},
		},
	}

	for scenario, testCase := range testCases {
		t.Run(scenario, func(t *testing.T) {
			require.NoError(t, os.Setenv("ENV", testCase.env))

			env := GetEnv()

			testCase.validateResult(t, env)

			require.NoError(t, os.Unsetenv("ENV"))
		})
	}
}

func TestIsLocalEnv(t *testing.T) {
	testCases := map[string]struct {
		env            string
		validateResult func(t *testing.T, isLocalEnv bool)
	}{
		"Empty": {
			env: "",
			validateResult: func(t *testing.T, isLocalEnv bool) {
				require.True(t, isLocalEnv)
			},
		},
		"Local": {
			env: "local",
			validateResult: func(t *testing.T, isLocalEnv bool) {
				require.True(t, isLocalEnv)
			},
		},
		"Another value": {
			env: "DEV",
			validateResult: func(t *testing.T, isLocalEnv bool) {
				require.False(t, isLocalEnv)
			},
		},
	}

	for scenario, testCase := range testCases {
		t.Run(scenario, func(t *testing.T) {
			require.NoError(t, os.Setenv("ENV", testCase.env))

			testCase.validateResult(t, IsLocalEnv())

			require.NoError(t, os.Unsetenv("ENV"))
		})
	}
}

func TestIsTestEnv(t *testing.T) {
	testCases := map[string]struct {
		env            string
		validateResult func(t *testing.T, isTestEnv bool)
	}{
		"False": {
			env: "local",
			validateResult: func(t *testing.T, isTestEnv bool) {
				require.False(t, isTestEnv)
			},
		},
		"True": {
			env: "test",
			validateResult: func(t *testing.T, isTestEnv bool) {
				require.True(t, isTestEnv)
			},
		},
	}

	for scenario, testCase := range testCases {
		t.Run(scenario, func(t *testing.T) {
			require.NoError(t, os.Setenv("ENV", testCase.env))

			testCase.validateResult(t, IsTestEnv())

			require.NoError(t, os.Setenv("ENV", "test"))
		})
	}
}

func TestIsProductionEnv(t *testing.T) {
	testCases := map[string]struct {
		env            string
		validateResult func(t *testing.T, isPrdEnv bool)
	}{
		"Local": {
			env: "local",
			validateResult: func(t *testing.T, isPrdEnv bool) {
				require.False(t, isPrdEnv)
			},
		},
		"Test": {
			env: "test",
			validateResult: func(t *testing.T, isPrdEnv bool) {
				require.False(t, isPrdEnv)
			},
		},
		"Production": {
			env: "production",
			validateResult: func(t *testing.T, isPrdEnv bool) {
				require.True(t, isPrdEnv)
			},
		},
	}

	for scenario, testCase := range testCases {
		t.Run(scenario, func(t *testing.T) {
			require.NoError(t, os.Setenv("ENV", testCase.env))

			testCase.validateResult(t, IsProductionEnv())

			require.NoError(t, os.Unsetenv("ENV"))
		})

		require.NoError(t, os.Setenv("ENV", "test"))
	}
}

func TestGetAPIVersion(t *testing.T) {
	testCases := map[string]struct {
		apiVersion     string
		validateResult func(t *testing.T, apiVersion string)
	}{
		"Empty": {
			apiVersion: "",
			validateResult: func(t *testing.T, apiVersion string) {
				require.Equal(t, "N/A", apiVersion)
			},
		},
		"Populated": {
			apiVersion: "1.0.3",
			validateResult: func(t *testing.T, apiVersion string) {
				require.Equal(t, "1.0.3", apiVersion)
			},
		},
	}

	for scenario, testCase := range testCases {
		t.Run(scenario, func(t *testing.T) {
			require.NoError(t, os.Setenv("API_VERSION", testCase.apiVersion))

			apiVersion := GetAPIVersion()

			testCase.validateResult(t, apiVersion)

			require.NoError(t, os.Unsetenv("API_VERSION"))
		})
	}
}
