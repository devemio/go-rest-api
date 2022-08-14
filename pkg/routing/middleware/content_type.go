package middleware

import "net/http"

type contentType struct{}

func NewContentType() *contentType {
	return &contentType{}
}

func (m *contentType) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
