package usecases

import (
	"github.com/backstagefood/backstagefood/internal/app/domain"
	"log/slog"
)

type ProductsRepositoryImp interface {
	ListAllProducts() ([]*domain.Product, error)
	FindProductById(id string) (*domain.Product, error)
}

type ProductsRepository struct {
	repository ProductsRepositoryImp
}

func NewListProduct(repo ProductsRepositoryImp) *ProductsRepository {
	return &ProductsRepository{repository: repo}
}

type DTO struct {
	id   string
	name string
}

func (l *ProductsRepository) List() ([]*DTO, error) {
	slog.Info("[products] list")
	l.repository.ListAllProducts()
	return []*DTO{}, nil
}
