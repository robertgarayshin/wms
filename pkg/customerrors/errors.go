package customerrors

import "errors"

var (
	// ErrWarehouseUnavailable - ошибка, склад недоступен.
	ErrWarehouseUnavailable = errors.New("warehouse is unavailable")
	// ErrNoReservation - ошибка, товар не зарезервирован.
	ErrNoReservation = errors.New("no reservation found")
	//	ErrNoWarehouse - ошибка, склада не существует
	ErrNoWarehouse = errors.New("no warehouse found")
)
