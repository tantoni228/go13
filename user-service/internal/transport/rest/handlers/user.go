package handlers

import (
	"context"
	"errors"

	"go13/user-service/internal/service"
	"go13/pkg/logger"
	api "go13/pkg/ogen/users-service"
	"go13/user-service/internal/models"

	"go.uber.org/zap"
)


type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(srv *service.UserService) *UserHandler {
	return &UserHandler{service: srv}
}

func (uh *UserHandler) SignUp(ctx context.Context, req *api.SignUpReq) (api.SignUpRes, error) {
	UserInfo := models.User{
		Username: req.GetUsername(),
		Email:    models.Email(req.GetEmail()),
		Password: models.Password(req.GetUsername()),
	}

	_, err := uh.service.SignUp(ctx, &api.SignUpReq{Email: api.Email(UserInfo.Email),
		Username: string(UserInfo.Username), Password: api.Password(UserInfo.Password)})
	if err != nil {
		if errors.Is(err, models.ErrEmailAlreadyExisting) || errors.Is(err, models.ErrUsernameAlreadyExisting) {
			return &api.SignUpConflict{}, nil
		}
		logger.FromCtx(ctx).Error("signing up", zap.Error(err))
		return &api.InternalErrorResponse{}, err
	}
	return &api.SignUpNoContent{}, nil
}

func (uh *UserHandler) SignIn(ctx context.Context, req *api.SignInReq) (api.SignInRes, error) {
	LogInInfo := models.User{
		Email:    models.Email(req.GetEmail()),
		Password: models.Password(req.GetPassword()),
	}

	user, err := uh.service.SignIn(ctx, &api.SignInReq{Email: api.Email(LogInInfo.Email),
		Password: api.Password(LogInInfo.Password)})
	if err != nil {
		if errors.Is(err, models.ErrUsernameNotFound) || errors.Is(err, models.ErrPasswordIsIncorrect) {
			return &api.SignInUnauthorized{}, nil
		}
		logger.FromCtx(ctx).Error("signing in", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}

	return user, nil
}

func (uh *UserHandler) ChangePassword(ctx context.Context, req *api.ChangePasswordReq) (api.ChangePasswordRes, error) {
	resp, err := uh.service.UserRepo.ChangePassword(ctx, &api.ChangePasswordReq{
		OldPassword: req.GetOldPassword(),
		NewPassword: req.GetNewPassword(),
	})
	if err != nil {
		logger.FromCtx(ctx).Error("changing password", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}
	return resp, nil
}

// func (uh *UserHandler) GetUserById(ctx context.Context, params api.GetUserByIdParams) (api.GetUserByIdRes, error) {
// 	resp, err := uh.service.UserRepo.GetUserById(ctx, params)
// 	if err != nil {
// 		if errors.Is(err, models.ErrIDNotFound) {
// 			return &api.UserNotFoundResponse{}, nil // используем правильный тип ответа
// 		}
// 		logger.FromCtx(ctx).Error("getting user by id", zap.Error(err))
// 		return &api.InternalErrorResponse{}, nil
// 	}

// 	return resp, nil
// }

// CheckToken implements api.Handler.
func (h *UserHandler) CheckToken(ctx context.Context) (api.CheckTokenRes, error) {
	panic("unimplemented")
}

// GetMe implements api.Handler.
func (h *UserHandler) GetMe(ctx context.Context) (api.GetMeRes, error) {
	panic("unimplemented")
}

// GetUserById implements api.Handler.
func (h *UserHandler) GetUserById(ctx context.Context, params api.GetUserByIdParams) (api.GetUserByIdRes, error) {
	panic("unimplemented")
}

// UpdateMe implements api.Handler.
func (h *UserHandler) UpdateMe(ctx context.Context, req *api.UserInput) (api.UpdateMeRes, error) {
	panic("unimplemented")
}
