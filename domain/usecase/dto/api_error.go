package dto

// APIErrorOutputDTO shall be used for whenever the
// application must return an error to the client.
type APIErrorOutputDTO struct {
	// Error is the error that happened during the
	// request handling.
	Error error `json:"error"`
}
