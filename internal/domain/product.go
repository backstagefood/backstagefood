package domain

import "time"

// Product represents a product in the system
type Product struct {
	ID              string          `json:"id"`
	IDCategory      string          `json:"id_category"`
	Description     string          `json:"description"`
	Ingredients     string          `json:"ingredients"`
	Price           float64         `json:"price"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	ProductCategory ProductCategory `json:"product_category"`
}
