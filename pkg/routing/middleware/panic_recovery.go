package middleware

import (
	"encoding/json"
	"net/http"
	"runtime/debug"

	"github.com/devemio/go-rest-api/pkg/contracts"
)

type panicRecovery struct {
	log contracts.Logger
}

func NewPanicRecovery(log contracts.Logger) *panicRecovery {
	return &panicRecovery{
		log: log,
	}
}

func (m *panicRecovery) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)

				_ = json.NewEncoder(w).Encode(map[string]string{
					"message": "internal server error",
				})

				m.log.WithFields(map[string]interface{}{
					"err":    err,
					"url":    r.URL,
					"method": r.Method,
					"trace":  string(debug.Stack()),
				}).Error("panic")
			}
		}()

		next.ServeHTTP(w, r)
	})
}
