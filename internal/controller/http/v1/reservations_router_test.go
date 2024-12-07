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

	"github.com/robertgarayshin/wms/internal/usecase"
	"github.com/robertgarayshin/wms/internal/usecase/mock"
	"github.com/robertgarayshin/wms/pkg/customerrors"
	"github.com/robertgarayshin/wms/pkg/logger"
)

type reservationsSuite struct {
	handler      *gin.RouterGroup
	reservations usecase.Reservations
	logger       logger.Interface
}

func defaultReservationsSuite(t *testing.T) *reservationsSuite {
	ctrl := gomock.NewController(t)

	return &reservationsSuite{
		handler:      gin.New().Group("/v1"),
		reservations: mock.NewMockReservations(ctrl),
		logger:       &mockLogger{},
	}
}

func TestNewReservationsAPIRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suit := defaultReservationsSuite(t)
			if tt.prepareFn != nil {
				tt.prepareFn(suit)
			}
			newReservationsAPIRoutes(suit.handler, suit.reservations, suit.logger)
		})
	}
}

func TestReserve(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name         string
		requestBody  string
		prepareFn    func(suite *reservationsSuite)
		expectedCode int
		expectedBody map[string]interface{}
	}{
		{
			name:        "success",
			requestBody: `{"ids":["id1","id2"]}`,
			prepareFn: func(suite *reservationsSuite) {
				suite.reservations.(*mock.MockReservations).
					EXPECT().
					Reserve(gomock.Any(), []string{"id1", "id2"}).
					Return(nil)
			},
			expectedCode: http.StatusCreated,
			expectedBody: map[string]interface{}{
				"status":         201,
				"status_message": "Created",
				"message":        "reservation successfully created",
				"error":          nil,
			},
		},
		{
			name:         "invalid JSON",
			requestBody:  `{"ids":`,
			prepareFn:    nil,
			expectedCode: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"status":         400,
				"status_message": "Bad Request",
				"message":        nil,
				"error":          "invalid request",
			},
		},
		{
			name:        "warehouse unavailable",
			requestBody: `{"ids":["id1","id2"]}`,
			prepareFn: func(suite *reservationsSuite) {
				suite.reservations.(*mock.MockReservations).
					EXPECT().
					Reserve(gomock.Any(), []string{"id1", "id2"}).
					Return(customerrors.ErrWarehouseUnavailable)
			},
			expectedCode: http.StatusForbidden,
			expectedBody: map[string]interface{}{
				"status":         403,
				"status_message": "Forbidden",
				"message":        nil,
				"error":          "warehouse is unavailable",
			},
		},
		{
			name:        "internal error",
			requestBody: `{"ids":["id1","id2"]}`,
			prepareFn: func(suite *reservationsSuite) {
				suite.reservations.(*mock.MockReservations).
					EXPECT().
					Reserve(gomock.Any(), []string{"id1", "id2"}).
					Return(errors.New("reservation failed"))
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"status":         500,
				"status_message": "Internal Server Error",
				"message":        nil,
				"error":          "reservation error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suite := defaultReservationsSuite(t)
			if tt.prepareFn != nil {
				tt.prepareFn(suite)
			}

			reservations := &reservationsAPIRoutes{
				reservations: suite.reservations,
				l:            suite.logger,
			}

			req := httptest.NewRequest(http.MethodPost, "/reserve", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(w)
			c.Request = req

			reservations.reserve(c)

			assert.Equal(t, tt.expectedCode, w.Code)

			resp, _ := json.Marshal(tt.expectedBody)
			assert.JSONEq(t, string(resp), w.Body.String())
		})
	}
}

func TestDeleteReservation(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name         string
		requestBody  string
		prepareFn    func(suite *reservationsSuite)
		expectedCode int
		expectedBody map[string]interface{}
	}{
		{
			name:        "success",
			requestBody: `{"ids":["id1","id2"]}`,
			prepareFn: func(suite *reservationsSuite) {
				suite.reservations.(*mock.MockReservations).
					EXPECT().
					CancelReservation(gomock.Any(), []string{"id1", "id2"}).
					Return(nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: map[string]interface{}{
				"status":         200,
				"status_message": "OK",
				"message":        "reservation successfully cancelled",
				"error":          nil,
			},
		},
		{
			name:         "invalid JSON",
			requestBody:  `{"ids":`,
			prepareFn:    nil,
			expectedCode: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"status":         400,
				"status_message": "Bad Request",
				"message":        nil,
				"error":          "invalid request body",
			},
		},
		{
			name:        "no reservations",
			requestBody: `{"ids":["id1","id2"]}`,
			prepareFn: func(suite *reservationsSuite) {
				suite.reservations.(*mock.MockReservations).
					EXPECT().
					CancelReservation(gomock.Any(), []string{"id1", "id2"}).
					Return(customerrors.ErrNoReservation)
			},
			expectedCode: http.StatusForbidden,
			expectedBody: map[string]interface{}{
				"status":         403,
				"status_message": "Forbidden",
				"message":        nil,
				"error":          "item have no reservations",
			},
		},
		{
			name:        "internal error",
			requestBody: `{"ids":["id1","id2"]}`,
			prepareFn: func(suite *reservationsSuite) {
				suite.reservations.(*mock.MockReservations).
					EXPECT().
					CancelReservation(gomock.Any(), []string{"id1", "id2"}).
					Return(errors.New("cancel reservation error"))
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"status":         500,
				"status_message": "Internal Server Error",
				"message":        nil,
				"error":          "error canceling reservation",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suite := defaultReservationsSuite(t)
			if tt.prepareFn != nil {
				tt.prepareFn(suite)
			}

			reservations := &reservationsAPIRoutes{
				reservations: suite.reservations,
				l:            suite.logger,
			}

			req := httptest.NewRequest(http.MethodDelete, "/reserve", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(w)
			c.Request = req

			reservations.deleteReservation(c)

			assert.Equal(t, tt.expectedCode, w.Code)

			resp, _ := json.Marshal(tt.expectedBody)
			assert.JSONEq(t, string(resp), w.Body.String())
		})
	}
}
