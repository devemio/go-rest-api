package routes

import (
	"net/http"

	"github.com/devemio/go-rest-api/internal/app/providers"
	"github.com/devemio/go-rest-api/pkg/contracts"
	"github.com/devemio/go-rest-api/pkg/handlers"
	"github.com/devemio/go-rest-api/pkg/routing"
)

func Create() http.Handler {
	router := routing.New()

	register(router, providers.App().Ctrls)

	return router
}

func register(r contracts.Router, c *providers.Ctrls) {
	r.Get("/ping", handlers.WrapGet(c.Ping.Ping))

	r.Get("/users", handlers.WrapGet(c.User.Get))
	r.Get("/users/{id}", handlers.WrapGetWithReq(c.User.Find))
	r.Post("/users", handlers.WrapPost(c.User.Create))
	r.Delete("/users/{id}", handlers.WrapGetWithReq(c.User.Delete))

	r.Get("/users/{id}/images", handlers.WrapGetWithReq(c.Image.Get))
}
