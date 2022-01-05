package main

import (
	"fmt"
	collection "github.com/devemio/go-rest-api/internal/illuminate"
	"github.com/devemio/go-rest-api/internal/infrastructure/application"
	"github.com/devemio/go-rest-api/internal/infrastructure/config"
	"log"
)

func main() {
	app := application.New(config.New())
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

	c := collection.New()
	c.Put("test", 1)
	fmt.Println(c.Contains("test"))
}
