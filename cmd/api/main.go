package main

import (
	"github.com/devemio/go-rest-api/cmd/api/users"
	"github.com/devemio/go-rest-api/internal/infrastructure/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Start application")

	r := mux.NewRouter()

	r.HandleFunc("/users", handle(users.Get)).Methods("GET")
	r.HandleFunc("/users/{id}", handle(users.Find)).Methods("GET")
	r.HandleFunc("/users", handle(users.Create)).Methods("POST")
	r.HandleFunc("/users/{id}", handle(users.Delete)).Methods("DELETE")

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))

	//app := application.New(config.New())
	//if err := app.Start(); err != nil {
	//	log.Fatal(err)
	//}
	//
	//c := collection.New()
	//c.Put("test", 1)
	//fmt.Println(c.Contains("test"))
}

func handle(next http.HandlerFunc) http.HandlerFunc {
	middlewares := []middleware.Middleware{
		middleware.Cors,
		middleware.ContentType,
	}

	return middleware.Apply(next, middlewares...)
}
