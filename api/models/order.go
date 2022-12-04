package models

import "time"

type Order struct {
	Id            int       `json:"id"`
	HotelId       int       `json:"hotel_id"`
	RoomId        int       `json:"room_id"`
	FullName      string    `json:"full_name"`
	DateFirst     time.Time `json:"date_first"`
	DateLast      time.Time `json:"date_last"`
	AdultsCount   int       `json:"adults_count"`
	ChildrenCount int       `json:"children_count"`
	UserId        int       `json:"user_id"`
}

type CreateOrder struct {
	HotelId       int       `json:"hotel_id"`
	RoomId        int       `json:"room_id"`
	FullName      string    `json:"full_name"`
	DateFirst     time.Time `json:"date_first"`
	DateLast      time.Time `json:"date_last"`
	AdultsCount   int       `json:"adults_count"`
	ChildrenCount int       `json:"children_count"`
	UserId        int       `json:"user_id"`
}
