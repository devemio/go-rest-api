package middleware

import "net/http"

type Middleware interface {
	Handle(next http.HandlerFunc) http.HandlerFunc
}

type RMiddleware interface {
	Handle(next http.Handler) http.Handler
}
