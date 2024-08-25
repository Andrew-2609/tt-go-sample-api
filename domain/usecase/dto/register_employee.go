package dto

import "time"

// RegisterEmployeeInputDTO represents the
// necessary attributes to register an
// employee.
type RegisterEmployeeInputDTO struct {
	Name string `json:"name"`
}

// RegisterEmployeeOutputDTO represents the
// response of a sucessful employee registration.
type RegisterEmployeeOutputDTO struct {
	Message     string                         `json:"message"`
	NewEmployee NewRegisteredEmployeeOutputDTO `json:"newEmployee"`
	Suggestion  string                         `json:"suggestion,omitempty"`
}

// NewRegisteredEmployeeOutputDTO represents
// the newly registrated employee's information.
type NewRegisteredEmployeeOutputDTO struct {
	ID        int64     `json:"id"`
	PublicID  string    `json:"publicId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
