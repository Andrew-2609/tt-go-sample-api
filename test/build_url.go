package test

import "fmt"

// baseUrl is the base application's URL for
// API routes.
const baseUrl = "/api/v1"

// BuildURL takes the given path and returns
// it at the end of the base application URL.
func BuildURL(path string) string {
	return fmt.Sprintf("%s/%s", baseUrl, path)
}
