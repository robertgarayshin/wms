package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/robertgarayshin/wms/internal/usecase"
	"github.com/robertgarayshin/wms/pkg/logger"

	// Документация Swagger.
	_ "github.com/robertgarayshin/wms/docs"
)

// NewRouter - метод создания нового роутера.
// Swagger spec:
// @title       WMS - Warehouses Management System
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(
	handler *gin.Engine,
	l logger.Interface,
	i usecase.Items,
	r usecase.Reservations,
	w usecase.Warehouse,
) {
	// Опции хэндлера запросов
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Prometheus метрики сервера
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Создание роутеров для API
	h := handler.Group("/v1")
	{
		newWarehousesAPIRoutes(h, w, l)
		newReservationsAPIRoutes(h, r, l)
		newItemsAPIRoutes(h, i, l)
	}
}
