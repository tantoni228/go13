package rest

import (
	"context"
	"fmt"
	"go13/chats-service/internal/transport/rest/handlers"
	"go13/pkg/auth"
	"go13/pkg/middlewares"
	api "go13/pkg/ogen/chats-service"
	"net/http"

	"go.uber.org/zap"
)

type Server struct {
	srv *http.Server
}

type handler struct {
	*handlers.ChatsHandler
	*handlers.RolesHandler
}

func NewServer(
	chatsHandler *handlers.ChatsHandler,
	rolesHandler *handlers.RolesHandler,
	l *zap.Logger,
	port int,
) (*Server, error) {

	apiSrv, err := api.NewServer(&handler{
		chatsHandler,
		rolesHandler,
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
