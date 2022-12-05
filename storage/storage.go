package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/samandar2605/booking/storage/postgres"
	"github.com/samandar2605/booking/storage/repo"
)

type StorageI interface {
	User() repo.UserStorageI
	Hotel() repo.HotelStorageI
	Comment() repo.CommentStorageI
	Like() repo.LikeStorageI
	Order() repo.OrderStorageI
}

type storagePg struct {
	userRepo    repo.UserStorageI
	hotelRepo   repo.HotelStorageI
	commentRepo repo.CommentStorageI
	likeRepo    repo.LikeStorageI
	orderRepo   repo.OrderStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		userRepo:    postgres.NewUser(db),
		commentRepo: postgres.NewComment(db),
		hotelRepo:   postgres.NewHotel(db),
		likeRepo:    postgres.NewLike(db),
		orderRepo:   postgres.NewOrder(db),
	}
}

func (s *storagePg) User() repo.UserStorageI {
	return s.userRepo
}

func (s *storagePg) Comment() repo.CommentStorageI {
	return s.commentRepo
}

func (s *storagePg) Hotel() repo.HotelStorageI {
	return s.hotelRepo
}

func (s *storagePg) Like() repo.LikeStorageI {
	return s.likeRepo
}

func (s *storagePg) Order() repo.OrderStorageI {
	return s.orderRepo
}
