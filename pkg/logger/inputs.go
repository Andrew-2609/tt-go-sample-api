package logger

// LogInput receives a message and an optional
// data object to be logged.
type LogInput struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// GetMessage returns the LogInput's inner message.
func (li LogInput) GetMessage() string {
	return li.Message
}

// GetFields returns the LogInput's inner
// data, formatted as a "data" map[string]any
// value.
func (li LogInput) GetFields() map[string]any {
	fields := map[string]any{"data": li.Data}

	removeEmptyAttributes(fields)

	return fields
}
