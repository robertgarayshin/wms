package repo

import (
	"context"
	"fmt"

	"github.com/robertgarayshin/wms/internal/entity"
	"github.com/robertgarayshin/wms/pkg/postgres"
)

type WarehousesRepo struct {
	*postgres.Postgres
}

func (r *WarehousesRepo) CreateWarehouse(ctx context.Context, warehouse entity.Warehouse) error {
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error starting transaction. %w", err)
	}

	stmt, args, err := r.Builder.Insert("warehouses").
		Columns("name", "is_available").
		Values(warehouse.Name, warehouse.Availability).ToSql()
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return fmt.Errorf("transaction already closed. %w", rollbackErr)
		}

		return fmt.Errorf("error creating statement. %w", err)
	}

	_, err = tx.Exec(ctx, stmt, args...)
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return fmt.Errorf("transaction already closed. %w", rollbackErr)
		}

		return fmt.Errorf("error creating warehouse. %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return fmt.Errorf("transaction already closed. %w", rollbackErr)
		}

		return fmt.Errorf("error commiting transaction. %w", err)
	}

	return nil
}

func NewWarehousesRepo(p *postgres.Postgres) *WarehousesRepo {
	return &WarehousesRepo{p}
}
