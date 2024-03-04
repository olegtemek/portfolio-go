package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/olegtemek/portfolio-go/internal/config"
	"github.com/olegtemek/portfolio-go/internal/logger"
	"github.com/olegtemek/portfolio-go/internal/repository"
	"github.com/olegtemek/portfolio-go/internal/server/rest"
	"github.com/olegtemek/portfolio-go/internal/service"
	"github.com/olegtemek/portfolio-go/internal/storage"
)

// @title Porfolio-go
// @version 1.0

// @Host olegtemek.site
// @BasePath /

// @securitydefinitions.basic BasicAuth
// @in header
// @name BasicAuth
func main() {

	cfg, err := config.New()

	if err != nil {
		panic("Cannot get .env file")
	}

	log := logger.New(cfg)

	db, err := storage.NewPostgresConnect(log, cfg)
	if err != nil {
		panic("Cannot connect to database")
	}

	repositories := repository.New(log, db)
	services := service.NewService(log, repositories)
	restHandler := rest.NewHandler(log, cfg, services)
	server := restHandler.Init()

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("ERROR", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Info("Server is shutting down", slog.String("sig", sig.String()))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		log.Info("Could not gracefully shutdown the server", err)
	}
	log.Info("Server was stopped")
}
