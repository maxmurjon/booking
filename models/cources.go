package models

import (
	"time"
)

type Course struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	PriceTutor   float64   `json:"price_tutor"`
	PriceNoTutor float64   `json:"price_no_tutor"`
	ImageUrl     string    `json:"image_url"`
	VideoUrl     string    `json:"video_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateCourse struct {
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	PriceTutor   float64 `json:"price_tutor"`
	PriceNoTutor float64 `json:"price_no_tutor"`
	ImageUrl     string    `json:"image_url"`
	VideoUrl     string    `json:"video_url"`
}

type UpdateCourse struct {
	ID           int      `json:"id"`
	Title        *string  `json:"title,omitempty"`
	Description  *string  `json:"description,omitempty"`
	PriceTutor   *float64 `json:"price_tutor,omitempty"`
	PriceNoTutor *float64 `json:"price_no_tutor,omitempty"`
	ImageUrl     *string  `json:"image_url,omitempty"`
	VideoUrl     *string   `json:"video_url,omitempty"`
}

type GetListCourseRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListCourseResponse struct {
	Count   int       `json:"count"`
	Courses []*Course `json:"courses"`
}
