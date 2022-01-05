package main

import (
	"github.com/devemio/go-rest-api/internal/infrastructure/middleware"
	"github.com/devemio/go-rest-api/internal/infrastructure/routing"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var Route *routing.Route

func init() {
	middlewares := []middleware.Middleware{
		middleware.Cors,
		middleware.ContentType,
	}

	Route = routing.New(mux.NewRouter(), middlewares)
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", Route.Router()))
}
