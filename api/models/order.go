package models

import "time"

type Order struct {
	Id            int       `json:"id"`
	HotelId       int       `json:"hotel_id"`
	RoomId        int       `json:"room_id"`
	DateFirst     time.Time `json:"date_first"`
	DateLast      time.Time `json:"date_last"`
	Days          int       `json:"days"`
	AdultsCount   int       `json:"adults_count"`
	ChildrenCount int       `json:"children_count"`
	UserId        int       `json:"user_id"`
	Summa         float64   `json:"summa"`
}

type CreateOrder struct {
	HotelId       int       `json:"hotel_id"`
	RoomId        int       `json:"room_id"`
	DateFirst     time.Time `json:"date_first"`
	Days          int       `json:"days"`
	AdultsCount   int       `json:"adults_count"`
	ChildrenCount int       `json:"children_count"`
}
