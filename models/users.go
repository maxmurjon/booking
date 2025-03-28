package models

import "time"

type User struct {
	Id        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	RoleId    *string   `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserPrimaryKey struct {
	Id    *string `json:"id,omitempty"`
	Email *string `json:"email,omitempty"`
}

type CreateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	RoleId	*string `json:"role_id"`
}

type UpdateUser struct {
	Id        string  `json:"id"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Password  *string `json:"password,omitempty"`
	RoleId    *string `json:"role_id,omitempty"`
	UserName  *string `json:"user_name,omitempty"`
	Email     *string `json:"email,omitempty"`
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
