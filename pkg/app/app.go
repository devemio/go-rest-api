package app

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/devemio/go-rest-api/pkg/contracts"
	"github.com/devemio/go-rest-api/pkg/env"
	"github.com/devemio/go-rest-api/pkg/routing/middleware"
)

type app struct {
	middlewares []contracts.Middleware
	log         contracts.Logger
	handler     http.Handler
	port        string
}

func Default(handler http.Handler) *app {
	log := logrus.New()

	return &app{
		middlewares: []contracts.Middleware{
			middleware.NewTiming(log),
			middleware.NewContentType(),
			middleware.NewPanicRecovery(log),
		},
		log:     log,
		handler: handler,
		port:    env.Get("APP_PORT", "8080"),
	}
}

func (a *app) Run() {
	for _, m := range a.middlewares {
		a.handler = m.Handle(a.handler)
	}

	a.log.WithField("port", a.port).Info("starting application")

	if err := http.ListenAndServe(":"+a.port, a.handler); err != nil {
		a.log.WithField("event", "start application").Fatal(err)
	}
}
