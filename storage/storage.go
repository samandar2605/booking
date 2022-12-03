package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/samandar2605/booking/storage/postgres"
	"github.com/samandar2605/booking/storage/repo"
)

type StorageI interface {
	User() repo.UserStorageI
	Hotel() repo.HotelStorageI
}

type storagePg struct {
	userRepo  repo.UserStorageI
	hotelRepo repo.HotelStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		userRepo:  postgres.NewUser(db),
		hotelRepo: postgres.NewHotel(db),
	}
}

func (s *storagePg) User() repo.UserStorageI {
	return s.userRepo
}

func (s *storagePg) Hotel() repo.HotelStorageI {
	return s.hotelRepo
}
