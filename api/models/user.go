package models

import (
	"time"
)

type User struct {
	Id              int       `json:"id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	ProfileImageUrl *string   `json:"profile_image_url"`
	Username        string    `json:"username"`
	Password        string    `json:"password"`
	Email           string    `json:"email"`
	PhoneNumber     *string   `json:"phone_number"`
	Type            string    `json:"type"`
	CreatedAt       time.Time `json:"created_at"`
}

type CreateUser struct {
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Username        string    `json:"username"`
	Password        string    `json:"password"`
	Email           string    `json:"email"`
	Type            string    `json:"type"`
}

type GetAllUsersParams struct {
	Limit      int    `json:"limit" binding:"required" default:"10"`
	Page       int    `json:"page" binding:"required" default:"1"`
	Search     string `json:"search"`
	SortByDate string `json:"sort_by_date" enums:"asc,desc"`
}

type GetAllUsersResponse struct {
	Users []*User `json:"users"`
	Count int     `json:"count"`
}
