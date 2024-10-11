package service

import (
	"fmt"

	"github.com/backstagefood/backstagefood/internal/domain"
)

type ProductInterface interface {
	ListProducts() ([]*domain.Product, error)
	FindProductById(id string) (*domain.Product, error)
}

type ProductService struct {
	productRepository ProductInterface
}

type ProductDTO struct {
	Id          string  `json:"id"`
	Description string  `json:"description"`
	Ingredients string  `json:"ingredients"`
	Price       float64 `json:"price"`
}

func NewProductService(repository ProductInterface) *ProductService {
	return &ProductService{productRepository: repository}
}

func (p *ProductService) GetProductById(id string) (*ProductDTO, error) {
	product, err := p.productRepository.FindProductById(id)
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

func (p *ProductService) GetProducts() ([]*ProductDTO, error) {
	productList, err := p.productRepository.ListProducts()
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
