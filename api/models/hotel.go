package models

type Hotel struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	ImageUrl    string  `json:"image_url"`
	Address     string  `json:"address"`
	Rating      float64 `json:"rating"`
	Email       string  `json:"email"`
	Price       float64 `json:"price"`
	PhoneNumber string  `json:"phone_number"`
	RoomsCount  int     `json:"rooms_count"`
	Rooms       []*Room `json:"room"`
}

type HotelAll struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	ImageUrl    string  `json:"image_url"`
	Address     string  `json:"address"`
	Rating      float64 `json:"rating"`
	Email       string  `json:"email"`
	Price       float64 `json:"price"`
	PhoneNumber string  `json:"phone_number"`
	RoomsCount  int     `json:"rooms_count"`
}

type CreateHotel struct {
	Name        string `json:"name"`
	ImageUrl    string `json:"image_url"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
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
type RoomInfo struct {
	Id               int             `json:"id"`
	HotelId          int             `json:"hotel_id"`
	ImageUrl         string          `json:"image_url"`
	IsActive         bool            `json:"is_active"`
	RoomType         string          `json:"room_type"`
	PriceForChildren float64         `json:"price_for_children"`
	PriceForAdults   float64         `json:"price_for_adults"`
	Images           []*AddRoomImage `json:"images"`
}

type AddRoom struct {
	HotelId          int     `json:"hotel_id"`
	ImageUrl         string  `json:"image_url"`
	IsActive         bool    `json:"is_active"`
	RoomType         string  `json:"room_type"`
	PriceForChildren float64 `json:"price_for_children"`
	PriceForAdults   float64 `json:"price_for_adults"`
}

type AddRoomImage struct {
	Id             int    `json:"id"`
	HotelId        int    `json:"hotel_id"`
	RoomId         int    `json:"room_id"`
	ImageUrl       string `json:"room_url"`
	SequenceNumber int    `json:"sequence_number"`
}

type AddHoltelImage struct {
	Id             int    `json:"id"`
	HotelId        int    `json:"hotel_id"`
	ImageUrl       string `json:"room_url"`
	SequenceNumber int    `json:"sequence_number"`
}
type CreateAddHotelImage struct {
	HotelId        int    `json:"hotel_id"`
	ImageUrl       string `json:"room_url"`
	SequenceNumber int    `json:"sequence_number"`
}

type CreateAddRoomImage struct {
	HotelId        int    `json:"hotel_id"`
	RoomId         int    `json:"room_id"`
	ImageUrl       string `json:"room_url"`
	SequenceNumber int    `json:"sequence_number"`
}

type GetAllHotelsParams struct {
	Limit       int    `json:"limit" binding:"required" default:"10"`
	Page        int    `json:"page" binding:"required" default:"1"`
	Search      string `json:"search"`
	SortByPrice string `json:"sort_by_price" enums:"asc,desc"`
}

type GetAllHotelsResponse struct {
	Hotels []*HotelAll `json:"hotels"`
	Count  int         `json:"count"`
}

type GetRoomReq struct {
	Id      int `json:"id"`
	HotelId int `json:"hotel_id"`
}
