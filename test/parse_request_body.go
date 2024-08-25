package test

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

// ParseRequestBody tries to parse any object to an
// io.Reader to be used in HTTP requests.
func ParseRequestBody(t *testing.T, v any) io.Reader {
	body, err := json.Marshal(v)
	require.NoError(t, err)

	return bytes.NewReader(body)
}
