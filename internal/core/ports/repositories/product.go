package repositories

import (
	"github.com/backstagefood/backstagefood/internal/core/domain"
)

type Product interface {
	ListProducts(description string) ([]*domain.Product, error)
	FindProductById(id string) (*domain.Product, error)
	CreateProduct(product *domain.Product) (*domain.Product, error)
	GetCategoryID(categoryName string) (string, error)
}
