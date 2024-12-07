package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/robertgarayshin/wms/internal/entity"
	"github.com/robertgarayshin/wms/internal/usecase"
	"github.com/robertgarayshin/wms/internal/usecase/mock"
	"github.com/robertgarayshin/wms/pkg/logger"
)

type warehousesSuite struct {
	handler   *gin.RouterGroup
	warehouse usecase.Warehouse
	logger    logger.Interface
}

func defaultWarehousesSuite(t *testing.T) *warehousesSuite {
	ctrl := gomock.NewController(t)

	return &warehousesSuite{
		handler:   gin.New().Group("/v1"),
		warehouse: mock.NewMockWarehouse(ctrl),
		logger:    &mockLogger{},
	}
}

func TestNewWarehousesAPIRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suit := defaultWarehousesSuite(t)
			if tt.prepareFn != nil {
				tt.prepareFn(suit)
			}
			newWarehousesAPIRoutes(suit.handler, suit.warehouse, suit.logger)
		})
	}
}

func TestCreateWarehouse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name         string
		requestBody  string
		prepareFn    func(suite *warehousesSuite)
		expectedCode int
		expectedBody map[string]interface{}
	}{
		{
			name:        "success",
			requestBody: `{"warehouse":{"name":"Warehouse A", "availability": true}}`,
			prepareFn: func(suite *warehousesSuite) {
				suite.warehouse.(*mock.MockWarehouse).
					EXPECT().
					WarehouseCreate(gomock.Any(), entity.Warehouse{
						Name:         "Warehouse A",
						Availability: true,
					}).
					Return(nil)
			},
			expectedCode: http.StatusCreated,
			expectedBody: map[string]interface{}{
				"status":         201,
				"status_message": "Created",
				"message":        "warehouse created successfully",
				"error":          nil,
			},
		},
		{
			name:         "invalid JSON",
			requestBody:  `{"warehouse":`,
			prepareFn:    nil,
			expectedCode: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"status":         400,
				"status_message": "Bad Request",
				"message":        nil,
				"error":          "provided data is invalid",
			},
		},
		{
			name:        "internal error",
			requestBody: `{"warehouse":{"name":"Warehouse A", "availability": true}}`,
			prepareFn: func(suite *warehousesSuite) {
				suite.warehouse.(*mock.MockWarehouse).
					EXPECT().
					WarehouseCreate(gomock.Any(), entity.Warehouse{
						Name:         "Warehouse A",
						Availability: true,
					}).
					Return(errors.New("warehouse creation failed"))
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"status":         500,
				"status_message": "Internal Server Error",
				"message":        nil,
				"error":          "warehouse service problems",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suite := defaultWarehousesSuite(t)
			if tt.prepareFn != nil {
				tt.prepareFn(suite)
			}

			warehouses := &warehousesAPIRoutes{
				warehouses: suite.warehouse,
				l:          suite.logger,
			}

			req := httptest.NewRequest(http.MethodPost, "/warehouses", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(w)
			c.Request = req

			warehouses.createWarehouse(c)

			assert.Equal(t, tt.expectedCode, w.Code)

			resp, _ := json.Marshal(tt.expectedBody)
			assert.JSONEq(t, string(resp), w.Body.String())
		})
	}
}
