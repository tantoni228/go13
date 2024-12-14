package service

import (
  "context"
  api "go13/pkg/ogen/users-service"
)

type AuthRepo interface {
  CheckToken(ctx context.Context) (api.CheckTokenRes, error)
  SignIn(ctx context.Context, req *api.SignInReq) (api.SignInRes, error)
  SignUp(ctx context.Context, req *api.SignUpReq) (api.SignUpRes, error)
}

type UsersRepo interface {
  ChangePassword(ctx context.Context, req *api.ChangePasswordReq) (api.ChangePasswordRes, error)
  GetMe(ctx context.Context) (api.GetMeRes, error)
  UpdateMe(ctx context.Context, req *api.UserInput) (api.UpdateMeRes, error)
  GetUserById(ctx context.Context, params api.GetUserByIdParams) (api.GetUserByIdRes, error)
}

type UserService struct {
  UserRepo UsersRepo
  AuthRepo AuthRepo
}

func NewUserService(userRepo UsersRepo, authRepo AuthRepo) *UserService {
  return &UserService{userRepo, authRepo}
}

func (s *UserService) CheckToken(ctx context.Context) (api.CheckTokenRes, error) {
  return s.AuthRepo.CheckToken(ctx)
}

func (s *UserService) SignIn(ctx context.Context, req *api.SignInReq) (api.SignInRes, error) {
  return s.AuthRepo.SignIn(ctx, req)
}

func (s *UserService) SignUp(ctx context.Context, req *api.SignUpReq) (api.SignUpRes, error) {
  return s.AuthRepo.SignUp(ctx, req)
}

func (s *UserService) ChangePassword(ctx context.Context, req *api.ChangePasswordReq) (api.ChangePasswordRes, error) {
  return s.UserRepo.ChangePassword(ctx, req)
}

func (s *UserService) GetMe(ctx context.Context) (api.GetMeRes, error) {
  return s.UserRepo.GetMe(ctx)
}

func (s *UserService) GetUserById(ctx context.Context, params api.GetUserByIdParams) (api.GetUserByIdRes, error) {
  return s.UserRepo.GetUserById(ctx, params)
}