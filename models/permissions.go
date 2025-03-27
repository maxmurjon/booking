package models

import "time"

type Permission struct {
	Id          int    `json:"id"`
	Name   string    `json:"name"`
	Description    string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreatePermission struct {
	Name   string `json:"name"`
	Description    string `json:"description"`
}

type UpdatePermission struct {
	Id          int  `json:"id"`
	Name   *string `json:"name"`
	Description    *string `json:"description"`
}

type GetListPermissionRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListPermissionResponse struct {
	Count       int           `json:"count"`
	Permissions []*Permission `json:"Permissions"`
}
