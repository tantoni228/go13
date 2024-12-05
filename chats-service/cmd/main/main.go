package main

import (
	"context"
	"flag"
	"go13/chats-service/internal/config"
	"go13/chats-service/internal/transport/rest"
	"go13/chats-service/internal/transport/rest/handlers"
	"go13/pkg/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
		log.Fatal(l)
	}

	defer l.Sync() //nolint:errcheck

	chatsHandler := handlers.NewChatsHandler()
	rolesHandler := handlers.NewRolesHandler()

	server, err := rest.NewServer(chatsHandler, rolesHandler, l, cfg.Port)
	if err != nil {
		l.Fatal("error while init server", zap.Error(err))
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	defer func() {
		l.Info("server stopped")
	}()

	go func() {
		l.Info("starting rest server", zap.Int("port", cfg.Port))
		if err := server.Run(ctx); err != nil && err != http.ErrServerClosed {
			l.Fatal("error while starting server", zap.Error(err))
		}
	}()

	<-sigCh
	ctx, cancel := context.WithTimeout(ctx, shutdownTimeout)
	defer cancel()

	l.Info("gracefully shutting down")
	if err := server.Shutdown(ctx); err != nil {
		l.Fatal("error while shutdown server", zap.Error(err))
	}

}
