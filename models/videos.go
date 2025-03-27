package models

import (
	"time"
)

type Video struct {
	ID          int       `json:"id" db:"id"`
	SectionID   int       `json:"section_id" db:"section_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	URL         string    `json:"url" db:"url"`
	Duration    int       `json:"duration" db:"duration"`
	Order       int       `json:"order" db:"order"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}



type CreateVideo struct {
	SectionID   int       `json:"section_id" db:"section_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	URL         string    `json:"url" db:"url"`
	Duration    int       `json:"duration" db:"duration"`
	Order       int       `json:"order" db:"order"`
}

type UpdateVideo struct {
	ID        int       `json:"id" db:"id"`
	SectionID   *int       `json:"section_id" db:"section_id"`
	Title       *string    `json:"title" db:"title"`
	Description *string    `json:"description" db:"description"`
	URL         *string    `json:"url" db:"url"`
	Duration    *int       `json:"duration" db:"duration"`
	Order       *int       `json:"order" db:"order"`
}

type GetListVideoRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListVideoResponse struct {
	Count int     `json:"count"`
	Videos []*Video `json:"videos"`
}
