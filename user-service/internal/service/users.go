package service

import (
	"context"
	"fmt"
	"go13/user-service/internal/dto"
	"go13/user-service/internal/models"
	"go13/user-service/internal/repo"

	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	usersRepo repo.UsersRepo
}

func NewUsersService(usersRepo repo.UsersRepo) *UsersService {
	return &UsersService{
		usersRepo: usersRepo,
	}
}

func (us *UsersService) ChangePassword(ctx context.Context, userId string, input dto.ChangePasswordInput) error {
	op := "UsersService.ChangePassword"

	user, err := us.usersRepo.GetUserById(ctx, userId)
	if err != nil {
		return fmt.Errorf("%s: usersRepo.GetUserById: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(input.OldPassword)); err != nil {
		return models.ErrInvalidCredentials
	}

	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%s: bcrypt.GenerateFromPassword: %w", op, err)
	}

	_, err = us.usersRepo.UpdateUser(ctx, userId, dto.UpdateUserInput{
		HashedPassword: dto.NewOptString(string(newHashedPassword)),
	})
	if err != nil {
		return fmt.Errorf("%s: usersRepo.UpdateUser: %w", op, err)
	}

	return nil
}

func (us *UsersService) GetUserById(ctx context.Context, userId string) (dto.UserInfo, error) {
	op := "UsersService.GetUserById"

	user, err := us.usersRepo.GetUserById(ctx, userId)
	if err != nil {
		return dto.UserInfo{}, fmt.Errorf("%s: usersRepo.GetUserById: %w", op, err)
	}

	return dto.UserInfo{
		Id:       userId,
		Username: user.Username,
		Bio:      user.Bio,
	}, nil
}

func (us *UsersService) UpdateMe(ctx context.Context, userId string, input dto.UpdateMeInput) (dto.UserInfo, error) {
	op := "UsersService.UpdateMe"

	user, err := us.usersRepo.UpdateUser(ctx, userId, dto.UpdateUserInput{
		Username: dto.NewOptString(input.Username),
		Bio:      dto.NewOptString(input.Bio),
	})
	if err != nil {
		return dto.UserInfo{}, fmt.Errorf("%s: usersRepo.UpdateUser: %w", op, err)
	}

	return dto.UserInfo{
		Id:       user.Id,
		Username: user.Username,
		Bio:      user.Bio,
	}, nil
}
