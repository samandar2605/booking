package repo

type Hotel struct {
	Id          int     `db:"id"`
	Name        string  `db:"name"`
	ImageUrl    string  `db:"image_url"`
	PhoneNumber string  `db:"phone_number"`
	Email       string  `db:"email"`
	Address     string  `db:"address"`
	Price       float64 `db:"price"`
	Rating      float64 `db:"rating"`
	RoomsCount  int     `db:"rooms_count"`
	Room        []Room  `db:"room"`
}

type Room struct {
	Id               int     `db:"id"`
	HotelId          int     `db:"hotel_id"`
	ImageUrl         string  `db:"image_url"`
	IsActive         bool    `db:"is_active"`
	RoomType         string  `db:"room_type"`
	PriceForChildren float64 `db:"price_for_children"`
	PriceForAdults   float64 `db:"price_for_adults"`
}

type RoomsImage struct {
	Id             int    `db:"id"`
	RoomId         int    `db:"room_id"`
	HotelId        int    `db:"hotel_id"`
	ImageUrl       string `db:"image_url"`
	SequenceNumber int    `db:"sequence_number"`
}

type HotelImage struct {
	Id             int    `db:"id"`
	HotelId        int    `db:"hotel_id"`
	ImageUrl       string `db:"image_url"`
	SequenceNumber int    `db:"sequence_number"`
}

type HotelStorageI interface {
	Create(h *Hotel) (*Hotel, error)
	AddRoom(room *Room) (*Room, error)
	AddRoomsImage(roomImage *RoomsImage) (*RoomsImage, error)
	AddHotelImage(roomImage *HotelImage) (*HotelImage, error)
	GetHotel(id int) (*Hotel, error)
	GetRoom(id, hotelId int) (*Room, error)
	GetAllRooms(id int) ([]*Room, error)
	GetAllRoomsImage(id, hotelId int) ([]*RoomsImage, error)
	GetAllHotels(params GetHotelsQuery) (*GetAllsHotelsResult, error)
	AddHotelsImage(hotelImage *HotelImage) (*HotelImage, error)
	DeleteHotel(id int) error
	DeleteRoom(id int) error
}

type GetHotelsQuery struct {
	Page        int    `json:"page" db:"page" binding:"required" default:"1"`
	Limit       int    `json:"limit" db:"limit" binding:"required" default:"10"`
	Search      string `json:"search"`
	SortByPrice string `json:"sort_by_price" enums:"asc,desc" default:"desc"`
}

type GetAllsHotelsResult struct {
	Hotels []*Hotel
	Count  int
}
