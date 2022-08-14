package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type In any
type Out any

func WrapGet[TOut Out](fn func() (TOut, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		out, err := fn()
		handle(w, req, out, err)
	}
}

func WrapGetWithReq[TOut Out](fn func(*http.Request) (TOut, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		out, err := fn(req)
		handle(w, req, out, err)
	}
}

func WrapPost[TIn In, TOut Out](fn func(TIn) (TOut, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		bytes, err := io.ReadAll(req.Body)
		if err != nil {
			internalServerError(w, fmt.Errorf("IO: %w", err))
			return
		}

		var in TIn
		if err = json.Unmarshal(bytes, &in); err != nil {
			internalServerError(w, fmt.Errorf("DTO: %w", err))
			return
		}

		out, err := fn(in)
		handle(w, req, out, err)
	}
}

func handle(w http.ResponseWriter, req *http.Request, out Out, err error) {
	if err == nil {
		switch r := out.(type) {
		case *Response:
			w.WriteHeader(r.statusCode)

			if r.data != nil {
				_ = json.NewEncoder(w).Encode(map[string]any{
					"data": r.data,
				})
			}
		default:
			w.WriteHeader(http.StatusOK)

			_ = json.NewEncoder(w).Encode(out)
		}

		return
	}

	var e *Error
	if errors.As(err, &e) {
		w.WriteHeader(e.statusCode)

		_ = json.NewEncoder(w).Encode(map[string]string{
			"message": e.message,
		})

		return
	}

	internalServerError(w, errors.New("internal server error"))

	log.WithFields(log.Fields{
		"err":    err,
		"url":    req.URL,
		"method": req.Method,
	}).Error("5xx response")
}

func internalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": err.Error(),
	})
}
