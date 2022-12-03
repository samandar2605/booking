package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samandar2605/booking/api/models"
	"github.com/samandar2605/booking/storage/repo"
)

// @Router /hotels/{id} [get]
// @Summary Get Hotel by id
// @Description Get Hotel by id
// @Tags hotels
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Hotel
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetHotel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Hotel().Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Hotel{
		Id:          resp.Id,
		Name:        resp.Name,
		ImageUrl:    resp.ImageUrl,
		PhoneNumber: resp.PhoneNumber,
		Email:       resp.Email,
		Address:     resp.Address,
		Rating:      resp.Rating,
		RoomsCount:  resp.RoomsCount,
		Room: models.Room{
			ImageUrl:         resp.Room.ImageUrl,
			RoomType:         resp.Room.RoomType,
			IsActive:         resp.Room.IsActive,
			PriceForChildren: resp.Room.PriceForChildren,
			PriceForAdults:   resp.Room.PriceForAdults,
		},
	})
}

// @Security ApiKeyAuth
// @Router /hotels [post]
// @Summary Create a Hotel
// @Description Create a Hotel
// @Tags hotels
// @Accept json
// @Produce json
// @Param Hotel body models.CreateHotel true "Hotel"
// @Success 201 {object} models.Hotel
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateHotel(c *gin.Context) {
	var (
		req models.CreateHotel
	)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Hotel().Create(&repo.Hotel{
		Name:        req.Name,
		ImageUrl:    req.ImageUrl,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Address:     req.Address,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.Hotel{
		Id:          resp.Id,
		Name:        resp.Name,
		ImageUrl:    resp.ImageUrl,
		Address:     resp.Address,
		Rating:      resp.Rating,
		Email:       resp.Email,
		PhoneNumber: resp.PhoneNumber,
		RoomsCount:  resp.RoomsCount,
	})
}
