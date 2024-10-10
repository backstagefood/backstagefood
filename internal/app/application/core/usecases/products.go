package usecases

import (
	"fmt"
	"github.com/backstagefood/backstagefood/internal/app/domain"
)

type ProductsRepositoryImp interface {
	ListAllProducts() ([]*domain.Product, error)
	FindProductById(id string) (*domain.Product, error)
}

type ProductsRepository struct {
	repository ProductsRepositoryImp
}

type ProductDTO struct {
	Id          string  `json:"id"`
	Description string  `json:"description"`
	Ingredients string  `json:"ingredients"`
	Price       float64 `json:"price"`
}

func NewProductsRepository(repository ProductsRepositoryImp) *ProductsRepository {
	return &ProductsRepository{repository: repository}
}

func (p *ProductsRepository) GetProductById(id string) (*ProductDTO, error) {
	product, err := p.repository.FindProductById(id)
	if err != nil {
		return nil, fmt.Errorf("product with id: %s not found", id)
	}
	return &ProductDTO{
		Id:          product.ID,
		Description: product.Description,
		Ingredients: product.Ingredients,
		Price:       product.Price,
	}, nil
}

func (p *ProductsRepository) GetProducts() ([]*ProductDTO, error) {
	productList, err := p.repository.ListAllProducts()
	if err != nil {
		return []*ProductDTO{}, fmt.Errorf("products not found")
	}
	var output []*ProductDTO
	for _, product := range productList {
		output = append(output, &ProductDTO{
			Id:          product.ID,
			Description: product.Description,
			Ingredients: product.Ingredients,
			Price:       product.Price,
		})
	}
	return output, nil
}
