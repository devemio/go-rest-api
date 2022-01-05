package routing

import (
	"github.com/devemio/go-rest-api/internal/infrastructure/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	r *mux.Router
	m []middleware.Middleware
}

func New(r *mux.Router, m []middleware.Middleware) *Route {
	return &Route{
		r: r,
		m: m,
	}
}

func (r *Route) handle(method, uri string, h http.HandlerFunc, m ...middleware.Middleware) {
	r.r.HandleFunc(uri, middleware.Apply(h, append(r.m, m...)...)).Methods(method)
}

func (r *Route) Get(uri string, h http.HandlerFunc, m ...middleware.Middleware) {
	r.handle(http.MethodGet, uri, h, m...)
}

func (r *Route) Post(uri string, h http.HandlerFunc, m ...middleware.Middleware) {
	r.handle(http.MethodPost, uri, h, m...)
}

func (r *Route) Delete(uri string, h http.HandlerFunc, m ...middleware.Middleware) {
	r.handle(http.MethodDelete, uri, h, m...)
}

func (r *Route) Router() http.Handler {
	return r.r
}
