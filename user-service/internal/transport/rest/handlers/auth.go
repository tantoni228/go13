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

type AuthService interface {
	SignUp(ctx context.Context, input dto.SignUpInput) error
	SignIn(ctx context.Context, input dto.SignInInput) (dto.SignInRes, error)
	CheckToken(ctx context.Context, tokenString string) error
}

type AuthHandler struct {
	authService AuthService
}

func NewAuthHandler(authService AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// SignUp implements signUp operation.
//
// Sign up to messanger.
//
// POST /auth/sign-up
func (ah *AuthHandler) SignUp(ctx context.Context, req *api.SignUpReq) (api.SignUpRes, error) {
	err := ah.authService.SignUp(ctx, dto.SignUpInput{
		Username: req.GetUsername(),
		Email:    string(req.GetEmail()),
		Password: string(req.GetPassword()),
	})
	if err != nil {
		if errors.Is(err, models.ErrEmailIsTaken) {
			return &api.SignUpConflict{}, nil
		}
		logger.FromCtx(ctx).Error("sign up", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}

	return &api.SignUpNoContent{}, nil
}

// SignIn implements signIn operation.
//
// Sign in to messanger.
//
// POST /auth/sign-in
func (ah *AuthHandler) SignIn(ctx context.Context, req *api.SignInReq) (api.SignInRes, error) {
	res, err := ah.authService.SignIn(ctx, dto.SignInInput{
		Email:    string(req.GetEmail()),
		Password: string(req.GetPassword()),
	})
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			return &api.SignInUnauthorized{}, nil
		}
		logger.FromCtx(ctx).Error("sign in", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}

	return &api.SignInResponse{Token: res.Token}, nil
}

// CheckToken implements checkToken operation.
//
// Check token.
//
// GET /auth/check
func (ah *AuthHandler) CheckToken(ctx context.Context) (api.CheckTokenRes, error) {
	err := ah.authService.CheckToken(ctx, auth.TokenFromCtx(ctx))
	if err != nil {
		if errors.Is(err, models.ErrInvalidToken) {
			return &api.UnauthenticatedResponse{}, nil
		}
		logger.FromCtx(ctx).Error("check token", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}

	return &api.CheckTokenNoContent{}, nil
}
