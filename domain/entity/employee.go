package entity

import (
	"strings"
	"time"
)

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

// NewEmployee returns a pointer of Employee
// with the given name.
//
// The returned employee will always have the
// first letter of his name uppercased.
func NewEmployee(name string) *Employee {
	return &Employee{
		Name: strings.ToUpper(string(name[0])) + name[1:],
	}
}
