package services

import "github.com/backstagefood/backstagefood/internal/core/domain"

// todo: must be domain
type ProductDTO struct {
	Id          string  `json:"id"`
	Description string  `json:"description"`
	Ingredients string  `json:"ingredients"`
	Price       float64 `json:"price"`
	IDCategory  string  `json:"category_id"`
	Category    string  `json:"category"`
}

type Product interface {
	// todo: must be domain
	GetProductById(id string) (*domain.Product, error)
	// todo: must be domain
	GetProducts(description string) ([]*domain.Product, error)
	// todo: must be domain
	CreateProduct(productDTO *ProductDTO) (*domain.Product, error)
	UpdateProduct(productDTO *ProductDTO) (*domain.Product, error)
	DeleteProduct(productID string) error
	GetCategoryID(categoryName string) (string, error)
	GetCategories() ([]*domain.ProductCategory, error)
}
