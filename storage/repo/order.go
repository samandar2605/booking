package repo

import "time"

type Order struct {
	Id            int
	HotelId       int
	RoomId        int
	DateFirst     time.Time
	DateLast      time.Time
	Days          int
	AdultsCount   int
	ChildrenCount int
	UserId        int
	Summa         float64
}

type OrderStorageI interface {
	CreateOrder(order *Order) (*Order, error)
	GetOrder(userId int) ([]*Order, error)
	CheckRoom(order *Order) bool
}
