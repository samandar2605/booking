package repo

import (
	"time"
)

type User struct {
	Id              int       `db:"id"`
	FirstName       string    `db:"first_name"`
	LastName        string    `db:"last_name"`
	PhoneNumber     *string   `db:"phone_number"`
	Email           string    `db:"email"`
	UserName        string    `db:"user_name"`
	Password        string    `db:"password"`
	ProfileImageUrl *string   `db:"profile_image_url"`
	Type            string    `db:"type"`
	CreatedAt       time.Time `db:"created_at"`
}
const (
	UserTypeSuperadmin = "superadmin"
	UserTypeUser       = "user"
)

type UserProfile struct {
	Id              int     `json:"id"`
	FirstName       string  `json:"first_name"`
	LastName        string  `json:"last_name"`
	Email           string  `json:"email"`
	ProfileImageUrl *string `json:"profile_image_url"`
}


type UpdatePassword struct {
	UserID   int64
	Password string
}

type UserStorageI interface {
	Create(u *User) (*User, error)
	Get(id int) (*User, error)
	GetAll(param GetUserQuery) (*GetAllUsersResult, error)
	Update(usr *User) (*User, error)
	Delete(id int) error
	GetByEmail(email string) (*User, error)
	CheckInfo(email, username string) (*User, error)
	UpdatePassword(req *UpdatePassword) error
	GetUserProfileInfo(usrId int) (*User, error)


}

type GetUserQuery struct {
	Page       int    `json:"page" db:"page" binding:"required" default:"1"`
	Limit      int    `json:"limit" db:"limit" binding:"required" default:"10"`
	Search     string `json:"search"`
	SortByDate string `json:"sort_by_date" enums:"asc,desc" default:"desc"`
}

type GetAllUsersResult struct {
	Users []*User
	Count int
}
