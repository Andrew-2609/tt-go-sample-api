package dto

// RequireEmployeeFromHRInputDTO represens the
// necessary information to require an employee
// from the HR team.
type RequireEmployeeFromHRInputDTO struct {
	Stack string `json:"stack" validate:"notblank"`
}

// RequireEmployeeFromHROutputDTO represents the
// response of a successful employee requisition.
type RequireEmployeeFromHROutputDTO struct {
	Message string `json:"message"`
}
