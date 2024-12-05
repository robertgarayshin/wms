package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/robertgarayshin/wms/pkg/customerrors"

	"github.com/robertgarayshin/wms/internal/entity"
)

// ItemsUseCase - usecase, реализующий интерфейс работы с товарами.
type ItemsUseCase struct {
	itemsRepository ItemsRepo
}

func NewItemsUsecase(i ItemsRepo) ItemsUseCase {
	return ItemsUseCase{
		itemsRepository: i,
	}
}

// CreateItems - создание товаров. Принимает на вход слайс товаров.
func (uc *ItemsUseCase) CreateItems(ctx context.Context, items []entity.Item) error {
	if err := uc.itemsRepository.StoreItems(ctx, items); errors.Is(err, customerrors.ErrNoWarehouse) {
		return err
	} else if err != nil {
		return fmt.Errorf("error store items in db. %w", err)
	}

	return nil
}

// Quantity - Счетчик количества товаров на складе по ID склада. На вход принимает id склада.
func (uc *ItemsUseCase) Quantity(ctx context.Context, id int) (map[string]int, error) {
	quantity, err := uc.itemsRepository.QuantityByWarehouse(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error get quantity by warehouse %d: %w", id, err)
	}

	return quantity, nil
}
