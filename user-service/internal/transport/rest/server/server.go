package server

import (
	"context"
	"fmt"
	api "go13/pkg/ogen/users-service"
	"go13/user-service/internal/transport/rest/auth"
	"go13/user-service/internal/transport/rest/handlers"
	"go13/user-service/internal/transport/rest/middlewares"
	"net/http"

	"go.uber.org/zap"
)

type Server struct {
	srv *http.Server
}

type handler struct {
	*handlers.AuthHandler
	*handlers.UsersHandler
}

func NewServer(
	authHandler *handlers.AuthHandler,
	usersHandler *handlers.UsersHandler,
	l *zap.Logger,
	port int,
) (*Server, error) {

	apiSrv, err := api.NewServer(&handler{
		authHandler,
		usersHandler,
	}, auth.NewSecurityHandler())
	if err != nil {
		return nil, err
	}

	handler := middlewares.Apply(
		apiSrv,
		middlewares.LoggerProvider(l),
		middlewares.Logging(),
	)

	srv := &http.Server{
		Handler: handler,
		Addr:    fmt.Sprintf(":%d", port),
	}

	return &Server{srv: srv}, nil
}

func (s *Server) Run(ctx context.Context) error {
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
