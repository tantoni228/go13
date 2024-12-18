package service

import (
	"context"
	"errors"
	"fmt"
	"go13/user-service/internal/dto"
	"go13/user-service/internal/models"
	"go13/user-service/internal/repo"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	tokenExpirationTime = time.Hour * 24 * 30
)

type AuthService struct {
	usersRepo repo.UsersRepo
	jwtSecret string
}

func NewAuthService(usersRepo repo.UsersRepo, jwtSecret string) *AuthService {
	return &AuthService{
		usersRepo: usersRepo,
		jwtSecret: jwtSecret,
	}
}

func (as *AuthService) SignUp(ctx context.Context, input dto.SignUpInput) error {
	op := "AuthService.SignUp"

	user := models.User{Username: input.Username, Email: input.Email}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%s: bcrypt.GenerateFromPassword: %w", op, err)
	}

	user.HashedPassword = string(hashedPassword)

	_, err = as.usersRepo.AddUser(ctx, user)
	if err != nil {
		return fmt.Errorf("%s: usersRepo.AddUser: %w", op, err)
	}

	return nil
}

func (as *AuthService) SignIn(ctx context.Context, input dto.SignInInput) (dto.SignInRes, error) {
	op := "AuthService.SignIn"

	user, err := as.usersRepo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			return dto.SignInRes{}, models.ErrInvalidCredentials
		}
		return dto.SignInRes{}, fmt.Errorf("%s: usersRepo.GetUserByEmail: %w", op, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(input.Password))
	if err != nil {
		return dto.SignInRes{}, models.ErrInvalidCredentials
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":     "user_service",
		"exp":     time.Now().Add(tokenExpirationTime).Unix(),
		"iat":     time.Now().Unix(),
		"nbf":     time.Now().Unix(),
		"user_id": user.Id,
	})

	signedToken, err := token.SignedString([]byte(as.jwtSecret))
	if err != nil {
		return dto.SignInRes{}, fmt.Errorf("%s: token.SignedString: %w", op, err)
	}

	return dto.SignInRes{Token: signedToken}, nil
}

func (as *AuthService) CheckToken(ctx context.Context, tokenString string) error {
	// op := "AuthService.CheckToken"

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(as.jwtSecret), nil
	}, jwt.WithExpirationRequired())
	if err != nil {
		return models.ErrInvalidToken
	}

	if !token.Valid {
		return models.ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return models.ErrInvalidToken
	}

	userIdAny, ok := claims["user_id"]
	if !ok {
		return models.ErrInvalidToken
	}

	userId, ok := userIdAny.(string)
	if !ok {
		return models.ErrInvalidToken
	}

	if _, err := uuid.Parse(userId); err != nil {
		return models.ErrInvalidToken
	}

	return nil
}
