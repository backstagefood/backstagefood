package usecases

import (
	"fmt"
	"log/slog"

	"github.com/backstagefood/backstagefood/internal/app/domain"
)

type ProductsRepositoryImp interface {
	ListAllProducts() ([]*domain.Product, error)
	FindProductById(id string) (*domain.Product, error)
}

type ProductsRepository struct {
	repository ProductsRepositoryImp
}

func NewProducts(repo ProductsRepositoryImp) *ProductsRepository {
	return &ProductsRepository{repository: repo}
}

type DTO struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

func (p *ProductsRepository) GetProductById(id string) (*DTO, error) {
	product, err := p.repository.FindProductById(id)
	if err != nil {
		return nil, fmt.Errorf("produto com id: %s não encontrado", id)
	}

	return &DTO{
		Id:          product.ID,
		Description: product.Description,
	}, nil
}

func (p *ProductsRepository) GetProducts() ([]*DTO, error) {
	slog.Info("[products] list")
	productList, err := p.repository.ListAllProducts()
	if err != nil || len(productList) == 0 {
		return []*DTO{}, fmt.Errorf("produtos não foram encontrado")
	}

	var output []*DTO
	for _, product := range productList {
		var tmpOutput DTO

		tmpOutput.Id = product.ID
		tmpOutput.Description = product.Description

		output = append(output, &tmpOutput)
	}
	return output, nil
}
