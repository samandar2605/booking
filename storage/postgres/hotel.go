package postgres

import (
	"fmt"

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
		INSERT INTO hotels(
			name,
			phone_number,
			email,
			image_url,
			address
		)values($1,$2,$3,$4,$5)
		RETURNING id,rooms_count,price
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
		&h.Price,
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
		)values($1,$2,$3,$4,$5,$6)
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

	_, err := hr.db.Exec("update hotels set rooms_count=rooms_count+1,price=(select min(price_for_adults) from rooms where hotel_id=$1) where id=$2", room.HotelId, room.HotelId)
	if err != nil {
		return nil, err
	}
	if err := row.Scan(&room.Id); err != nil {
		return nil, err
	}

	return room, nil
}

func (hr *hotelRepo) AddRoomsImage(roomImage *repo.RoomsImage) (*repo.RoomsImage, error) {
	query := `
		insert into room_images(
			room_id,
			hotel_id,
			image_url,
			sequence_number
		)values($1,$2,$3,$4)
		returning id
	`

	row := hr.db.QueryRow(
		query,
		roomImage.RoomId,
		roomImage.HotelId,
		roomImage.ImageUrl,
		roomImage.SequenceNumber,
	)

	if err := row.Scan(&roomImage.Id); err != nil {
		return nil, err
	}

	return roomImage, nil
}

func (hr *hotelRepo) AddHotelImage(roomImage *repo.HotelImage) (*repo.HotelImage, error) {
	query := `
		insert into hotel_images(
			hotel_id,
			image_url,
			sequence_number
		)values($1,$2,$3)
		returning id
	`

	row := hr.db.QueryRow(
		query,
		roomImage.HotelId,
		roomImage.ImageUrl,
		roomImage.SequenceNumber,
	)

	if err := row.Scan(&roomImage.Id); err != nil {
		return nil, err
	}

	return roomImage, nil
}

func (hr *hotelRepo) GetHotel(id int) (*repo.Hotel, error) {
	var hotel repo.Hotel

	query := `
		SELECT
			id,
			name,
			image_url,
			phone_number,
			email,
			address,
			rating,
			price,
			rooms_count
		from hotels where id=$1
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
		&hotel.Price,
		&hotel.RoomsCount,
	); err != nil {
		return nil, err
	}

	return &hotel, nil
}

func (hr *hotelRepo) GetAll(params repo.GetHotelsQuery) (*repo.GetAllsHotelsResult, error) {
	result := repo.GetAllsHotelsResult{
		Hotels: make([]*repo.Hotel, 0),
	}

	offset := (params.Page - 1) * params.Limit
	limit := fmt.Sprintf("LIMIT %d OFFSET %d", params.Limit, offset)
	filter := "where true"

	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter += fmt.Sprintf(`
			and h.name ILIKE '%s' OR h.email ILIKE '%s' OR h.address ILIKE '%s' `, str, str, str)
	}

	if params.SortByPrice == "" {
		params.SortByPrice = "desc"
	}

	query := `
		Select 
			id,
			name,
			image_url,
			phone_number,
			email,
			price,
			address,
			rating,
			rooms_count
		from hotels ` + filter + ` order by price ` + params.SortByPrice + ` ` + limit

	rows, err := hr.db.Query(query)
	if err != nil {
		return nil, err
	}

	fmt.Println(query)

	defer rows.Close()
	for rows.Next() {
		var hotel repo.Hotel
		if err := rows.Scan(
			&hotel.Id,
			&hotel.Name,
			&hotel.ImageUrl,
			&hotel.PhoneNumber,
			&hotel.Email,
			&hotel.Price,
			&hotel.Address,
			&hotel.Rating,
			&hotel.RoomsCount,
		); err != nil {
			return nil, err
		}
		result.Hotels = append(result.Hotels, &hotel)
	}
	queryCount := `SELECT count(1) FROM hotels ` + filter
	err = hr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}
	fmt.Println(result)
	return &result, nil

}

func (hr *hotelRepo) AddHotelsImage(hotelImage *repo.HotelImage) (*repo.HotelImage, error) {
	query := `
		insert into hotel_images(
			hotel_id,
			image_url,
			sequence_number
		)values($1,$2,$3)
		returning id
	`

	row := hr.db.QueryRow(
		query,
		hotelImage.HotelId,
		hotelImage.ImageUrl,
		hotelImage.SequenceNumber,
	)

	if err := row.Scan(&hotelImage.Id); err != nil {
		return nil, err
	}

	return hotelImage, nil
}
