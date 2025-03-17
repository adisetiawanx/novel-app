package helper

import "net/http"

type BaseError struct {
	Message    string
	StatusCode int
}

func (e *BaseError) Error() string {
	return e.Message
}

func NewUserInputError(message string) *BaseError {
	return &BaseError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
}

func NewAuthenticationError(message string) *BaseError {
	return &BaseError{
		Message:    message,
		StatusCode: http.StatusUnauthorized,
	}
}

func NewPermissionError(message string) *BaseError {
	return &BaseError{
		Message:    message,
		StatusCode: http.StatusForbidden,
	}
}

func NewConflictError(message string) *BaseError {
	return &BaseError{
		Message:    message,
		StatusCode: http.StatusConflict,
	}
}

func NewInternalServerError() *BaseError {
	return &BaseError{
		Message:    "unable to process request, please try again later",
		StatusCode: http.StatusInternalServerError,
	}
}
