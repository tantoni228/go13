package models

import "errors"

var (
	ErrEmailIsTaken       = errors.New("email is taken")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidToken       = errors.New("invalid token")
)
