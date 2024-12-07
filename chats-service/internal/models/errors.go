package models

import "errors"

var (
	ErrChatNotFound       = errors.New("chat not found")
	ErrRoleNotFound       = errors.New("role not found")
	ErrChatOrRoleNotFound = errors.New("chat or role not found")
	ErrRoleAlreadyExists  = errors.New("role already exists")
	ErrUserAlreadyInChat  = errors.New("user already in chat")
)
