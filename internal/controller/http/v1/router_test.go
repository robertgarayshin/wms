package v1

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/robertgarayshin/wms/internal/usecase/mock"
)

func TestNewRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockItems := mock.NewMockItems(ctrl)
	mockReservations := mock.NewMockReservations(ctrl)
	mockWarehouse := mock.NewMockWarehouse(ctrl)

	r := gin.New()

	NewRouter(r, mockLogger{}, mockItems, mockReservations, mockWarehouse)

	t.Run("test warehouses route", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/v1/warehouses/", nil)
		require.NoError(t, err)

		w := performRequest(r, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("test reservations route", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/v1/reserve", nil)
		require.NoError(t, err)

		w := performRequest(r, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("test items route", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/v1/items/quantity", nil)
		require.NoError(t, err)

		w := performRequest(r, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func performRequest(r http.Handler, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
