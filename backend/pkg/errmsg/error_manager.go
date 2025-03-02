package errmsg

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CustomError represents an application-specific error
type CustomError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
	Err     error  `json:"-"`
}

// Error implements the error interface
func (e *CustomError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %s -> %v", e.Code, e.Message, e.Details, e.Err)
	}
	return fmt.Sprintf("[%s] %s: %s", e.Code, e.Message, e.Details)
}

// Unwrap allows errors.As() to extract the original error
func (e *CustomError) Unwrap() error {
	return e.Err
}

// ErrorResponse represents a JSON response for Echo
type ErrorResponse struct {
	Error CustomError `json:"error"`
}

// Predefined error definitions
var ErrorDefinitions = map[string]struct {
	StatusCode int
	Message    string
}{
	"ERR_BAD_REQUEST":     {http.StatusBadRequest, "Invalid request parameters."},
	"ERR_UNAUTHORIZED":    {http.StatusUnauthorized, "Authentication required."},
	"ERR_NOT_FOUND":       {http.StatusNotFound, "Resource not found."},
	"ERR_INTERNAL_SERVER": {http.StatusInternalServerError, "Unexpected server error."},
}

// NewError creates a new CustomError
func NewError(errorCode, details string) error {
	return &CustomError{
		Code:    errorCode,
		Message: getErrorMessage(errorCode),
		Details: details,
	}
}

// WrapError wraps an existing error into a CustomError
func WrapError(errorCode, details string, err error) error {
	return &CustomError{
		Code:    errorCode,
		Message: getErrorMessage(errorCode),
		Details: details,
		Err:     err,
	}
}

// GetErrorCode extracts the error code from a wrapped error
func GetErrorCode(err error) string {
	var customErr *CustomError
	if errors.As(err, &customErr) {
		return customErr.Code
	}
	return "ERR_INTERNAL_SERVER"
}

// RespondError sends an error response using Echo
func RespondError(c echo.Context, err error) error {
	var customErr *CustomError
	if !errors.As(err, &customErr) {
		customErr = &CustomError{
			Code:    "ERR_INTERNAL_SERVER",
			Message: "Unexpected server error.",
		}
	}

	statusCode := getErrorStatusCode(customErr.Code)
	response := ErrorResponse{
		Error: *customErr,
	}

	return c.JSON(statusCode, response)
}

// getErrorMessage retrieves the error message based on errorCode
func getErrorMessage(errorCode string) string {
	if def, exists := ErrorDefinitions[errorCode]; exists {
		return def.Message
	}
	return "Unexpected error occurred."
}

// getErrorStatusCode retrieves the HTTP status code based on errorCode
func getErrorStatusCode(errorCode string) int {
	if def, exists := ErrorDefinitions[errorCode]; exists {
		return def.StatusCode
	}
	return http.StatusInternalServerError
}
