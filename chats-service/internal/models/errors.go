package models

import "errors"

var (
	ErrChatNotFound       = errors.New("chat not found")
	ErrRoleNotFound       = errors.New("role not found")
	ErrMemberNotFound     = errors.New("member not found")
	ErrMessageNotFound    = errors.New("message not found")
	ErrAccessForbidden    = errors.New("access forbidden")
	ErrRoleAlreadyExists  = errors.New("role already exists")
	ErrUserAlreadyInChat  = errors.New("user already in chat")
	ErrInvalidJoinCode    = errors.New("invalid join code")
	ErrUserIsBanned       = errors.New("user is banned")
	ErrEndpointNotFound   = errors.New("endpoint not found")
	ErrInvalidRouteParams = errors.New("invalid route params")
)
