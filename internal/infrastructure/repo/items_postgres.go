package repo

import (
	"context"
	"fmt"

	"github.com/robertgarayshin/wms/pkg/customerrors"

	"github.com/robertgarayshin/wms/internal/entity"
	"github.com/robertgarayshin/wms/pkg/postgres"
)

type ItemsRepo struct {
	*postgres.Postgres
}

func (r *ItemsRepo) StoreItems(ctx context.Context, items []entity.Item) error {
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error starting transaction. %w", err)
	}

	for _, item := range items {
		var exists bool
		checkWarehouseStmt := `SELECT EXISTS(SELECT 1 FROM warehouses WHERE id = $1)`
		row := tx.QueryRow(ctx, checkWarehouseStmt, item.WarehouseID)

		err = row.Scan(&exists)
		if err != nil {
			rollbackError := tx.Rollback(ctx)
			if rollbackError != nil {
				return fmt.Errorf("transaction already closed. %w", rollbackError)
			}

			return fmt.Errorf("error checking existence of warehouse with ID %d. %w", item.WarehouseID, err)
		} else if !exists {
			rollbackError := tx.Rollback(ctx)
			if rollbackError != nil {
				return fmt.Errorf("transaction already closed. %w", rollbackError)
			}

			return customerrors.ErrNoWarehouse
		}

		stmt := `INSERT INTO items (unique_code, name, size, quantity, warehouse_id) 
							VALUES ($1, $2, $3, $4, $5)
								ON CONFLICT (unique_code) DO UPDATE
								SET name = $2, 
								    size = $3, 
								    quantity = items.quantity + $4, 
								    warehouse_id = $5
								WHERE items.unique_code = $1`
		_, err = tx.Exec(ctx, stmt, item.UniqueID, item.Name, item.Size, item.Quantity, item.WarehouseID)
		if err != nil {
			rollbackError := tx.Rollback(ctx)
			if rollbackError != nil {
				return fmt.Errorf("transaction already closed. %w", rollbackError)
			}

			return fmt.Errorf("error executing insert item statement. %w", err)
		}
	}

	if err = tx.Commit(ctx); err != nil {
		if rollbackError := tx.Rollback(ctx); rollbackError != nil {
			return fmt.Errorf("transaction already closed. %w", rollbackError)
		}
		return fmt.Errorf("error commiting transaction. %w", err)
	}

	return nil
}

func (r *ItemsRepo) QuantityByWarehouse(ctx context.Context, warehouseID int) (map[string]int, error) {
	res := make(map[string]int)
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	stmt, args, err := r.Builder.Select("unique_code, quantity").
		From("items").
		Where("warehouse_id = ?", warehouseID).ToSql()
	if err != nil {
		rollbackError := tx.Rollback(ctx)
		if rollbackError != nil {
			return nil, fmt.Errorf("transaction already closed. %w", rollbackError)
		}

		return nil, fmt.Errorf("error building query. %w", err)
	}

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		rollbackError := tx.Rollback(ctx)
		if rollbackError != nil {
			return nil, fmt.Errorf("transaction already closed. %w", rollbackError)
		}

		return nil, fmt.Errorf("error executing query. %w", err)
	}

	for rows.Next() {
		var id string
		var quantity int

		err = rows.Scan(&id, &quantity)
		if err != nil {
			rollbackError := tx.Rollback(ctx)
			if rollbackError != nil {
				return nil, fmt.Errorf("transaction already closed. %w", rollbackError)
			}

			return nil, fmt.Errorf("error scanning row value. %w", err)
		}

		res[id] = quantity
	}

	if err = tx.Commit(ctx); err != nil {
		if rollbackError := tx.Rollback(ctx); rollbackError != nil {
			return nil, fmt.Errorf("transaction already closed. %w", rollbackError)
		}

		return nil, fmt.Errorf("error commiting transaction. %w", err)
	}

	return res, nil
}

func NewItemsRepository(pg *postgres.Postgres) *ItemsRepo {
	return &ItemsRepo{pg}
}
