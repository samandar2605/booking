package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samandar2605/booking/api/models"
	"github.com/samandar2605/booking/storage/repo"
)

// @Security ApiKeyAuth
// @Router /likes [post]
// @Summary Create or update like
// @Description Create or update like
// @Tags like
// @Accept json
// @Produce json
// @Param like body models.CreateOrUpdateLikeRequest true "like"
// @Success 201 {object} models.Like
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateOrUpdateLike(c *gin.Context) {
	var (
		req models.CreateOrUpdateLikeRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload, err := h.GetAuthPayload(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = h.storage.Like().CreateOrUpdate(&repo.Like{
		UserId:  payload.UserId,
		HotelId: int(req.HotelId),
		Status:  req.Status,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Successfully finished",
	})
}


// @Security ApiKeyAuth
// @Router /likes/user-hotel [get]
// @Summary Get like by user and hotel
// @Description Get like by user and hotel
// @Tags like
// @Accept json
// @Produce json
// @Param hotel_id query int true "Hotel ID"
// @Success 200 {object} models.Like
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetLike(c *gin.Context) {
	hotelId, err := strconv.Atoi(c.Query("hotel_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload, err := h.GetAuthPayload(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp, err := h.storage.Like().Get((payload.UserId), hotelId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.Like{
		Id:      resp.ID,
		HotelId: resp.HotelId,
		UserId:  resp.UserId,
		Status:  resp.Status,
	})
}
