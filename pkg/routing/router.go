package routing

import (
	"net/http"

	"github.com/gorilla/mux"
)

type router struct {
	r *mux.Router
}

func New() *router {
	return &router{
		r: mux.NewRouter().StrictSlash(true),
	}
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.r.ServeHTTP(w, req)
}

func (r *router) Get(path string, handlerFunc http.HandlerFunc) {
	r.r.HandleFunc(path, handlerFunc).Methods(http.MethodGet)
}

func (r *router) Post(path string, handlerFunc http.HandlerFunc) {
	r.r.HandleFunc(path, handlerFunc).Methods(http.MethodPost)
}

func (r *router) Put(path string, handlerFunc http.HandlerFunc) {
	r.r.HandleFunc(path, handlerFunc).Methods(http.MethodPut)
}

func (r *router) Patch(path string, handlerFunc http.HandlerFunc) {
	r.r.HandleFunc(path, handlerFunc).Methods(http.MethodPatch)
}

func (r *router) Delete(path string, handlerFunc http.HandlerFunc) {
	r.r.HandleFunc(path, handlerFunc).Methods(http.MethodDelete)
}
