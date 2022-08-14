package controllers

import (
	"net/http"

	"github.com/devemio/go-rest-api/internal/domain/users"
	"github.com/devemio/go-rest-api/pkg/handlers"
)

type ImageCtrl struct {
	Repo    users.UserRepo
	Service *users.Service
}

type imagesOut struct {
	Data []imageOut `json:"data"`
}

type imageOut struct {
	ID  int64  `json:"id"`
	Url string `json:"url"`
}

func (c *ImageCtrl) Get(r *http.Request) (*imagesOut, error) {
	id, err := getUserID(r)
	if err != nil {
		return nil, err
	}

	entity, err := c.Repo.Find(id)
	if err != nil {
		return nil, handlers.MapToErr(err, users.ErrUserNotFound, handlers.ErrNotFound())
	}

	images := make([]imageOut, 0, len(entity.Images))
	for _, image := range entity.Images {
		images = append(images, imageOut{
			ID:  image.ID,
			Url: c.Service.GetTemporaryUrl(image.Url),
		})
	}

	return &imagesOut{Data: images}, nil
}
