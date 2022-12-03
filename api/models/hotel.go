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
	Room        Room    `json:"room"`
}

type CreateHotel struct {
	Name        string  `json:"name"`
	ImageUrl    string  `json:"image_url"`
	Address     string  `json:"address"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phone_number"`
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
