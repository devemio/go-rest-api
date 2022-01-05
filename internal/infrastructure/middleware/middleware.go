package middleware

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Apply(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	if len(m) == 0 {
		return h
	}

	wrapped := h

	for i := range m {
		wrapped = m[i](wrapped)
	}

	return wrapped
}
