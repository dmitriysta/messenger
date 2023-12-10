//go:generate mockery

package interfaces

import (
	"context"

	"github.com/dmitriysta/messenger/user/internal/models"
)

type UserService interface {
	CreateUser(ctx context.Context, name, email, password string) (*models.User, error)
	GetUserById(ctx context.Context, userId int) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, userId int) error
}
