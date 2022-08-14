package repositories

import (
	"github.com/devemio/go-rest-api/internal/domain/users/models"
	"github.com/devemio/go-rest-api/internal/domain/users/repositories"
)

type UserRepo struct{}

var users = []models.User{
	{
		ID:           1,
		Username:     "A",
		EmailAddress: "a@gmai.com",
	},
	{
		ID:           2,
		Username:     "B",
		EmailAddress: "b@gmai.com",
	},
}

func (r *UserRepo) Get() ([]models.User, error) {
	return users, nil
}

func (r *UserRepo) Find(id int64) (*models.User, error) {
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}

	return nil, repositories.ErrUserNotFound
}

func (r *UserRepo) Save(user *models.User) error {
	users = append(users, *user)

	return nil
}

func (r *UserRepo) Delete(id int64) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)

			return nil
		}
	}

	return repositories.ErrUserNotFound
}
