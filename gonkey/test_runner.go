package gonkey

import (
	"database/sql"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lamoda/gonkey/fixtures"
	"github.com/lamoda/gonkey/runner"

	"github.com/robertgarayshin/wms/config"
	v1 "github.com/robertgarayshin/wms/internal/controller/http/v1"
	repo2 "github.com/robertgarayshin/wms/internal/infrastructure/repo"
	"github.com/robertgarayshin/wms/internal/usecase"
	"github.com/robertgarayshin/wms/pkg/logger"
	"github.com/robertgarayshin/wms/pkg/postgres"
)

type TestRunner struct {
	caseDirs []string
}

func (r TestRunner) Run(t *testing.T) {
	cfg, err := config.NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	gin.SetMode(gin.TestMode)

	l := logger.New(cfg.Log.Level)

	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		t.Fatal(fmt.Errorf("error creating postgres repository. %w", err))
	}
	defer pg.Close()

	// создайте экземпляр сервера вашего приложения
	itemsUsecase := usecase.NewItemsUsecase(
		repo2.NewItemsRepository(pg),
	)

	reservationsUsecase := usecase.NewReservationsUsecase(
		repo2.NewReservationRepo(pg),
	)

	warehousesUsecase := usecase.NewWarehousesUsecase(
		repo2.NewWarehousesRepo(pg),
	)

	handler := gin.New()
	v1.NewRouter(handler,
		l,
		itemsUsecase,
		reservationsUsecase,
		warehousesUsecase,
	)
	db, _ := sql.Open("postgres", cfg.PG.URL)

	srv := httptest.NewServer(handler)
	defer srv.Close()

	for _, dir := range r.caseDirs {
		runner.RunWithTesting(t, &runner.RunWithTestingParams{
			Server:      srv,
			TestsDir:    dir,
			DB:          db,
			DbType:      fixtures.Postgres,
			FixturesDir: "./fixtures",
		})
	}
}
