package auth

import (
	"context"
	"errors"
	api "go13/pkg/ogen/chats-service"

	"github.com/golang-jwt/jwt/v5"
)

type userIdCtxKey struct{}

type jwtClaims struct {
	jwt.RegisteredClaims
	UserId string `json:"user_id"`
}

// SecurityHandler implements api.SecurityHandler.
// HandleBearerAuth parses token and sets userId to request context.
type SecurityHandler struct{}

func NewSecurityHandler() *SecurityHandler {
	return &SecurityHandler{}
}

func (s *SecurityHandler) HandleBearerAuth(ctx context.Context, _ api.OperationName, auth api.BearerAuth) (context.Context, error) {
	var claims jwtClaims
	_, err := jwt.ParseWithClaims(auth.GetToken(), &claims, nil)
	if err != nil && !errors.Is(err, jwt.ErrTokenUnverifiable) {
		return ctx, err
	}

	userId := claims.UserId
	ctx = context.WithValue(ctx, userIdCtxKey{}, userId)
	return ctx, nil
}

// UserIdFromCtx returns userId associated with context.
// If no userId is associated, the empty string is returned.
func UserIdFromCtx(ctx context.Context) string {
	return ctx.Value(userIdCtxKey{}).(string)
}
