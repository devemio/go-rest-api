package controllers

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/devemio/go-rest-api/internal/domain/users"
	"github.com/devemio/go-rest-api/pkg/handlers"
	"github.com/devemio/go-rest-api/pkg/routing"
)

type UserCtrl struct {
	Repo users.UserRepoContract
}

type usersOut struct {
	Data []userOut `json:"data"`
}

type userOut struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

func (c *UserCtrl) Get() (*usersOut, error) {
	entities, err := c.Repo.Get()
	if err != nil {
		return nil, err
	}

	out := &usersOut{
		Data: make([]userOut, 0, len(entities)),
	}

	for _, entity := range entities {
		out.Data = append(out.Data, userOut{
			ID:       entity.ID,
			Username: entity.Username,
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

	entity, err := c.Repo.Find(id)
	if err != nil {
		return nil, handlers.MapToErr(err, users.ErrUserNotFound, handlers.ErrNotFound())
	}

	return &userOut{
		ID:       entity.ID,
		Username: entity.Username,
	}, nil
}

type createUserIn struct {
	Username string `json:"username"`
}

func (c *UserCtrl) Create(in *createUserIn) (*handlers.Response, error) {
	id := rand.Int63()
	entity := &users.User{
		ID:           id,
		Username:     in.Username,
		EmailAddress: in.Username + "@gmail.com",
	}

	if err := c.Repo.Save(entity); err != nil {
		return nil, err
	}

	entity, err := c.Repo.Find(id)
	if err != nil {
		return nil, err
	}

	return handlers.ResCreated(&userOut{
		ID:       entity.ID,
		Username: entity.Username,
	}), nil
}

func (c *UserCtrl) Delete(r *http.Request) (*handlers.Response, error) {
	params := routing.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		return nil, handlers.ErrValidation("id should be int")
	}

	if err = c.Repo.Delete(id); err != nil {
		return nil, handlers.MapToErr(err, users.ErrUserNotFound, handlers.ErrNotFound())
	}

	return handlers.ResNoContent(), nil
}
