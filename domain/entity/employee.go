package entity

import "time"

// Employee represents an employee of TT.
//
// Its attributes don't have the `json` tag
// because an Employee entity should never
// reach the handler layer. Input and Output
// DTOs should be used instead.
type Employee struct {
	ID        int64
	PublicID  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
