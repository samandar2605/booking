package repo

import (
	"time"
)

type GetCommentQuery struct {
	Page       int    `json:"page" db:"page" binding:"required" default:"1"`
	Limit      int    `json:"limit" db:"limit" binding:"required" default:"10"`
	HotelId     int    `json:"hotel_id" db:"hotel_id"`
	UserId     int    `json:"user_id" db:"user_id"`
	SortByDate string `json:"sort_by_date" enums:"asc,desc" default:"desc"`
}

type GetAllCommentsResult struct {
	Comments []*Comment
	Count    int
}

type Comment struct {
	Id          int         `json:"id" db:"id"`
	HotelId      int         `json:"hotel_id" db:"hotel_id"`
	UserId      int         `json:"user_id" db:"user_id"`
	Description string      `json:"description" db:"description"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
	User        UserProfile `json:"user"`
}

type CommentStorageI interface {
	Create(comment *Comment) (*Comment, error)
	Get(id int) (*Comment, error)
	GetAll(param GetCommentQuery) (*GetAllCommentsResult, error)
	Update(cr *Comment) (*Comment, error)
	GetUserInfo(id int) int
	Delete(id int) error
}
