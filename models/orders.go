package models

import (
	"time"
)


type Order struct {
	ID        int       `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	CourseID  int       `json:"course_id" db:"course_id"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CreateOrder struct {
	UserID   string `json:"user_id" db:"user_id"`
	CourseID int    `json:"course_id" db:"course_id"`
	Status   string `json:"status" db:"status"`
}

type UpdateOrder struct {
	ID       int    `json:"id" db:"id"`
	UserID   string `json:"user_id" db:"user_id"`
	CourseID int    `json:"course_id" db:"course_id"`
	Status   string `json:"status" db:"status"`
}

type GetListOrderRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListOrderResponse struct {
	Count  int      `json:"count"`
	Orders []*Order `json:"Orders"`
}
