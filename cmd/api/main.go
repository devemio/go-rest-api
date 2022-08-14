package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/devemio/go-rest-api/internal/infrastructure/middleware"
	"github.com/devemio/go-rest-api/internal/infrastructure/routing"
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
