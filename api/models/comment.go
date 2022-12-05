package models

import (
	"time"
)

type Comment struct {
	Id          int          `json:"id" db:"id"`
	HotelId     int          `json:"hotel_id" db:"hotel_id"`
	UserId      int          `json:"user_id" db:"user_id"`
	Description string       `json:"description" db:"description"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
	User        *UserProfile `json:"user"`
}

type CommentUser struct {
	Id              int     `json:"id"`
	FirstName       string  `json:"first_name"`
	LastName        string  `json:"last_name"`
	Email           string  `json:"email"`
	ProfileImageUrl *string `json:"profile_image_url"`
}

type CreateComment struct {
	HotelId     int    `json:"hotel_id" db:"hotel_id"`
	Description string `json:"description" db:"description"`
}

type UpdateComment struct {
	Id          int       `json:"id" db:"id"`
	HotelId     int       `json:"hotel_id" db:"hotel_id"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type GetAllCommentsParams struct {
	Limit      int    `json:"limit" binding:"required" default:"10"`
	Page       int    `json:"page" binding:"required" default:"1"`
	UserID     int    `json:"user_id"`
	HotelId    int    `json:"hotel_id"`
	SortByDate string `json:"sort_by_date" binding:"required,oneof=asc desc" default:"desc"`
}

type GetAllCommentsResponse struct {
	Comments []*Comment `json:"comments"`
	Count    int        `json:"count"`
}
