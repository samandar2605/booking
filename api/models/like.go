package models

type Like struct {
	Id     int  `json:"id"`
	HotelId int  `json:"hotel_id"`
	UserId int  `json:"user_id"`
	Status bool `json:"status"`
}

type CreateOrUpdateLikeRequest struct {
	HotelId int `json:"hotel_id" binding:"required"`
	Status bool  `json:"status"`
}
