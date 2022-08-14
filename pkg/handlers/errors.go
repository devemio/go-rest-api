package handlers

import (
	"net/http"
)

type errBase struct {
	statusCode int
	message    string
}

func (e *errBase) Error() string {
	return e.message
}

func Err(statusCode int, message string) *errBase {
	return &errBase{
		statusCode: statusCode,
		message:    message,
	}
}

func ErrValidation(message string) *errBase {
	return Err(http.StatusBadRequest, message)
}

func ErrUnauthorized() *errBase {
	return Err(http.StatusUnauthorized, "unauthorized")
}

func ErrNotFound() *errBase {
	return Err(http.StatusNotFound, "not found")
}
