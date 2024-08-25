package dto

import "time"

// ListEmployeesInputDTO represents the
// required parameters in order to retrieve
// a list of employees.
type ListEmployeesInputDTO struct {
	Page  int32 `json:"page" validate:"required,min=1"`
	Limit int32 `json:"limit" validate:"omitempty,min=1,max=100"`
}

// ListEmployeesOutputDTO represents the
// response of an employees listing, that
// may contain several items.
type ListEmployeesOutputDTO struct {
	Items []EmployeeItemOutputDTO `json:"items"`
}

// EmployeeItemOutputDTO represents a single
// employee contained in the output DTO's list.
type EmployeeItemOutputDTO struct {
	PublicID  string    `json:"publicId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
