package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/mo0Oonnn/password-generator/internal/config"
	"github.com/mo0Oonnn/password-generator/internal/http-server/routes"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "production"
)

func main() {
	cfg := config.MustLoad()

	logger := setupLogger(cfg.Environment)

	logger.Info("starting server", slog.String("address", cfg.Address), slog.String("environment", cfg.Environment))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      routes.Routes(logger),
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Error(fmt.Sprintf("failed to start server: %v", err))
	}

	logger.Error("server stopped")
}

func setupLogger(environment string) *slog.Logger {
	var logger *slog.Logger
	logsFile, err := os.Create("logs.txt")
	if err != nil {
		log.Fatalf("error creating logs file: %v", err)
	}

	if environment == envLocal {
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	} else if environment == envDev || environment == envProd {
		logger = slog.New(
			slog.NewJSONHandler(logsFile, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}

	return logger
}
