package main

import (
	"context"
	"flag"
	"go13/chats-service/internal/config"
	pgrepo "go13/chats-service/internal/repo/postgres"
	"go13/chats-service/internal/service"
	"go13/chats-service/internal/transport/rest"
	"go13/chats-service/internal/transport/rest/handlers"
	"go13/pkg/logger"
	"go13/pkg/postgres"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"go.uber.org/zap"
)

var (
	shutdownTimeout = time.Second * 10
)

var (
	configPath = flag.String("config", "configs/example.yml", "path to config file")
)

func main() {
	flag.Parse()

	ctx := context.Background()

	cfg, err := config.Get(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	l, err := logger.Get(cfg.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	defer l.Sync() //nolint:errcheck

	pg, err := postgres.Get(ctx, cfg.PostgresCfg)
	if err != nil {
		l.Fatal("postgres.Get", zap.Error(err))
	}

	trManager, err := manager.New(trmsqlx.NewDefaultFactory(pg.DB))
	if err != nil {
		l.Fatal("manager.Get", zap.Error(err))
	}

	rolesRepo := pgrepo.NewRolesRepo(pg)
	chatsRepo := pgrepo.NewChatsRepo(pg)
	membersRepo := pgrepo.NewMembersRepo(pg)

	chatsService := service.NewChatsService(chatsRepo, rolesRepo, membersRepo, trManager)

	chatsHandler := handlers.NewChatsHandler(chatsService)
	rolesHandler := handlers.NewRolesHandler()

	server, err := rest.NewServer(chatsHandler, rolesHandler, l, cfg.Port)
	if err != nil {
		l.Fatal("rest.NewServer", zap.Error(err))
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	defer func() {
		l.Info("server stopped")
	}()

	go func() {
		l.Info("starting rest server", zap.Int("port", cfg.Port))
		if err := server.Run(ctx); err != nil && err != http.ErrServerClosed {
			l.Fatal("server.Run", zap.Error(err))
		}
	}()

	<-sigCh
	ctx, cancel := context.WithTimeout(ctx, shutdownTimeout)
	defer cancel()

	l.Info("gracefully shutting down")
	if err := server.Shutdown(ctx); err != nil {
		l.Fatal("server.Shutdown", zap.Error(err))
	}

}
