package users

import (
	"errors"
	"fmt"
)

var (
	ErrRepository   = errors.New("repository")
	ErrUserNotFound = fmt.Errorf("user not found: %w", ErrRepository)
)

type UserRepo interface {
	Get() ([]User, error)
	Find(id int64) (*User, error)
	Save(*User) error
	Delete(id int64) error
}
