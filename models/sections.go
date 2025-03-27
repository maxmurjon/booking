package models

import (
	"time"
)

type Section struct {
	ID        int       `json:"id" db:"id"`
	CourseID  int       `json:"course_id" db:"course_id"`
	Title     string    `json:"title" db:"title"`
	Order     int       `json:"order" db:"order"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CreateSection struct {
	CourseID  int       `json:"course_id" db:"course_id"`
	Title     string    `json:"title" db:"title"`
	Order     int       `json:"order" db:"order"`
}

type UpdateSection struct {
	ID        int       `json:"id" db:"id"`
	CourseID  int       `json:"course_id" db:"course_id"`
	Title     string    `json:"title" db:"title"`
	Order     int       `json:"order" db:"order"`
}

type GetListSectionRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListSectionResponse struct {
	Count int     `json:"count"`
	Sections []*Section `json:"sections"`
}
