package usecase

import (
	"context"
	"fmt"

	"github.com/robertgarayshin/wms/internal/entity"
)

// WarehousesUsecase - usecase, реализующий интерфейс работы со складами.
type WarehousesUsecase struct {
	warehousesRepository WarehousesRepo
}

// WarehouseCreate - метод создания нового склада. На вход принимает сущность нового склада.
func (w *WarehousesUsecase) WarehouseCreate(ctx context.Context, warehouse entity.Warehouse) error {
	if err := w.warehousesRepository.CreateWarehouse(ctx, warehouse); err != nil {
		return fmt.Errorf("error creating warehouse: %w", err)
	}

	return nil
}

func NewWarehousesUsecase(w WarehousesRepo) WarehousesUsecase {
	return WarehousesUsecase{
		warehousesRepository: w,
	}
}
