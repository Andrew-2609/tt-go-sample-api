package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewEmployee(t *testing.T) {
	lowercasedEmployee := NewEmployee("andrew")
	uppercasedEmployee := NewEmployee("Karine")

	require.Equal(t, "Andrew", lowercasedEmployee.Name)
	require.Equal(t, "Karine", uppercasedEmployee.Name)
}
