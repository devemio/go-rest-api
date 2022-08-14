package providers

import (
	"github.com/devemio/go-rest-api/internal/app/http/controllers"
	"github.com/devemio/go-rest-api/internal/domain/users"
	usersImpl "github.com/devemio/go-rest-api/internal/infrastructure/repositories/users"
)

type app struct {
	*Ctrls
}

type Ctrls struct {
	Ping  *controllers.PingPongCtrl
	User  *controllers.UserCtrl
	Image *controllers.ImageCtrl
}

type _repos struct {
	User users.UserRepo
}

func App() *app {
	repos := &_repos{
		User: &usersImpl.UserRepo{},
	}

	return &app{
		Ctrls: &Ctrls{
			Ping: &controllers.PingPongCtrl{},
			User: &controllers.UserCtrl{
				Repo: repos.User,
			},
			Image: &controllers.ImageCtrl{
				Repo:    repos.User,
				Service: &users.Service{},
			},
		},
	}
}
