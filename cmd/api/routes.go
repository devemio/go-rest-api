package main

import (
	"github.com/devemio/go-rest-api/cmd/api/users"
)

func init() {
	Route.Get("/users", users.Get)
	Route.Get("/users/{id}", users.Find)
	Route.Post("/users", users.Create)
	Route.Delete("/users/{id}", users.Delete)
}
