package util

import "reflect"

// IsZeroValue checks if the given value is Zero
// according to its type.
//
// It uses the reflect library
// underneath, so check out its documentation if you
// don't know what Zero value means in Go.
func IsZeroValue(param any) bool {
	valueOfParam := reflect.ValueOf(param)

	if !valueOfParam.IsValid() {
		return true
	}

	return valueOfParam.IsZero()
}
