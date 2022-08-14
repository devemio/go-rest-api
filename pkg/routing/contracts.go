package routing

import (
	"net/http"
)

type Router interface {
	Get(path string, handlerFunc http.HandlerFunc)
	Post(path string, handlerFunc http.HandlerFunc)
	Put(path string, handlerFunc http.HandlerFunc)
	Patch(path string, handlerFunc http.HandlerFunc)
	Delete(path string, handlerFunc http.HandlerFunc)
}
