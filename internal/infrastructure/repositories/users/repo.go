package users

import (
	domain "github.com/devemio/go-rest-api/internal/domain/users"
)

type UserRepo struct{}

var users = []domain.User{
	{
		ID:           1,
		Username:     "A",
		EmailAddress: "a@gmai.com",
		Images: []domain.Image{
			{
				ID:  1,
				Url: "https://images.com/001.jpg",
			},
			{
				ID:  2,
				Url: "https://images.com/002.jpg",
			},
		},
	},
	{
		ID:           2,
		Username:     "B",
		EmailAddress: "b@gmai.com",
		Images: []domain.Image{
			{
				ID:  3,
				Url: "https://images.com/003.jpg",
			},
			{
				ID:  4,
				Url: "https://images.com/004.jpg",
			},
		},
	},
}

func (r *UserRepo) Get() ([]domain.User, error) {
	return users, nil
}

func (r *UserRepo) Find(id int64) (*domain.User, error) {
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}

	return nil, domain.ErrUserNotFound
}

func (r *UserRepo) Save(user *domain.User) error {
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

	return domain.ErrUserNotFound
}
