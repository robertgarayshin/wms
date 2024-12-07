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
	"github.com/robertgarayshin/wms/pkg/customerrors"
	"github.com/robertgarayshin/wms/pkg/logger"
)

type mockLogger struct{}

func (m mockLogger) Info(_ string, _ ...interface{})       {}
func (m mockLogger) Error(_ interface{}, _ ...interface{}) {}
func (m mockLogger) Debug(_ interface{}, _ ...interface{}) {}
func (m mockLogger) Warn(_ string, _ ...interface{})       {}
func (m mockLogger) Fatal(_ interface{}, _ ...interface{}) {}

type suite struct {
	handler *gin.RouterGroup
	items   usecase.Items
	logger  logger.Interface
}

func defaultSuite(t *testing.T) *suite {
	ctrl := gomock.NewController(t)

	return &suite{
		handler: gin.New().Group("/v1"),
		items:   mock.NewMockItems(ctrl),
		logger:  &mockLogger{},
	}
}

func TestNewItemsAPIRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
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
			newItemsAPIRoutes(suit.handler, suit.items, suit.logger)
		})
	}
}

func TestGetItemsQuantity(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var c *gin.Context

	tests := []struct {
		name      string
		prepareFn func(suite *suite)
		wantErr   bool
		status    int
		response  map[string]interface{}
	}{
		{
			name: "success",
			prepareFn: func(suite *suite) {
				c.Params = gin.Params{
					{Key: "warehouse_id", Value: "1"},
				}
				suite.items.(*mock.MockItems).
					EXPECT().
					Quantity(gomock.Any(), 1).
					Return(map[string]int{"item1": 100}, nil)
			},
			wantErr: false,
			status:  http.StatusOK,
			response: map[string]interface{}{
				"status":         200,
				"status_message": "OK",
				"message":        map[string]int{"item1": 100},
				"error":          nil,
			},
		},
		{
			name:      "invalid warehouse id",
			prepareFn: nil,
			wantErr:   true,
			status:    http.StatusBadRequest,
			response: map[string]interface{}{
				"error":          "error converting warehouse_id to int",
				"message":        nil,
				"status":         400,
				"status_message": "Bad Request",
			},
		},
		{
			name: "usecase error",
			prepareFn: func(suite *suite) {
				c.Params = gin.Params{
					{Key: "warehouse_id", Value: "1"},
				}
				suite.items.(*mock.MockItems).
					EXPECT().
					Quantity(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("error"))
			},
			wantErr: true,
			status:  http.StatusInternalServerError,
			response: map[string]interface{}{
				"error":          "error getting items quantity",
				"message":        nil,
				"status":         500,
				"status_message": "Internal Server Error",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/items/1/quantity", nil)
			w := httptest.NewRecorder()

			c, _ = gin.CreateTestContext(w)
			c.Request = req

			suit := defaultSuite(t)
			if tt.prepareFn != nil {
				tt.prepareFn(suit)
			}

			items := &itemsAPIRouter{
				items: suit.items,
				l:     suit.logger,
			}

			items.getItemsQuantity(c)

			assert.Equal(t, tt.status, w.Code)
			resp, _ := json.Marshal(tt.response)
			assert.JSONEq(t, string(resp), w.Body.String())
		})
	}
}

func TestCreateItems(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name         string
		requestBody  string
		prepareFn    func(suite *suite)
		expectedCode int
		expectedBody map[string]interface{}
	}{
		{
			name:        "success",
			requestBody: `{"items":[{"name":"item1","quantity":10},{"name":"item2","quantity":20}]}`,
			prepareFn: func(suite *suite) {
				suite.items.(*mock.MockItems).
					EXPECT().
					CreateItems(gomock.Any(), []entity.Item{
						{Name: "item1", Quantity: 10},
						{Name: "item2", Quantity: 20},
					}).
					Return(nil)
			},
			expectedCode: http.StatusCreated,
			expectedBody: map[string]interface{}{
				"error":          nil,
				"message":        "items successfully created",
				"status":         201,
				"status_message": "Created",
			},
		},
		{
			name:         "invalid JSON",
			requestBody:  `{"items": [`,
			prepareFn:    nil,
			expectedCode: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error":          "provided data is invalid",
				"message":        nil,
				"status":         400,
				"status_message": "Bad Request",
			},
		},
		{
			name:        "warehouse not exist",
			requestBody: `{"items":[{"name":"item1","quantity":10}]}`,
			prepareFn: func(suite *suite) {
				suite.items.(*mock.MockItems).
					EXPECT().
					CreateItems(gomock.Any(), []entity.Item{
						{Name: "item1", Quantity: 10},
					}).
					Return(customerrors.ErrNoWarehouse)
			},
			expectedCode: http.StatusNotFound,
			expectedBody: map[string]interface{}{
				"error":          "warehouse is not exist",
				"message":        nil,
				"status":         404,
				"status_message": "Not Found",
			},
		},
		{
			name:        "internal service error",
			requestBody: `{"items":[{"name":"item1","quantity":10}]}`,
			prepareFn: func(suite *suite) {
				suite.items.(*mock.MockItems).
					EXPECT().
					CreateItems(gomock.Any(), []entity.Item{
						{Name: "item1", Quantity: 10},
					}).
					Return(errors.New("service error"))
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"error":          "items service problems",
				"message":        nil,
				"status":         500,
				"status_message": "Internal Server Error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suite := defaultSuite(t)
			if tt.prepareFn != nil {
				tt.prepareFn(suite)
			}

			items := &itemsAPIRouter{
				items: suite.items,
				l:     suite.logger,
			}

			req := httptest.NewRequest(http.MethodPut, "/items", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(w)
			c.Request = req

			items.createItems(c)

			assert.Equal(t, tt.expectedCode, w.Code)

			resp, _ := json.Marshal(tt.expectedBody)
			assert.JSONEq(t, string(resp), w.Body.String())
		})
	}
}
