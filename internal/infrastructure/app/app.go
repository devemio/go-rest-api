package app

import (
	"github.com/devemio/go-rest-api/internal/app/http/controllers"
	users2 "github.com/devemio/go-rest-api/internal/domain/users"
	"github.com/devemio/go-rest-api/internal/infrastructure/repositories/users"
)

type App struct {
	*Ctrls
}

type Ctrls struct {
	Ping *controllers.PingPongCtrl
	User *controllers.UserCtrl
}

type _repos struct {
	User users2.UserRepoContract
}

func New() *App { // @fixme move
	repos := &_repos{
		User: &users.UserRepo{},
	}

	return &App{
		Ctrls: &Ctrls{
			Ping: &controllers.PingPongCtrl{},
			User: &controllers.UserCtrl{
				Repo: repos.User,
			},
		},
	}
}
