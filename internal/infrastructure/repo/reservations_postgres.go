package repo

import (
	"context"
	"fmt"

	"github.com/robertgarayshin/wms/pkg/customerrors"
	"github.com/robertgarayshin/wms/pkg/postgres"
)

type ReservationsRepo struct {
	*postgres.Postgres
}

func (r *ReservationsRepo) CreateReservation(ctx context.Context, ids []string) error {
	reservations := r.reservedItemsCount(ids)

	for id, count := range reservations {
		tx, err := r.Pool.Begin(ctx)
		if err != nil {
			return fmt.Errorf("error begining transaction. %w", err)
		}
		// Запись склада Unstored (id = 0) для товаров, которые создаются при создании резервации
		/*
			Usecase: товар добавлен в предварительный резерв, но при этом еще не прибыл на склад.
			Записи об этом товаре нет в таблице товаров.
			Создается резервация -> уникальный код товара создает дефолтный товар (все поля пустые, кроме кода)
			Товар помещается на склад unstored.
			Количество товара на складе = 0 - количество товара в резервации (уменьшается при следующей резервации)
			Товар прибывает на склад: запись о нем обновляется, указывается реальный склад, добавляется вся информация
				количество товара = текущее количество + количество прибывшего
			Резервация товара сохраняется, новое количество товара учитывает зарезервированный
		*/
		var availability bool

		itemCreateStatement := `INSERT INTO items(unique_code) VALUES ($1) ON CONFLICT DO NOTHING`

		_, err = tx.Exec(ctx, itemCreateStatement, id)
		if err != nil {
			if rollbackError := tx.Rollback(ctx); rollbackError != nil {
				return fmt.Errorf("transaction already closed. %w", rollbackError)
			}

			return fmt.Errorf("error create item. %w", err)
		}

		checkWarehouseStatement := `SELECT w.is_available FROM items 
   										JOIN warehouses w ON w.id = items.warehouse_id
										WHERE unique_code = $1`
		row := tx.QueryRow(ctx, checkWarehouseStatement, id)

		if err = row.Scan(&availability); err != nil {
			if rollbackError := tx.Rollback(ctx); rollbackError != nil {
				return fmt.Errorf("transaction already closed. %w", rollbackError)
			}

			return fmt.Errorf("error checking warehouse availability. %w", err)
		}

		if !availability {
			if rollbackError := tx.Rollback(ctx); rollbackError != nil {
				return fmt.Errorf("transaction already closed. %w", rollbackError)
			}

			return customerrors.ErrWarehouseUnavailable
		}

		reservationCreateStatement := `UPDATE items 
											SET reserved = reserved + $1,
											    quantity = quantity - $1
										WHERE unique_code = $2`

		_, err = tx.Exec(ctx, reservationCreateStatement, count, id)
		if err != nil {
			if rollbackError := tx.Rollback(ctx); rollbackError != nil {
				return fmt.Errorf("transaction already closed. %w", rollbackError)
			}

			return fmt.Errorf("error executing insert reservation. %w", err)
		}

		if err = tx.Commit(ctx); err != nil {
			if rollbackError := tx.Rollback(ctx); rollbackError != nil {
				return fmt.Errorf("transaction already closed. %w", rollbackError)
			}

			return fmt.Errorf("error commiting transaction. %w", err)
		}
	}

	return nil
}

func (r *ReservationsRepo) DeleteReservation(ctx context.Context, ids []string) error {
	deleteReservations := r.reservedItemsCount(ids)

	for id, count := range deleteReservations {
		tx, err := r.Pool.Begin(ctx)
		if err != nil {
			return fmt.Errorf("error begining transaction. %w", err)
		}

		var reserved int
		var availability bool

		checkWarehouseStatement := `SELECT w.is_available FROM items 
   										JOIN warehouses w ON w.id = items.warehouse_id
										WHERE unique_code = $1`
		row := tx.QueryRow(ctx, checkWarehouseStatement, id)

		if err = row.Scan(&availability); err != nil {
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				return fmt.Errorf("transaction already closed. %w", rollbackErr)
			}

			return fmt.Errorf("error checking warehouse availability. %w", err)
		}

		if !availability {
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				return fmt.Errorf("transaction already closed. %w", rollbackErr)
			}

			return customerrors.ErrWarehouseUnavailable
		}

		reservationsCheck, args, err := r.Builder.Select("reserved").
			From("items").Where("unique_code = ?", id).
			ToSql()
		if err != nil {
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				return fmt.Errorf("transaction already closed. %w", rollbackErr)
			}

			return fmt.Errorf("error building statement. %w", err)
		}

		err = tx.QueryRow(ctx, reservationsCheck, args...).Scan(&reserved)
		if reserved == 0 {
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				return fmt.Errorf("transaction already closed. %w", rollbackErr)
			}

			return customerrors.ErrNoReservation
		} else if err != nil {
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				return fmt.Errorf("transaction already closed. %w", rollbackErr)
			}

			return fmt.Errorf("error scanning item reservations. %w", err)
		}

		reservationDeleteStatement := `UPDATE items
											SET reserved = reserved - $1,
											    quantity = quantity + $1
										WHERE unique_code = $2`

		_, err = tx.Exec(ctx, reservationDeleteStatement, count, id)
		if err != nil {
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				return fmt.Errorf("transaction already closed. %w", rollbackErr)
			}

			return fmt.Errorf("error delete reservation. %w", err)
		}

		if err = tx.Commit(ctx); err != nil {
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				return fmt.Errorf("transaction already closed. %w", rollbackErr)
			}

			return fmt.Errorf("error commiting transaction. %w", err)
		}
	}

	return nil
}

func (r *ReservationsRepo) reservedItemsCount(ids []string) map[string]int {
	res := make(map[string]int, len(ids))

	for _, id := range ids {
		res[id]++
	}

	return res
}

func NewReservationRepo(p *postgres.Postgres) *ReservationsRepo {
	return &ReservationsRepo{p}
}
