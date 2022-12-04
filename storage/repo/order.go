package repo

import "time"

type Order struct {
	Id            int
	HotelId       int
	RoomId        int
	FullName      string
	DateFirst     time.Time
	DateLast      time.Time
	AdultsCount   int
	ChildrenCount int
	UserId        int
}

type OrderStorageI interface {
	CreateOrder(order *Order) (*Order, error)
	GetOrder(userId int) ([]*Order, error)
}
