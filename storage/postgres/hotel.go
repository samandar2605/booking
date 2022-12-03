package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/samandar2605/booking/storage/repo"
)

type hotelRepo struct {
	db *sqlx.DB
}

func NewHotel(db *sqlx.DB) repo.HotelStorageI {
	return &hotelRepo{db: db}
}

func (hr *hotelRepo) Create(h *repo.Hotel) (*repo.Hotel, error) {
	query := `
		INSERT INTO users(
			name,
			image_url,
			phone_number,
			email,
			address
		)values($1,$2,$3,$4,$5)
		RETURNING id,rooms_count
	`

	row := hr.db.QueryRow(
		query,
		h.Name,
		h.ImageUrl,
		h.PhoneNumber,
		h.Email,
		h.Address,
	)

	if err := row.Scan(
		&h.Id,
		&h.RoomsCount,
	); err != nil {
		return nil, err
	}

	return h, nil
}

func (hr *hotelRepo) AddRoom(room *repo.Room) (*repo.Room, error) {
	query := `
		insert into rooms(
			hotel_id,
			image_url,
			is_active,
			room_type,
			price_for_children,
			price_for_adults
		)values($1,$2,$3,$4,$5)
		returning id
	`
	row := hr.db.QueryRow(
		query,
		room.HotelId,
		room.ImageUrl,
		room.IsActive,
		room.RoomType,
		room.PriceForChildren,
		room.PriceForAdults,
	)

	if err := row.Scan(&room.Id); err != nil {
		return nil, err
	}

	return room, nil
}

func (hr *hotelRepo) AddRoomsImage(roomImage *repo.RoomsImage) (*repo.RoomsImage, error) {
	query := `
		insert into room_images(
			room_id,
			image_url,
			sequence_number
		)values($1,$2,$3)
		returning id
	`
	row := hr.db.QueryRow(
		query,
		roomImage.RoomId,
		roomImage.ImageUrl,
		roomImage.SequenceNumber,
	)

	if err := row.Scan(&roomImage.Id); err != nil {
		return nil, err
	}

	return roomImage, nil
}

func (hr *hotelRepo) Get(id int) (*repo.Hotel, error) {
	var hotel repo.Hotel

	query := `
		SELECT
			h.id,
			h.name,
			h.image_url,
			h.phone_number,
			h.email,
			h.address,
			h.rating,
			h.rooms_count,
			r.image_url,
			r.room_type,
			r.is_active,
			r.price_for_children,
			r.price_for_adults
		from hotels h 
		inner join rooms r on h.id=r.hotel_id
		inner join hotel_images h_i on h.id=hotel_id
		where h.id=$1
	`
	row := hr.db.QueryRow(
		query,
		id,
	)

	if err := row.Scan(
		&hotel.Id,
		&hotel.Name,
		&hotel.ImageUrl,
		&hotel.PhoneNumber,
		&hotel.Email,
		&hotel.Address,
		&hotel.Rating,
		&hotel.RoomsCount,
		&hotel.Room.ImageUrl,
		&hotel.Room.RoomType,
		&hotel.Room.IsActive,
		&hotel.Room.PriceForChildren,
		&hotel.Room.PriceForAdults,
	); err != nil {
		return nil, err
	}

	return &hotel, nil
}
