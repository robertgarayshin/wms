package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/robertgarayshin/wms/config"
	v1 "github.com/robertgarayshin/wms/internal/controller/http/v1"
	repo2 "github.com/robertgarayshin/wms/internal/infrastructure/repo"
	"github.com/robertgarayshin/wms/internal/usecase"
	"github.com/robertgarayshin/wms/pkg/httpserver"
	"github.com/robertgarayshin/wms/pkg/logger"
	"github.com/robertgarayshin/wms/pkg/postgres"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("error creating postgres repository. %w", err))
	}
	defer pg.Close()

	// Use case
	itemsUsecase := usecase.NewItemsUsecase(
		repo2.NewItemsRepository(pg),
	)

	reservationsUsecase := usecase.NewReservationsUsecase(
		repo2.NewReservationRepo(pg),
	)

	warehousesUsecase := usecase.NewWarehousesUsecase(
		repo2.NewWarehousesRepo(pg),
	)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler,
		l,
		itemsUsecase,
		reservationsUsecase,
		warehousesUsecase,
	)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app server notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app server shutdown: %w", err))
	}
}
