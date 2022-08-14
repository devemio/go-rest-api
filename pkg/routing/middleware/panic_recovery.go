package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
)

type panicRecovery struct{}

func NewPanicRecovery() *panicRecovery {
	return &panicRecovery{}
}

func (m *panicRecovery) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)

				_ = json.NewEncoder(w).Encode(map[string]string{
					"message": "internal server error",
				})

				fmt.Println("Panic", r.Method, r.URL, string(debug.Stack()))
			}
		}()

		next(w, r)
	}
}
