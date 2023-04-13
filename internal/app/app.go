// Package app configures and runs application.
package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/HUSTtoKTH/lintserver/config"
	v1 "github.com/HUSTtoKTH/lintserver/internal/controller/http/v1"
	"github.com/HUSTtoKTH/lintserver/internal/usecase"
	"github.com/HUSTtoKTH/lintserver/internal/usecase/repo"
	"github.com/HUSTtoKTH/lintserver/internal/usecase/webapi"
	"github.com/HUSTtoKTH/lintserver/pkg/httpserver"
	"github.com/HUSTtoKTH/lintserver/pkg/logger"
	"github.com/HUSTtoKTH/lintserver/pkg/postgres"
	"github.com/gin-gonic/gin"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	print(cfg.PG.URL)
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()
	err = pg.Pool.Ping(context.Background())
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.Ping: %w", err))
	}

	// Use case
	translationUseCase := usecase.New(
		repo.New(pg),
		webapi.New(),
	)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, translationUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
