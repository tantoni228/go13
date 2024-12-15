package models

import "errors"

var (
	ErrEmailAlreadyExisting    = errors.New("email already existing")
	ErrUsernameAlreadyExisting = errors.New("username already existing")
	ErrPasswordAlreadyExisting = errors.New("password already existing")
	ErrEmailNotFound           = errors.New("user's email not found")
	ErrIDNotFound              = errors.New("user's id not found")
	ErrPasswordIsIncorrect     = errors.New("password is incorrect")
	ErrUsernameNotFound        = errors.New("username not found")
)
