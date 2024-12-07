package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/robertgarayshin/wms/pkg/customerrors"

	"github.com/gin-gonic/gin"

	"github.com/robertgarayshin/wms/internal/entity"
	"github.com/robertgarayshin/wms/internal/usecase"
	"github.com/robertgarayshin/wms/pkg/logger"
)

type itemsAPIRouter struct {
	items usecase.Items
	l     logger.Interface
}

func newItemsAPIRoutes(handler *gin.RouterGroup, i usecase.Items, l logger.Interface) {
	items := &itemsAPIRouter{
		items: i,
		l:     l,
	}

	h := handler.Group("/items")
	{
		h.GET("/:warehouse_id/quantity", items.getItemsQuantity)
		h.PUT("", items.createItems)
	}
}

// @Summary     Get items quantity
// @Description Count items in warehouse
// @ID          getItemsQuantity
// @Tags  	    itmes
// @Accept      json
// @Produce     json
// @Param 		warehouse_id	path 		int			true 	"warehouse_id"
// @Success     200 			{object} 	response
// @Failure		400				{object}	response
// @Failure     500 			{object} 	response
// @Router      /items/{warehouse_id}/quantity 			[get]
func (r *itemsAPIRouter) getItemsQuantity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("warehouse_id"))
	if err != nil {
		r.l.Error(err, "error converting warehouse_id to int")
		errorResponse(c, http.StatusBadRequest, "error converting warehouse_id to int")

		return
	}

	q, err := r.items.Quantity(c.Request.Context(), id)
	if err != nil {
		r.l.Error(err, "error getting items quantity")
		errorResponse(c, http.StatusInternalServerError, "error getting items quantity")

		return
	}

	successResponse(c, http.StatusOK, q)
}

type itemsCreateRequest struct {
	Items []entity.Item `json:"items"`
}

// @Summary     Create items
// @Description Create items in warehouse
// @ID          createItem
// @Tags  	    itmes
// @Accept      json
// @Produce     json
// @Param 		item	 		body 		itemsCreateRequest		true 	"items"
// @Success     201 			{object} 	response
// @Failure		400				{object}	response
// @Failure		404				{object}	response
// @Failure     500 			{object} 	response
// @Router      /items 			[put]
func (r *itemsAPIRouter) createItems(c *gin.Context) {
	var itemsReq itemsCreateRequest
	if err := c.BindJSON(&itemsReq); err != nil {
		r.l.Error(err, "error binding JSON")
		errorResponse(c, http.StatusBadRequest, "provided data is invalid")

		return
	}
	err := r.items.CreateItems(c.Request.Context(), itemsReq.Items)
	if errors.Is(err, customerrors.ErrNoWarehouse) {
		r.l.Error(err, "warehouse is not exist")
		errorResponse(c, http.StatusNotFound, "warehouse is not exist")

		return
	} else if err != nil {
		r.l.Error(err, "failed to create item")
		errorResponse(c, http.StatusInternalServerError, "items service problems")

		return
	}

	successResponse(c, http.StatusCreated, "items successfully created")
}
