package utils

import (
	"fmt"
)

// CustomError represents
type CustomError struct {
	ResourceName string
	Message      string
	Err          error
}

func (ce *CustomError) Error() string {
	if ce.Err != nil {
		return fmt.Sprintf("[%s] %s: %s", ce.ResourceName, ce.Message, ce.Err.Error())
	}
	return fmt.Sprintf("[%s] %s", ce.ResourceName, ce.Message)
}

// NewCustomError creates a new CustomError instance.
func NewCustomError(resourceName, message string, err error) *CustomError {
	return &CustomError{
		ResourceName: resourceName,
		Message:      message,
		Err:          err,
	}
}

// HandleError handles errors and returns a CustomError if necessary.
func HandleError(resourceName, message string, err error) error {
	if err != nil {
		return NewCustomError(resourceName, message, err)
	}
	return nil
}
