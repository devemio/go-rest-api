package main

import (
	"github.com/devemio/go-rest-api/internal/app/http/routes"
	"github.com/devemio/go-rest-api/pkg/app"
)

func main() {
	app.Default(routes.Create()).Run()
}
