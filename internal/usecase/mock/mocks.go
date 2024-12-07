// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package usecase is a generated GoMock package.
package mock

import (
	"context"
	"reflect"

	"github.com/golang/mock/gomock"

	"github.com/robertgarayshin/wms/internal/entity"
)

// MockItems is a mock of Items interface.
type MockItems struct {
	ctrl     *gomock.Controller
	recorder *MockItemsMockRecorder
}

// MockItemsMockRecorder is the mock recorder for MockItems.
type MockItemsMockRecorder struct {
	mock *MockItems
}

// NewMockItems creates a new mock instance.
func NewMockItems(ctrl *gomock.Controller) *MockItems {
	mock := &MockItems{ctrl: ctrl}
	mock.recorder = &MockItemsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItems) EXPECT() *MockItemsMockRecorder {
	return m.recorder
}

// CreateItems mocks base method.
func (m *MockItems) CreateItems(ctx context.Context, items []entity.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateItems", ctx, items)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateItems indicates an expected call of CreateItems.
func (mr *MockItemsMockRecorder) CreateItems(ctx, items interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateItems", reflect.TypeOf((*MockItems)(nil).CreateItems), ctx, items)
}

// Quantity mocks base method.
func (m *MockItems) Quantity(arg0 context.Context, arg1 int) (map[string]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Quantity", arg0, arg1)
	ret0, _ := ret[0].(map[string]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Quantity indicates an expected call of Quantity.
func (mr *MockItemsMockRecorder) Quantity(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Quantity", reflect.TypeOf((*MockItems)(nil).Quantity), arg0, arg1)
}

// MockReservations is a mock of Reservations interface.
type MockReservations struct {
	ctrl     *gomock.Controller
	recorder *MockReservationsMockRecorder
}

// MockReservationsMockRecorder is the mock recorder for MockReservations.
type MockReservationsMockRecorder struct {
	mock *MockReservations
}

// NewMockReservations creates a new mock instance.
func NewMockReservations(ctrl *gomock.Controller) *MockReservations {
	mock := &MockReservations{ctrl: ctrl}
	mock.recorder = &MockReservationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReservations) EXPECT() *MockReservationsMockRecorder {
	return m.recorder
}

// CancelReservation mocks base method.
func (m *MockReservations) CancelReservation(arg0 context.Context, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelReservation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelReservation indicates an expected call of CancelReservation.
func (mr *MockReservationsMockRecorder) CancelReservation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelReservation", reflect.TypeOf((*MockReservations)(nil).CancelReservation), arg0, arg1)
}

// Reserve mocks base method.
func (m *MockReservations) Reserve(arg0 context.Context, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reserve", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reserve indicates an expected call of Reserve.
func (mr *MockReservationsMockRecorder) Reserve(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reserve", reflect.TypeOf((*MockReservations)(nil).Reserve), arg0, arg1)
}

// MockWarehouse is a mock of Warehouse interface.
type MockWarehouse struct {
	ctrl     *gomock.Controller
	recorder *MockWarehouseMockRecorder
}

// MockWarehouseMockRecorder is the mock recorder for MockWarehouse.
type MockWarehouseMockRecorder struct {
	mock *MockWarehouse
}

// NewMockWarehouse creates a new mock instance.
func NewMockWarehouse(ctrl *gomock.Controller) *MockWarehouse {
	mock := &MockWarehouse{ctrl: ctrl}
	mock.recorder = &MockWarehouseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWarehouse) EXPECT() *MockWarehouseMockRecorder {
	return m.recorder
}

// WarehouseCreate mocks base method.
func (m *MockWarehouse) WarehouseCreate(ctx context.Context, warehouse entity.Warehouse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WarehouseCreate", ctx, warehouse)
	ret0, _ := ret[0].(error)
	return ret0
}

// WarehouseCreate indicates an expected call of WarehouseCreate.
func (mr *MockWarehouseMockRecorder) WarehouseCreate(ctx, warehouse interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WarehouseCreate", reflect.TypeOf((*MockWarehouse)(nil).WarehouseCreate), ctx, warehouse)
}

// MockItemsRepo is a mock of ItemsRepo interface.
type MockItemsRepo struct {
	ctrl     *gomock.Controller
	recorder *MockItemsRepoMockRecorder
}

// MockItemsRepoMockRecorder is the mock recorder for MockItemsRepo.
type MockItemsRepoMockRecorder struct {
	mock *MockItemsRepo
}

// NewMockItemsRepo creates a new mock instance.
func NewMockItemsRepo(ctrl *gomock.Controller) *MockItemsRepo {
	mock := &MockItemsRepo{ctrl: ctrl}
	mock.recorder = &MockItemsRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemsRepo) EXPECT() *MockItemsRepoMockRecorder {
	return m.recorder
}

// QuantityByWarehouse mocks base method.
func (m *MockItemsRepo) QuantityByWarehouse(arg0 context.Context, arg1 int) (map[string]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuantityByWarehouse", arg0, arg1)
	ret0, _ := ret[0].(map[string]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuantityByWarehouse indicates an expected call of QuantityByWarehouse.
func (mr *MockItemsRepoMockRecorder) QuantityByWarehouse(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuantityByWarehouse", reflect.TypeOf((*MockItemsRepo)(nil).QuantityByWarehouse), arg0, arg1)
}

// StoreItems mocks base method.
func (m *MockItemsRepo) StoreItems(arg0 context.Context, arg1 []entity.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreItems", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreItems indicates an expected call of StoreItems.
func (mr *MockItemsRepoMockRecorder) StoreItems(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreItems", reflect.TypeOf((*MockItemsRepo)(nil).StoreItems), arg0, arg1)
}

// MockReservationsRepo is a mock of ReservationsRepo interface.
type MockReservationsRepo struct {
	ctrl     *gomock.Controller
	recorder *MockReservationsRepoMockRecorder
}

// MockReservationsRepoMockRecorder is the mock recorder for MockReservationsRepo.
type MockReservationsRepoMockRecorder struct {
	mock *MockReservationsRepo
}

// NewMockReservationsRepo creates a new mock instance.
func NewMockReservationsRepo(ctrl *gomock.Controller) *MockReservationsRepo {
	mock := &MockReservationsRepo{ctrl: ctrl}
	mock.recorder = &MockReservationsRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReservationsRepo) EXPECT() *MockReservationsRepoMockRecorder {
	return m.recorder
}

// CreateReservation mocks base method.
func (m *MockReservationsRepo) CreateReservation(arg0 context.Context, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReservation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateReservation indicates an expected call of CreateReservation.
func (mr *MockReservationsRepoMockRecorder) CreateReservation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReservation", reflect.TypeOf((*MockReservationsRepo)(nil).CreateReservation), arg0, arg1)
}

// DeleteReservation mocks base method.
func (m *MockReservationsRepo) DeleteReservation(arg0 context.Context, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReservation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteReservation indicates an expected call of DeleteReservation.
func (mr *MockReservationsRepoMockRecorder) DeleteReservation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReservation", reflect.TypeOf((*MockReservationsRepo)(nil).DeleteReservation), arg0, arg1)
}

// MockWarehousesRepo is a mock of WarehousesRepo interface.
type MockWarehousesRepo struct {
	ctrl     *gomock.Controller
	recorder *MockWarehousesRepoMockRecorder
}

// MockWarehousesRepoMockRecorder is the mock recorder for MockWarehousesRepo.
type MockWarehousesRepoMockRecorder struct {
	mock *MockWarehousesRepo
}

// NewMockWarehousesRepo creates a new mock instance.
func NewMockWarehousesRepo(ctrl *gomock.Controller) *MockWarehousesRepo {
	mock := &MockWarehousesRepo{ctrl: ctrl}
	mock.recorder = &MockWarehousesRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWarehousesRepo) EXPECT() *MockWarehousesRepoMockRecorder {
	return m.recorder
}

// CreateWarehouse mocks base method.
func (m *MockWarehousesRepo) CreateWarehouse(arg0 context.Context, arg1 entity.Warehouse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWarehouse", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWarehouse indicates an expected call of CreateWarehouse.
func (mr *MockWarehousesRepoMockRecorder) CreateWarehouse(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWarehouse", reflect.TypeOf((*MockWarehousesRepo)(nil).CreateWarehouse), arg0, arg1)
}
