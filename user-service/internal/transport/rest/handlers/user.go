package handlers

import (
	"context"
	"errors"
	"go13/pkg/logger"
	api "go13/pkg/ogen/users-service"
	"go13/user-service/internal/dto"
	"go13/user-service/internal/models"
	"go13/user-service/internal/transport/rest/auth"

	"go.uber.org/zap"
)

type UsersService interface {
	ChangePassword(ctx context.Context, userId string, input dto.ChangePasswordInput) error
	GetUserById(ctx context.Context, userId string) (dto.UserInfo, error)
	UpdateMe(ctx context.Context, userId string, input dto.UpdateMeInput) (dto.UserInfo, error)
}

type UsersHandler struct {
	usersService UsersService
}

func NewUsersHandler(usersService UsersService) *UsersHandler {
	return &UsersHandler{
		usersService: usersService,
	}
}

// ChangePassword implements changePassword operation.
//
// Change password.
//
// POST /users/me/change-password
func (uh *UsersHandler) ChangePassword(ctx context.Context, req *api.ChangePasswordReq) (api.ChangePasswordRes, error) {
	err := uh.usersService.ChangePassword(ctx, auth.UserIdFromCtx(ctx), dto.ChangePasswordInput{
		OldPassword: string(req.GetOldPassword()),
		NewPassword: string(req.GetNewPassword()),
	})
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			return &api.ChangePasswordForbidden{}, nil
		}
		logger.FromCtx(ctx).Error("change password", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}

	return &api.ChangePasswordNoContent{}, nil
}

// GetMe implements getMe operation.
//
// Get user info by token.
//
// GET /users/me
func (uh *UsersHandler) GetMe(ctx context.Context) (api.GetMeRes, error) {
	userInfo, err := uh.usersService.GetUserById(ctx, auth.UserIdFromCtx(ctx))
	if err != nil {
		logger.FromCtx(ctx).Error("get me", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}
	return &api.User{
		ID:       api.UserId(userInfo.Id),
		Username: userInfo.Username,
		Bio:      userInfo.Bio,
	}, nil
}

// GetUserById implements getUserById operation.
//
// Get user by id.
//
// GET /users/{userId}
func (uh *UsersHandler) GetUserById(ctx context.Context, params api.GetUserByIdParams) (api.GetUserByIdRes, error) {
	userInfo, err := uh.usersService.GetUserById(ctx, string(params.UserId))
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			return &api.UserNotFoundResponse{}, nil
		}
		logger.FromCtx(ctx).Error("get user by id", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}
	return &api.User{
		ID:       api.UserId(userInfo.Id),
		Username: userInfo.Username,
		Bio:      userInfo.Bio,
	}, nil
}

// UpdateMe implements updateMe operation.
//
// Update user info.
//
// PUT /users/me
func (uh *UsersHandler) UpdateMe(ctx context.Context, req *api.UserInput) (api.UpdateMeRes, error) {
	userInfo, err := uh.usersService.UpdateMe(ctx, auth.UserIdFromCtx(ctx), dto.UpdateMeInput{
		Username: req.GetUsername(),
		Bio:      req.GetBio(),
	})
	if err != nil {
		logger.FromCtx(ctx).Error("update me", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}
	return &api.User{
		ID:       api.UserId(userInfo.Id),
		Username: userInfo.Username,
		Bio:      userInfo.Bio,
	}, nil
}
