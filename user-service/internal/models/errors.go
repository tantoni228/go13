package errors

import "errors"

var (
	ErrEmailAlreadyExisting = errors.New("email already existing")
	ErrUsernameAlreadyExisting = errors.New("username already existing")
	ErrPasswordAlreadyExisting = errors.New("password already existing")
	ErrUsernameNotFound = errors.New("username not found")
	ErrIDNotFound = errors.New("user id not found")
)