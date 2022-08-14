package repositories

import (
	"fmt"

	"github.com/devemio/go-rest-api/internal/domain/users/models"
)

var ErrUserNotFound = fmt.Errorf("user not found: %w", ErrRepository)

type UserRepoContract interface {
	Get() ([]models.User, error)
	Find(id int64) (*models.User, error)
	Save(*models.User) error
	Delete(id int64) error
}
