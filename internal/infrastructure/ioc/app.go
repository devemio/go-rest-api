package ioc

import (
	"github.com/devemio/go-rest-api/internal/app/http/controllers"
	"github.com/devemio/go-rest-api/internal/domain/users/repositories"
	repoImpl "github.com/devemio/go-rest-api/internal/infrastructure/impl/repositories"
)

type App struct {
	*Ctrls
}

type Ctrls struct {
	Ping *controllers.PingPongCtrl
	User *controllers.UserCtrl
}

type _repos struct {
	User repositories.UserRepoContract
}

func New() *App {
	repos := &_repos{
		User: &repoImpl.UserRepo{},
	}

	return &App{
		&Ctrls{
			Ping: &controllers.PingPongCtrl{},
			User: &controllers.UserCtrl{
				Repo: repos.User,
			},
		},
	}
}
