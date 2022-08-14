package handlers

import (
	"net/http"
)

type Error struct {
	statusCode int
	message    string
}

func (e *Error) Error() string {
	return e.message
}

func Err(statusCode int, message string) *Error {
	return &Error{
		statusCode: statusCode,
		message:    message,
	}
}

func ErrValidation(message string) *Error {
	return Err(http.StatusBadRequest, message)
}

func ErrUnauthorized() *Error {
	return Err(http.StatusUnauthorized, "unauthorized")
}

func ErrNotFound() *Error {
	return Err(http.StatusNotFound, "not found")
}
