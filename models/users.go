package models

import "time"

type User struct {
	Id          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
	ImageUrl    *string   `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserPrimaryKey struct {
	Id           string `json:"id"`
	Phone_number string `json:"phone_number"`
}

type CreateUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UpdateUser struct {
	Id          string  `json:"id"`
	FirstName   *string `json:"first_name,omitempty"`
	LastName    *string `json:"last_name,omitempty"`
	Password    *string `json:"password_hash,omitempty"`
	ImageUrl    *string `json:"image_url,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
}

type GetListUserRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListUserResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"users"`
}
