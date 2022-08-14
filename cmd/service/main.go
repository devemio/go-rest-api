package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devemio/go-rest-api/internal/app/http/routes"
	"github.com/devemio/go-rest-api/internal/infrastructure/ioc"
	"github.com/devemio/go-rest-api/pkg/routing"
)

func main() {
	fmt.Println("Application started at :8080")

	r := routing.New()

	routes.Register(r, ioc.New())

	log.Fatal(http.ListenAndServe(":8080", r))
}
