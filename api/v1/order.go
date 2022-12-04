package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samandar2605/booking/api/models"
	"github.com/samandar2605/booking/storage/repo"
)

// @Security ApiKeyAuth
// @Router /orders [post]
// @Summary Create a order
// @Description Create a order
// @Tags order
// @Accept json
// @Produce json
// @Param order body models.CreateOrder true "order"
// @Success 201 {object} models.Order
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateOrder(c *gin.Context) {
	var (
		req models.CreateOrder
	)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	payload, err := h.GetAuthPayload(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Order().CreateOrder(&repo.Order{
		RoomId:        req.RoomId,
		FullName:      payload.FirstName + ` ` + payload.LastName,
		DateFirst:     req.DateFirst,
		DateLast:      req.DateLast,
		AdultsCount:   req.AdultsCount,
		ChildrenCount: req.ChildrenCount,
		UserId:        payload.UserId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.Order{
		Id:            resp.Id,
		RoomId:        resp.RoomId,
		FullName:      resp.FullName,
		DateFirst:     resp.DateFirst,
		DateLast:      resp.DateLast,
		AdultsCount:   resp.AdultsCount,
		ChildrenCount: resp.ChildrenCount,
		UserId:        resp.UserId,
	})
}

// @Router /orders/{id} [get]
// @Summary Get user by id
// @Description Get user by id
// @Tags order
// @Accept json
// @Produce json
// @Success 200 {object} models.Order
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetOrder(c *gin.Context) {
	payload, err := h.GetAuthPayload(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Order().GetOrder(payload.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	for _, i := range resp {
		c.JSON(http.StatusOK, parseOrderModel(i))
	}
}

func parseOrderModel(order *repo.Order) models.Order {
	return models.Order{
		Id:            order.Id,
		RoomId:        order.RoomId,
		FullName:      order.FullName,
		DateFirst:     order.DateFirst,
		DateLast:      order.DateLast,
		AdultsCount:   order.AdultsCount,
		ChildrenCount: order.ChildrenCount,
		UserId:        order.UserId,
	}
}
