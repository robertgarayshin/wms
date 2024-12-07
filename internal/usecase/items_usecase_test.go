package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/robertgarayshin/wms/internal/entity"
	"github.com/robertgarayshin/wms/internal/usecase/mock"
	"github.com/robertgarayshin/wms/pkg/customerrors"
)

type suite struct {
	itemsRepo ItemsRepo
}

func defaultSuite(t *testing.T) *suite {
	ctrl := gomock.NewController(t)

	return &suite{
		itemsRepo: mock.NewMockItemsRepo(ctrl),
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name      string
		prepareFn func(suite *suite)
		wantErr   bool
	}{
		{
			name: "success",
			prepareFn: func(suite *suite) {
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suit := defaultSuite(t)
			if tt.prepareFn != nil {
				tt.prepareFn(suit)
			}
			NewItemsUsecase(suit.itemsRepo)
		})
	}
}

func TestItemsUseCase_CreateItems(t *testing.T) {
	type args struct {
		ctx   context.Context
		items []entity.Item
	}
	err := errors.New("dummy error")
	tests := []struct {
		name        string
		prepareFn   func(suite *suite)
		wantErr     bool
		wantedError error
		args        args
	}{
		{
			name: "success",
			prepareFn: func(suite *suite) {
				suite.itemsRepo.(*mock.MockItemsRepo).EXPECT().StoreItems(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "no warehouse error",
			prepareFn: func(suite *suite) {
				suite.itemsRepo.(*mock.MockItemsRepo).EXPECT().StoreItems(gomock.Any(), gomock.Any()).
					Return(customerrors.ErrNoWarehouse)
			},
			wantErr:     true,
			wantedError: customerrors.ErrNoWarehouse,
		},
		{
			name: "other error",
			prepareFn: func(suite *suite) {
				suite.itemsRepo.(*mock.MockItemsRepo).EXPECT().StoreItems(gomock.Any(), gomock.Any()).
					Return(err)
			},
			wantErr:     true,
			wantedError: err,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suit := defaultSuite(t)
			if tt.prepareFn != nil {
				tt.prepareFn(suit)
			}

			uc := NewItemsUsecase(suit.itemsRepo)

			err := uc.CreateItems(tt.args.ctx, tt.args.items)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !errors.Is(err, tt.wantedError) {
				t.Errorf("CreateItems() error = %v, wantErr %v", err, tt.wantedError)
			}
		})
	}
}

func TestItemsUseCase_Quantity(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	//err := errors.New("dummy error")
	tests := []struct {
		name      string
		prepareFn func(suite *suite)
		wantErr   bool
		args      args
	}{
		{
			name: "success",
			prepareFn: func(suite *suite) {
				suite.itemsRepo.(*mock.MockItemsRepo).EXPECT().QuantityByWarehouse(gomock.Any(), gomock.Any()).Return(map[string]int{}, nil)
			},
			wantErr: false,
		},
		{
			name: "error",
			prepareFn: func(suite *suite) {
				suite.itemsRepo.(*mock.MockItemsRepo).EXPECT().QuantityByWarehouse(gomock.Any(), gomock.Any()).Return(nil, errors.New("dummy error"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suit := defaultSuite(t)
			if tt.prepareFn != nil {
				tt.prepareFn(suit)
			}

			uc := NewItemsUsecase(suit.itemsRepo)

			_, err := uc.Quantity(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Quantity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
