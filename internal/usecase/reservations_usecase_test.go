package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/robertgarayshin/wms/internal/usecase/mock"
	"github.com/robertgarayshin/wms/pkg/customerrors"
)

type reservationsSuite struct {
	reservationsRepo ReservationsRepo
}

func defaultReservationsSuite(t *testing.T) *reservationsSuite {
	ctrl := gomock.NewController(t)

	return &reservationsSuite{
		reservationsRepo: mock.NewMockReservationsRepo(ctrl),
	}
}

func TestNewReservationsUsecase(t *testing.T) {
	tests := []struct {
		name      string
		prepareFn func(suite *reservationsSuite)
		wantErr   bool
	}{
		{
			name: "success",
			prepareFn: func(suite *reservationsSuite) {
			},
			wantErr: false,
		},
		{
			name: "reservationsRepo is nil",
			prepareFn: func(suite *reservationsSuite) {
				suite.reservationsRepo = nil
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suit := defaultReservationsSuite(t)
			if tt.prepareFn != nil {
				tt.prepareFn(suit)
			}
			NewReservationsUsecase(suit.reservationsRepo)
		})
	}
}

func TestReservationsUsecase_Reserve(t *testing.T) {
	type args struct {
		ctx context.Context
		ids []string
	}
	err := errors.New("dummy error")
	tests := []struct {
		name        string
		prepareFn   func(suite *reservationsSuite)
		wantErr     bool
		wantedError error
		args        args
	}{
		{
			name: "success",
			prepareFn: func(suite *reservationsSuite) {
				suite.reservationsRepo.(*mock.MockReservationsRepo).EXPECT().CreateReservation(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "warehouse unavailable error",
			prepareFn: func(suite *reservationsSuite) {
				suite.reservationsRepo.(*mock.MockReservationsRepo).EXPECT().CreateReservation(gomock.Any(), gomock.Any()).
					Return(customerrors.ErrWarehouseUnavailable)
			},
			wantErr:     true,
			wantedError: customerrors.ErrWarehouseUnavailable,
		},
		{
			name: "other error",
			prepareFn: func(suite *reservationsSuite) {
				suite.reservationsRepo.(*mock.MockReservationsRepo).EXPECT().CreateReservation(gomock.Any(), gomock.Any()).
					Return(err)
			},
			wantErr:     true,
			wantedError: err,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suit := defaultReservationsSuite(t)
			if tt.prepareFn != nil {
				tt.prepareFn(suit)
			}

			uc := NewReservationsUsecase(suit.reservationsRepo)

			err := uc.Reserve(tt.args.ctx, tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reserve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !errors.Is(err, tt.wantedError) {
				t.Errorf("Reserve() error = %v, wantErr %v", err, tt.wantedError)
			}
		})
	}
}

func TestItemsUseCase_CancelReservation(t *testing.T) {
	type args struct {
		ctx context.Context
		ids []string
	}
	err := errors.New("dummy error")
	tests := []struct {
		name      string
		prepareFn func(suite *reservationsSuite)
		wantErr   bool
		wantedErr error
		args      args
	}{
		{
			name: "success",
			prepareFn: func(suite *reservationsSuite) {
				suite.reservationsRepo.(*mock.MockReservationsRepo).EXPECT().DeleteReservation(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "error no reservation",
			prepareFn: func(suite *reservationsSuite) {
				suite.reservationsRepo.(*mock.MockReservationsRepo).EXPECT().DeleteReservation(gomock.Any(), gomock.Any()).Return(customerrors.ErrNoReservation)
			},
			wantErr:   true,
			wantedErr: customerrors.ErrNoReservation,
		},
		{
			name: "error",
			prepareFn: func(suite *reservationsSuite) {
				suite.reservationsRepo.(*mock.MockReservationsRepo).EXPECT().DeleteReservation(gomock.Any(), gomock.Any()).Return(err)
			},
			wantErr:   true,
			wantedErr: err,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suit := defaultReservationsSuite(t)
			if tt.prepareFn != nil {
				tt.prepareFn(suit)
			}

			uc := NewReservationsUsecase(suit.reservationsRepo)

			err := uc.CancelReservation(tt.args.ctx, tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("Quantity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
