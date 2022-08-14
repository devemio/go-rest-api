package middleware

import (
	"net/http"
	"time"

	"github.com/devemio/go-rest-api/pkg/contracts"
)

type timing struct {
	log contracts.Logger
}

func NewTiming(log contracts.Logger) *timing {
	return &timing{
		log: log,
	}
}

func (m *timing) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)

		m.log.WithFields(map[string]interface{}{
			"method":   r.Method,
			"url":      r.URL,
			"duration": duration,
		}).Info("response time")
	})
}
