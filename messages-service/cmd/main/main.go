package main

import (
	"context"
	"flag"
	"fmt"
	"go13/messages-service/internal/service"
	"go13/messages-service/internal/config"
	"go13/messages-service/internal/transport/rest"
	"go13/messages-service/internal/transport/rest/handlers"
	repo "go13/messages-service/internal/repo/postgres"
	"go13/pkg/logger"
	"go13/pkg/postgres"
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
		log.Fatal(err)
	}

	defer l.Sync() //nolint:errcheck

	db, err := postgres.Get(ctx, cfg.PostgresCfg)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	repo := repo.NewMessageRepository(db)

	srv := service.NewMessageService(repo)

	messageHandler := handlers.NewMessagesHandler(srv)

	server, err := rest.NewServer(messageHandler, l, cfg.Port)
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
