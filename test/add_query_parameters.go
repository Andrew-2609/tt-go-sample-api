package test

import (
	"fmt"
	"strings"
	"tt-go-sample-api/util"
)

// AddQueryParameters adds query parameters to the
// given base path.
func AddQueryParameters(basePath string, params map[string]any) string {
	completePath := basePath

	if !strings.Contains(completePath, "?") {
		completePath += "?"
	}

	var queryParams []string

	for k, v := range params {
		if util.IsZeroValue(v) {
			continue
		}

		queryParams = append(queryParams, fmt.Sprintf("%s=%v", k, v))
	}

	return completePath + strings.Join(queryParams, "&")
}
