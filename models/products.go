package models

import "time"

type Product struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Price         float64   `json:"price"`          // NUMERIC(10, 2) uchun float64 ishlatiladi
	StockQuantity int       `json:"stock_quantity"` // Ombordagi mahsulot miqdori
	CreatedAt     time.Time `json:"created_at"`     // Yaratilgan vaqt
	UpdatedAt     time.Time `json:"updated_at"`     // Yangilangan vaqt
}

type CreateProduct struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`          // NUMERIC(10, 2) uchun float64 ishlatiladi
	StockQuantity int     `json:"stock_quantity"` // Ombordagi mahsulot miqdori
}

type UpdateProduct struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`          // NUMERIC(10, 2) uchun float64 ishlatiladi
	StockQuantity int     `json:"stock_quantity"` // Ombordagi mahsulot miqdori
}

type GetListProductRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListProductResponse struct {
	Count    int        `json:"count"`
	Products []*Product `json:"products"`
}
