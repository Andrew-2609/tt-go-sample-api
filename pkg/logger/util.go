package logger

import "tt-go-sample-api/util"

// removeEmptyAttributes removes any empty attribute
// of a given map[string]any. It runs recursively if
// there are nested maps.
func removeEmptyAttributes(input map[string]any) {
	for key, value := range input {
		if innerValue, ok := value.(map[string]any); ok {
			removeEmptyAttributes(innerValue)
		}

		if util.IsZeroValue(value) {
			delete(input, key)
		}
	}
}
