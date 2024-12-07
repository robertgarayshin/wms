package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/robertgarayshin/wms/internal/entity"
	"github.com/robertgarayshin/wms/internal/usecase/mock"
)

type warehousesSuite struct {
	warehousesRepository WarehousesRepo
}

func defaultWarehousesSuite(t *testing.T) *warehousesSuite {
	ctrl := gomock.NewController(t)

	return &warehousesSuite{
		warehousesRepository: mock.NewMockWarehousesRepo(ctrl),
	}
}

func TestNewWarehousesUsecase(t *testing.T) {
	tests := []struct {
		name      string
		prepareFn func(suite *warehousesSuite)
		wantErr   bool
	}{
		{
			name: "success",
			prepareFn: func(suite *warehousesSuite) {
			},
			wantErr: false,
		},
		{
			name: "warehousesRepository is nil",
			prepareFn: func(suite *warehousesSuite) {
				suite.warehousesRepository = nil
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suit := defaultWarehousesSuite(t)
			if tt.prepareFn != nil {
				tt.prepareFn(suit)
			}
			NewWarehousesUsecase(suit.warehousesRepository)
		})
	}
}

func TestWarehousesUsecase_WarehouseCreate(t *testing.T) {
	type args struct {
		ctx       context.Context
		warehouse entity.Warehouse
	}
	tests := []struct {
		name      string
		prepareFn func(suite *warehousesSuite)
		wantErr   bool
		args      args
	}{
		{
			name: "success",
			prepareFn: func(suite *warehousesSuite) {
				suite.warehousesRepository.(*mock.MockWarehousesRepo).EXPECT().CreateWarehouse(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "error",
			prepareFn: func(suite *warehousesSuite) {
				suite.warehousesRepository.(*mock.MockWarehousesRepo).EXPECT().CreateWarehouse(gomock.Any(), gomock.Any()).
					Return(errors.New(""))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suit := defaultWarehousesSuite(t)
			if tt.prepareFn != nil {
				tt.prepareFn(suit)
			}

			uc := NewWarehousesUsecase(suit.warehousesRepository)

			err := uc.WarehouseCreate(tt.args.ctx, tt.args.warehouse)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateWarehouse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
