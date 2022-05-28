package client

// Error struct for wrapping api errors
type Error struct {
	ErrorCode string       `json:"errorCode"`
	Status    string       `json:"status"`
	Message   string       `json:"message"`
	Errors    []FieldError `json:"errors"`
}

// FieldError struct for wrapping validation errors of a particular field
type FieldError struct {
	InvalidValue string `json:"invalidValue"`
	Field        string `json:"field"`
	Message      string `json:"message"`
}

// Error returns an error's message
func (err *Error) Error() string {
	return err.Message
}

// NotFoundError is a specific error for 404 case
type NotFoundError struct {
	ErrorCode string `json:"errorCode"`
	Status    string `json:"status"`
	Message   string `json:"message"`
}

// Error returns an error's message
func (err *NotFoundError) Error() string {
	return err.Message
}
