package apierror

import (
	"fmt"
	"net/http"
)

// APIError define error to be returned in response
type APIError struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func New(code string) error {
	return &APIError{
		Code:    code,
		Message: "An error has occured: " + code,
	}
}

func WithMessage(code string, message string) error {
	return &APIError{
		Code:    code,
		Message: message,
	}
}

func WithData(code string, message string, data interface{}) error {
	return &APIError{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func FromError(code string, err error) error {
	return &APIError{
		Code:    code,
		Message: err.Error(),
	}
}

func InternalError(err error) error {
	return &APIError{
		Code:    CodeInternalServerError,
		Message: err.Error(),
	}
}

// GetCode return error code from err object. will return InternalError if err is not APIError
func GetCode(err error) string {
	if f, ok := err.(*APIError); ok {
		return f.Code
	}
	return CodeInternalServerError
}

// ToHTTPStatusFrom return appropiate http status from error
func GetHTTPStatus(err error) int {
	code := GetCode(err)
	if errorDictionary == nil || errorDictionary[code] == 0 {
		return http.StatusInternalServerError
	}
	return errorDictionary[code]
}

const (
	// CodeInternalServerError is default code to handle internal error (500)
	CodeInternalServerError = "InternalServerError"
)
