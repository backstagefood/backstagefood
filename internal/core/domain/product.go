package domain

import (
	"time"

	"github.com/pkg/errors"
)

var (
	// ErrInvalidProductCategory is returned when the product category is invalid
	ErrInvalidProductCategory = errors.New("invalid product category")
	// ErrInvalidProductDescription is returned when the product description is invalid
	ErrInvalidProductDescription = errors.New("invalid product description")
	// ErrInvalidProductIngredients is returned when the product ingredients are invalid
	ErrInvalidProductIngredients = errors.New("invalid product ingredients")
	// ErrInvalidProductPrice is returned when the product price is invalid
	ErrInvalidProductPrice = errors.New("invalid product price")
)

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

// ValidateProduct validates a product
func ValidateProduct(product *Product) error {
	// if product.IDCategory == "" {
	// 	return ErrInvalidProductCategory
	// }
	if product.Description == "" {
		return ErrInvalidProductDescription
	}
	if product.Ingredients == "" {
		return ErrInvalidProductIngredients
	}
	if product.Price <= 0 {
		return ErrInvalidProductPrice
	}
	return nil
}

func GetNow() time.Time {
	return time.Now()
}
