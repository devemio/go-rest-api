package routing

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/devemio/go-rest-api/pkg/routing/middleware"
)

type router struct {
	r *mux.Router
	m []middleware.Middleware
}

func New() *router {
	return &router{
		r: mux.NewRouter().StrictSlash(true),
		m: []middleware.Middleware{
			middleware.NewPanicRecovery(),
			middleware.NewTiming(),
			middleware.NewContentType(), // @fixme
			middleware.NewLogging(),
		},
	}
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.r.ServeHTTP(w, req)
}

func (r *router) Get(path string, handlerFunc http.HandlerFunc) {
	r.r.HandleFunc(path, r.decorate(handlerFunc)).Methods(http.MethodGet)
}

func (r *router) Post(path string, handlerFunc http.HandlerFunc) {
	r.r.HandleFunc(path, r.decorate(handlerFunc)).Methods(http.MethodPost)
}

func (r *router) Put(path string, handlerFunc http.HandlerFunc) {
	r.r.HandleFunc(path, r.decorate(handlerFunc)).Methods(http.MethodPut)
}

func (r *router) Patch(path string, handlerFunc http.HandlerFunc) {
	r.r.HandleFunc(path, r.decorate(handlerFunc)).Methods(http.MethodPatch)
}

func (r *router) Delete(path string, handlerFunc http.HandlerFunc) {
	r.r.HandleFunc(path, r.decorate(handlerFunc)).Methods(http.MethodDelete)
}

func (r *router) decorate(fn http.HandlerFunc) http.HandlerFunc {
	for _, m := range r.m {
		fn = m.Handle(fn)
	}

	return fn
}
