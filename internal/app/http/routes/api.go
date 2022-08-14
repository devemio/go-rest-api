package routes

import (
	"github.com/devemio/go-rest-api/internal/infrastructure/ioc"
	"github.com/devemio/go-rest-api/pkg/handlers"
	"github.com/devemio/go-rest-api/pkg/routing"
)

func Register(r routing.Router, app *ioc.App) {
	r.Get("/ping", handlers.WrapGet(app.Ctrls.Ping.Ping))

	r.Get("/users", handlers.WrapGet(app.Ctrls.User.Get))
	r.Get("/users/{id}", handlers.WrapGet(app.Ctrls.User.Find))
	r.Post("/users", handlers.WrapPost(app.Ctrls.User.Create))
	r.Delete("/users/{id}", handlers.WrapGet(app.Ctrls.User.Delete))
}
