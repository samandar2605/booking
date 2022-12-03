package models

import "time"

type RegisterRequest struct {
	FirstName       string `json:"first_name" binding:"required,min=2,max=30"`
	LastName        string `json:"last_name" binding:"required,min=2,max=30"`
	ProfileImageUrl string `json:"profile_image_url"`
	Username        string `json:"username" binding:"required,min=2,max=30"`
	Password        string `json:"password" binding:"required,min=6,max=16"`
	Email           string `json:"email" binding:"required,email"`
	PhoneNumber     string `json:"phone_number"`
	Type            string `json:"type" binding:"required,oneof=superadmin user" default:"user"`
}

type AuthResponse struct {
	ID              int       `json:"id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	ProfileImageUrl string    `json:"profile_image_url"`
	Username        string    `json:"username"`
	Password        string    `json:"password"`
	Email           string    `json:"email"`
	PhoneNumber     string    `json:"phone_number"`
	Type            string    `json:"type"`
	CreatedAt       time.Time `json:"created_at"`
	AccessToken     string    `json:"access_token"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=16"`
}
type VerifyRequest struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required"`
}

type RedisUser struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	VerificationCode int    `json:"verification_code"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type UpdatePasswordRequest struct {
	Password string `json:"password" binding:"required"`
}
