package test

import (
	"io"
	"net/http"
)

// NewRequest configures a new *http.Request with all
// required headers for the application.
func NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	return req, nil
}
