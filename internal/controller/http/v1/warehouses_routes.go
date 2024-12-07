package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/robertgarayshin/wms/internal/entity"
	"github.com/robertgarayshin/wms/internal/usecase"
	"github.com/robertgarayshin/wms/pkg/logger"
)

type warehousesAPIRoutes struct {
	warehouses usecase.Warehouse
	l          logger.Interface
}

func newWarehousesAPIRoutes(handler *gin.RouterGroup, w usecase.Warehouse, l logger.Interface) {
	r := &warehousesAPIRoutes{w, l}

	h := handler.Group("/warehouses")
	{
		h.POST("/", r.createWarehouse)
	}
}

type createWarehouseRequest struct {
	Warehouse entity.Warehouse `json:"warehouse"`
}

// @Summary     Create warehouse
// @Description Create warehouse by provided data
// @ID          createWarehouse
// @Tags  	    warehouses
// @Accept      json
// @Produce     json
// @Param 		warehouse 		body 		createWarehouseRequest		true 	"warehouse"
// @Success     201 			{object} 	response
// @Failure		400				{object}	response
// @Failure     500 			{object} 	response
// @Router      /warehouses		[post]
func (w *warehousesAPIRoutes) createWarehouse(c *gin.Context) {
	var wh createWarehouseRequest
	if err := c.ShouldBindJSON(&wh); err != nil {
		w.l.Error(err, "error binding JSON")
		errorResponse(c, http.StatusBadRequest, "provided data is invalid")

		return
	}

	if err := w.warehouses.WarehouseCreate(c.Request.Context(), wh.Warehouse); err != nil {
		w.l.Error(err, "failed to create warehouse")
		errorResponse(c, http.StatusInternalServerError, "warehouse service problems")

		return
	}

	successResponse(c, http.StatusCreated, "warehouse created successfully")
}
