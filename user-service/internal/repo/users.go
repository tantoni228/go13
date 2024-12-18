package repo

import (
	"context"
	"go13/user-service/internal/dto"
	"go13/user-service/internal/models"
)

type UsersRepo interface {
	AddUser(ctx context.Context, user models.User) (models.User, error)
	GetUserById(ctx context.Context, userId string) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	UpdateUser(ctx context.Context, userId string, input dto.UpdateUserInput) (models.User, error)
}
