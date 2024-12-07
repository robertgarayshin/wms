package gonkey

import (
	"testing"
)

func TestReserve(t *testing.T) {
	TestRunner{
		caseDirs: []string{"cases/reservations/create"},
	}.Run(t)
}

func TestDeleteReservation(t *testing.T) {
	TestRunner{
		caseDirs: []string{"cases/reservations/delete"},
	}.Run(t)
}
