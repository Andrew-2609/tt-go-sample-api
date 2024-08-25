package apierr

// ClientErrorCode is an error that happened due to
// a client's action.
type ClientErrorCode string

// ServerErrorCode is an error that happened due to
// an internal problem.
type ServerErrorCode string

// Code is an unique identifier for a custom error.
//
// Each code is used in a context. If there's no error Code for
// a certain context, it must be created.
type Code interface {
	ClientErrorCode | ServerErrorCode
}

// ClientErrorCodes
const (
	// CodeRequestValidationError is used whenever a request
	// validation failed, e.g. a required field as not given.
	CodeRequestValidationError ClientErrorCode = "TTVALERR"
)

// ServerErrorCodes
const (
	// CodeUnknownServerError is used whenever the application
	// returns an unexpected and unknown error.
	CodeUnknownServerError ServerErrorCode = "TTINTERR"

	// CodeSQLListEmployeesFailedErrorCode is used whenever
	// the application can't list employees because of an
	// SQL error.
	CodeSQLListEmployeesFailedErrorCode ServerErrorCode = "TTSQLERR001"
)
