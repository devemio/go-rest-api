package main

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithField("port", 4000).Info("starting server")

	http.HandleFunc("/", rootHandler)

	if err := http.ListenAndServe(":4000", logRequest(http.DefaultServeMux)); err != nil {
		log.WithField("event", "start server").Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<h1>Hello World</h1><div>Welcome to whereever you are</div>")
}

type responseWriter struct {
	http.ResponseWriter

	b          []byte
	statusCode int

	name        func() string
	wroteHeader bool
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.b = b
	return w.ResponseWriter.Write(b)
}

func (w *responseWriter) WriteHeader(statusCode int) {
	if w.wroteHeader == false {
		w.ResponseWriter.Header().Set("Server", w.name())
		w.wroteHeader = true
	}

	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func logRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		var duration time.Duration

		rw := &responseWriter{
			ResponseWriter: w,
			name: func() string {
				return duration.String()
			},
		}

		h.ServeHTTP(rw, r)

		duration = time.Since(start)

		log.WithFields(log.Fields{
			"uri":      r.RequestURI,
			"method":   r.Method,
			"duration": duration,
			"status":   rw.statusCode,
			"body":     string(rw.b),
		}).Info("> request")
	})
}
