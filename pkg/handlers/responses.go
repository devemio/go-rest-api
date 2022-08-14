package handlers

import (
	"net/http"
)

type Response struct {
	statusCode int
	data       any
}

func Res(statusCode int, data any) *Response {
	return &Response{
		statusCode: statusCode,
		data:       data,
	}
}

func ResCreated(data any) *Response {
	return Res(http.StatusCreated, data)
}

func ResNoContent() *Response {
	return Res(http.StatusNoContent, nil)
}
