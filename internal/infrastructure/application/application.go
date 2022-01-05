package application

import (
	"github.com/devemio/go-rest-api/internal/infrastructure/config"
)

type Application struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Application {
	return &Application{
		cfg: cfg,
	}
}

func (a *Application) Start() error {
	return nil
}
