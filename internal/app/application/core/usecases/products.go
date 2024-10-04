package usecases

import (
	"github.com/backstagefood/backstagefood/internal/app/domain"
)

type ListProductsImp interface {
	ListAll() ([]*domain.Product, error)
	FindById(id string) ([]*domain.Product, error)
}

type ListProducts struct {
	repositoryList ListProductsImp
}

func NewListProduct(repo ListProductsImp) *ListProducts {
	return &ListProducts{repositoryList: repo}
}

type DTO struct {
	id   string
	name string
}

func (l *ListProducts) List() ([]*DTO, error) {
	l.repositoryList.ListAll()
	return []*DTO{}, nil
}
