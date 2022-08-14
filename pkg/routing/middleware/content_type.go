package middleware

import "net/http"

type contentType struct{}

func NewContentType() *contentType {
	return &contentType{}
}

func (m *contentType) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next(w, r)
	}
}
