package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type timing struct{}

func NewTiming() *timing {
	return &timing{}
}

func (m *timing) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next(w, r)

		time.Sleep(10 * time.Millisecond) // @fixme

		fmt.Println("Response time", r.Method, r.URL, time.Since(start))
	}
}
