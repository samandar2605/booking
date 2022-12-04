package v1

import (
	"errors"
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

	resp, err := h.storage.Hotel().GetHotel(id)
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
		Price:       resp.Price,
		Rating:      resp.Rating,
		RoomsCount:  resp.RoomsCount,
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
// @Success 201 {object} models.CreateHotel
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
	payload, err := h.GetAuthPayload(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	if (payload.UserType != "superadmin" || payload.UserType != "admin") && payload.UserType == "user" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: errors.New("you don't create hotel").Error(),
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
		Name:        resp.Name,
		ImageUrl:    resp.ImageUrl,
		Address:     resp.Address,
		Rating:      resp.Rating,
		Price:       resp.Price,
		Email:       resp.Email,
		PhoneNumber: resp.PhoneNumber,
		RoomsCount:  resp.RoomsCount,
	})
}

// @Security ApiKeyAuth
// @Router /hotels/add-room [post]
// @Summary Add a room
// @Description Add a room
// @Tags hotels
// @Accept json
// @Produce json
// @Param AddRoom body models.AddRoom true "AddRoom"
// @Success 201 {object} models.Room
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) AddRoom(c *gin.Context) {
	var (
		req models.AddRoom
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
	if (payload.UserType != "superadmin" || payload.UserType != "admin") && payload.UserType == "user" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: errors.New("you don't create hotel").Error(),
		})
		return
	}

	resp, err := h.storage.Hotel().AddRoom(&repo.Room{
		HotelId:          req.HotelId,
		ImageUrl:         req.ImageUrl,
		IsActive:         req.IsActive,
		RoomType:         req.RoomType,
		PriceForChildren: req.PriceForChildren,
		PriceForAdults:   req.PriceForAdults,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.Room{
		Id:               resp.Id,
		HotelId:          resp.HotelId,
		ImageUrl:         resp.ImageUrl,
		IsActive:         resp.IsActive,
		RoomType:         resp.RoomType,
		PriceForChildren: resp.PriceForChildren,
		PriceForAdults:   resp.PriceForAdults,
	})
}

// @Router /hotels/add-rooms-images [post]
// @Summary Add a rooms Images
// @Description Add a rooms Images
// @Tags hotels
// @Accept json
// @Produce json
// @Param AddRoomImage body models.CreateAddRoomImage true "AddRoomImage"
// @Success 201 {object} models.AddRoomImage
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) AddRoomImage(c *gin.Context) {
	var (
		req models.CreateAddRoomImage
	)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Hotel().AddRoomsImage(&repo.RoomsImage{
		RoomId:         req.RoomId,
		HotelId:        req.HotelId,
		ImageUrl:       req.ImageUrl,
		SequenceNumber: req.SequenceNumber,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.AddRoomImage{
		Id:             resp.Id,
		RoomId:         resp.RoomId,
		HotelId:        resp.HotelId,
		ImageUrl:       resp.ImageUrl,
		SequenceNumber: resp.SequenceNumber,
	})
}

// @Router /hotels/add-hotels-images [post]
// @Summary Add a hotels Images
// @Description Add a hotels Images
// @Tags hotels
// @Accept json
// @Produce json
// @Param CreateAddHotelImage body models.CreateAddHotelImage true "CreateAddHotelImage"
// @Success 201 {object} models.AddHoltelImage
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) AddHotelImage(c *gin.Context) {
	var (
		req models.CreateAddHotelImage
	)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Hotel().AddHotelImage(&repo.HotelImage{
		HotelId:        req.HotelId,
		ImageUrl:       req.ImageUrl,
		SequenceNumber: req.SequenceNumber,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.AddHoltelImage{
		Id:             resp.Id,
		HotelId:        resp.HotelId,
		ImageUrl:       resp.ImageUrl,
		SequenceNumber: resp.SequenceNumber,
	})
}

// @Router /hotels [get]
// @Summary Get all Hotels
// @Description Get all Hotels
// @Tags hotels
// @Accept json
// @Produce json
// @Param filter query models.GetAllHotelsParams false "Filter"
// @Success 200 {object} models.GetAllHotelsResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAllHotels(c *gin.Context) {
	req, err := HotelsParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	result, err := h.storage.Hotel().GetAll(repo.GetHotelsQuery{
		Page:   req.Page,
		Limit:  req.Limit,
		Search: req.Search,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, HotelsResponse(result))
}

func HotelsParams(c *gin.Context) (*models.GetAllHotelsParams, error) {
	var (
		limit       int = 10
		page        int = 1
		sortByPrice string
		err         error
	)

	if c.Query("limit") != "" {
		limit, err = strconv.Atoi(c.Query("limit"))
		if err != nil {
			return nil, err
		}
	}

	if c.Query("page") != "" {
		page, err = strconv.Atoi(c.Query("page"))
		if err != nil {
			return nil, err
		}
	}

	if c.Query("sort_by_price") != "" &&
		(c.Query("sort_by_price") == "desc" || c.Query("sort_by_price") == "asc" || c.Query("sort_by_price") == "") {
		sortByPrice = c.Query("sort_by_price")
	}

	return &models.GetAllHotelsParams{
		Page:        page,
		Limit:       limit,
		Search:      c.Query("search"),
		SortByPrice: sortByPrice,
	}, nil
}
func HotelsResponse(data *repo.GetAllsHotelsResult) *models.GetAllHotelsResponse {
	response := models.GetAllHotelsResponse{
		Hotels: make([]*models.Hotel, 0),
		Count:  data.Count,
	}

	for _, hotel := range data.Hotels {
		p := parseHotelModel(hotel)
		response.Hotels = append(response.Hotels, &p)
	}

	return &response
}

func parseHotelModel(hotel *repo.Hotel) models.Hotel {
	return models.Hotel{
		Id:          hotel.Id,
		Name:        hotel.Name,
		ImageUrl:    hotel.ImageUrl,
		Address:     hotel.Address,
		Rating:      hotel.Rating,
		Email:       hotel.Email,
		PhoneNumber: hotel.PhoneNumber,
		RoomsCount:  hotel.RoomsCount,
	}
}
