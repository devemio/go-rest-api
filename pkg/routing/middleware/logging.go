package middleware

import (
	"fmt"
	"net/http"
)

type logging struct{}

func NewLogging() *logging {
	return &logging{}
}

func (m *logging) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(">", r.Method, r.URL)

		next(w, r)

		fmt.Println("<", r.Response.Status)
	}
}
