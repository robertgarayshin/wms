package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/robertgarayshin/wms/pkg/customerrors"
)

// ReservationsUsecase - usecase реализующий интерфейс работы с резервациями.
type ReservationsUsecase struct {
	reservationsRepository ReservationsRepo
}

// Reserve - метод для создания резерва товара. На вход принимает слайс уникальных кодов товара.
func (uc *ReservationsUsecase) Reserve(ctx context.Context, ids []string) error {
	err := uc.reservationsRepository.CreateReservation(ctx, ids)
	if errors.Is(err, customerrors.ErrWarehouseUnavailable) {
		return customerrors.ErrWarehouseUnavailable
	} else if err != nil {
		return fmt.Errorf("error create reservation. %w", err)
	}

	return nil
}

// CancelReservation - метод для удаления резервации. На вход принимает слайс уникальных кодов товара.
func (uc *ReservationsUsecase) CancelReservation(ctx context.Context, ids []string) error {
	err := uc.reservationsRepository.DeleteReservation(ctx, ids)
	if errors.Is(err, customerrors.ErrNoReservation) {
		return err
	} else if err != nil {
		return fmt.Errorf("error delete reservation. %w", err)
	}

	return nil
}

func NewReservationsUsecase(reservationsRepository ReservationsRepo) Reservations {
	return &ReservationsUsecase{
		reservationsRepository: reservationsRepository,
	}
}
