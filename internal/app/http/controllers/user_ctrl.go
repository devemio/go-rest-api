package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/devemio/go-rest-api/internal/domain/users/models"
	"github.com/devemio/go-rest-api/internal/domain/users/repositories"
	"github.com/devemio/go-rest-api/pkg/handlers"
	"github.com/devemio/go-rest-api/pkg/routing"
)

type UserCtrl struct {
	Repo repositories.UserRepoContract
}

type usersOut struct {
	Data []userOut `json:"data"`
}

type userOut struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

func (c *UserCtrl) Get(*http.Request) (*usersOut, error) {
	users, err := c.Repo.Get()
	if err != nil {
		return nil, err
	}

	out := &usersOut{
		Data: make([]userOut, 0, len(users)),
	}

	for _, u := range users {
		out.Data = append(out.Data, userOut{
			ID:       u.ID,
			Username: u.Username,
		})
	}

	return out, nil
}

func (c *UserCtrl) Find(r *http.Request) (*userOut, error) {
	params := routing.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		return nil, handlers.ErrValidation("id should be int")
	}

	user, err := c.Repo.Find(id)
	if err != nil {
		return nil, handlers.MapToErr(err, repositories.ErrUserNotFound, handlers.ErrNotFound())
	}

	return &userOut{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}

type createUserIn struct {
	Username string `json:"username"`
}

func (c *UserCtrl) Create(_ *http.Request, in *createUserIn) (*handlers.Response, error) {
	fmt.Println("IN", in)

	id := rand.Int63()
	user := &models.User{
		ID:           id,
		Username:     in.Username,
		EmailAddress: in.Username + "@gmail.com",
	}

	if err := c.Repo.Save(user); err != nil {
		return nil, err
	}

	user, err := c.Repo.Find(id)
	if err != nil {
		return nil, err
	}

	return handlers.ResCreated(&userOut{
		ID:       user.ID,
		Username: user.Username,
	}), nil
}

func (c *UserCtrl) Delete(r *http.Request) (*handlers.Response, error) {
	params := routing.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		return nil, handlers.ErrValidation("id should be int")
	}

	if err = c.Repo.Delete(id); err != nil {
		return nil, handlers.MapToErr(err, repositories.ErrUserNotFound, handlers.ErrNotFound())
	}

	return handlers.ResNoContent(), nil
}
