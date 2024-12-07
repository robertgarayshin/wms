package usecase

import (
	"context"

	"github.com/robertgarayshin/wms/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks.go -package=usecase
type (
	// Items - интерфейс сервиса товаров. Позволяет создавать товары и обновлять их.
	Items interface {
		CreateItems(ctx context.Context, items []entity.Item) error
		Quantity(context.Context, int) (map[string]int, error)
	}

	// Reservations - интерфейс сервиса резерваций. Позволяет создавать и удалять резервации товаров.
	Reservations interface {
		Reserve(context.Context, []string) error
		CancelReservation(context.Context, []string) error
	}

	// Warehouse - интерфейс сервиса складов. Позволяет создавать склады для товаров.
	Warehouse interface {
		WarehouseCreate(ctx context.Context, warehouse entity.Warehouse) error
	}

	// ItemsRepo - интерфейс репозитория товаров. Используется для создания,
	// изменения товаров или для получения их количества по складу.
	ItemsRepo interface {
		StoreItems(context.Context, []entity.Item) error
		QuantityByWarehouse(context.Context, int) (map[string]int, error)
	}

	// ReservationsRepo - интерфейс репозитория резерваций. Используется для
	// создания и удаления резерваций товаров по их уникальному коду.
	ReservationsRepo interface {
		CreateReservation(context.Context, []string) error
		DeleteReservation(context.Context, []string) error
	}

	// WarehousesRepo - интерфейс репозитория складов. Используется для создания
	// и удаления складов.
	WarehousesRepo interface {
		CreateWarehouse(context.Context, entity.Warehouse) error
	}
)
