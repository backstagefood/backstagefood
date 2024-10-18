package services

import (
	"fmt"
	"log/slog"

	"github.com/backstagefood/backstagefood/internal/core/domain"
	portRepository "github.com/backstagefood/backstagefood/internal/core/ports/repositories"
	portService "github.com/backstagefood/backstagefood/internal/core/ports/services"
	"github.com/google/uuid"
)

type ProductService struct {
	productRepository portRepository.Product
}

func NewProductService(repository portRepository.Product) portService.Product {
	return &ProductService{productRepository: repository}
}

func (p *ProductService) GetProductById(id string) (*portService.ProductDTO, error) {
	product, err := p.productRepository.FindProductById(id)
	if err != nil {
		slog.Error("error:", err)
		return nil, fmt.Errorf("product with id: %s not found", id)
	}

	return &portService.ProductDTO{
		Id:          product.ID,
		Description: product.Description,
		Ingredients: product.Ingredients,
		Price:       product.Price,
		IDCategory:  product.ProductCategory.ID,
		Category:    product.ProductCategory.Description,
	}, nil
}

func (p *ProductService) GetProducts(description string) ([]*portService.ProductDTO, error) {
	productList, err := p.productRepository.ListProducts(description)
	if err != nil {
		slog.Error("error:", err)
		return []*portService.ProductDTO{}, fmt.Errorf("products not found")
	}

	var output []*portService.ProductDTO
	for _, product := range productList {
		output = append(output, &portService.ProductDTO{
			Id:          product.ID,
			Description: product.Description,
			Ingredients: product.Ingredients,
			Price:       product.Price,
			IDCategory:  product.ProductCategory.ID,
			Category:    product.ProductCategory.Description,
		})
	}
	return output, nil
}

func (p *ProductService) CreateProduct(productDTO *portService.ProductDTO) (*portService.ProductDTO, error) {
	productToCreate := &domain.Product{
		ID:          uuid.New().String(),
		IDCategory:  productDTO.IDCategory,
		Description: productDTO.Description,
		Ingredients: productDTO.Ingredients,
		Price:       productDTO.Price,
		CreatedAt:   domain.GetNow(),
		UpdatedAt:   domain.GetNow(),
	}

	//Validate Product
	if err := domain.ValidateProduct(productToCreate); err != nil {
		return nil, err
	}

	createdProduct, err := p.productRepository.CreateProduct(productToCreate)
	if err != nil {
		return nil, fmt.Errorf("error creating product")
	}

	return &portService.ProductDTO{
		Id:          createdProduct.ID,
		Description: createdProduct.Description,
		Ingredients: createdProduct.Ingredients,
		Price:       createdProduct.Price,
	}, nil
}

func (p *ProductService) GetCategoryID(categoryName string) (string, error) {
	categoryID, err := p.productRepository.GetCategoryID(categoryName)
	if err != nil {
		return "", fmt.Errorf("category not found")
	}
	return categoryID, nil
}
