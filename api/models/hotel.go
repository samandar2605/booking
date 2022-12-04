package models

type Hotel struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	ImageUrl    string  `json:"image_url"`
	Address     string  `json:"address"`
	Rating      float64 `json:"rating"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phone_number"`
	RoomsCount  int     `json:"rooms_count"`
}

type CreateHotel struct {
	Name        string  `json:"name"`
	ImageUrl    string  `json:"image_url"`
	Address     string  `json:"address"`
	Rating      float64 `json:"rating"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phone_number"`
	RoomsCount  int     `json:"rooms_count"`
}

type Room struct {
	Id               int     `json:"id"`
	HotelId          int     `json:"hotel_id"`
	ImageUrl         string  `json:"image_url"`
	IsActive         bool    `json:"is_active"`
	RoomType         string  `json:"room_type"`
	PriceForChildren float64 `json:"price_for_children"`
	PriceForAdults   float64 `json:"price_for_adults"`
}

type AddRoom struct {
	HotelId          int     `json:"hotel_id"`
	ImageUrl         string  `json:"image_url"`
	IsActive         bool    `json:"is_active"`
	RoomType         string  `json:"room_type"`
	PriceForChildren float64 `json:"price_for_children"`
	PriceForAdults   float64 `json:"price_for_adults"`
}

type GetAllHotelsParams struct {
	Limit       int    `json:"limit" binding:"required" default:"10"`
	Page        int    `json:"page" binding:"required" default:"1"`
	Search      string `json:"search"`
	SortByPrice string `json:"sort_by_price" enums:"asc,desc"`
}

type GetAllHotelsResponse struct {
	Hotels []*Hotel `json:"hotels"`
	Count  int      `json:"count"`
}
