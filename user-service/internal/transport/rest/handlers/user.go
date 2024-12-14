package handlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	// "go13/messages-service/internal/models"
	"go13/user-service/internal/models"
	"go13/user-service/internal/service"
	api "go13/pkg/ogen/users-service"

	"github.com/lib/pq"
)

type AuthRepo interface {
	CheckToken(ctx context.Context) (api.CheckTokenRes, error)
	SignIn(ctx context.Context, req *SignInReq) (api.SignInRes, error)
	SignUp(ctx context.Context, req *SignUpReq) (api.SignUpRes, error) //Update Bio забыл
  }
  
  type UsersRepo interface {
	ChangePassword(ctx context.Context, req *ChangePasswordReq) (api.ChangePasswordRes, error)
	GetMe(ctx context.Context) (api.GetMeRes, error)
	UpdateMe(ctx context.Context, req *UserInput) (ape.UpdateMeRes, error)
	GetUserById(ctx context.Context, params GetUserByIdParams) (api.GetUserByIdRes, error)
  }
  
  type UserService struct {
	UserRepo UsersRepo
	AuthRepo AuthRepo
  }
  

type UserHandler struct {
	service UserService
}

func NewUserHandler(srv *service.UserService) *UserHandler {
	return &UserHandler{service: srv}
  }
  
  func (uh *UserHandler) SignUp(ctx context.Context, req *SignUpReq) (api.SignUpRes, error) {
	UserInfo := models.User{
	  Username: req.GetUsername(),
	  Email:    req.GetEmail(),
	  Password: req.GetPassword(),
	}
	user, err := uh.service.SignUp(ctx, req)
	if err != nil {
	  if errors.Is(err, models.ErrEmailAlreadyExisting) || errors.Is(err, models.ErrUsernameAlreadyExisting) {
		return &api.SignUpConflict{}, nil
	  }
	  logger.FromCtx(ctx).Error("signing up", zap.Error(err))
	  return &api.InternalErrorResponse{}, nil
	}
	return &api.SignUpNoContent{}, nil
  }
  
  func (uh *UserHandler) SignIn(ctx context.Context, req *SignInReq) (api.SignInRes, error) {
	LogInInfo := models.User{
	  Email:    req.GetEmail(),
	  Password: req.GetPassword(),
	}
	user, err := uh.service.SignIn(ctx, req)
	if err != nil {
	  if errors.Is(err, models.ErrUsernameNotFound) || errors.Is(err, models.ErrPasswordIsIncorrect) {
		return &api.SignInUnauthorized{}, nil
	  }
	  logger.FromCtx(ctx).Error("signing in", zap.Error(err))
	  return &api.InternalErrorResponse{}, nil
	}
	return user, nil // или ваш объект SignInRes
  }
  
  func (uh *UserHandler) ChangePassword(ctx context.Context, req *ChangePasswordReq) (api.ChangePasswordRes, error) {
	resp, err := uh.service.ChangePassword(ctx, &api.ChangePasswordReq{
	  OldPassword: req.GetOldPassword(),
	  NewPassword: req.GetNewPassword(),
	})
	if err != nil {
	  logger.FromCtx(ctx).Error("changing password", zap.Error(err))
	  return &api.InternalErrorResponse{}, nil
	}
	return resp, nil
  }