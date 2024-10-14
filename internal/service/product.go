package service

import (
	"fmt"
	"log/slog"

	"github.com/backstagefood/backstagefood/internal/domain"
	"github.com/google/uuid"
)

type ProductInterface interface {
	ListProducts(description string) ([]*domain.Product, error)
	FindProductById(id string) (*domain.Product, error)
	CreateProduct(product *domain.Product) (*domain.Product, error)
	GetCategoryID(categoryName string) (string, error)
}

type ProductService struct {
	productRepository ProductInterface
}

type ProductDTO struct {
	Id          string  `json:"id"`
	Description string  `json:"description"`
	Ingredients string  `json:"ingredients"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	IDCategory  string  `json:"category_id"`
}

func NewProductService(repository ProductInterface) *ProductService {
	return &ProductService{productRepository: repository}
}

func (p *ProductService) GetProductById(id string) (*ProductDTO, error) {
	product, err := p.productRepository.FindProductById(id)
	if err != nil {
		slog.Error("error:", err)
		return nil, fmt.Errorf("product with id: %s not found", id)
	}

	return &ProductDTO{
		Id:          product.ID,
		Description: product.Description,
		Ingredients: product.Ingredients,
		Price:       product.Price,
		Category:    product.ProductCategory.Description,
	}, nil
}

func (p *ProductService) GetProducts(description string) ([]*ProductDTO, error) {
	productList, err := p.productRepository.ListProducts(description)
	if err != nil {
		slog.Error("error:", err)
		return []*ProductDTO{}, fmt.Errorf("products not found")
	}

	var output []*ProductDTO
	for _, product := range productList {
		output = append(output, &ProductDTO{
			Id:          product.ID,
			Description: product.Description,
			Ingredients: product.Ingredients,
			Price:       product.Price,
			Category:    product.ProductCategory.Description,
		})
	}
	return output, nil
}

func (p *ProductService) CreateProduct(productDTO *ProductDTO) (*ProductDTO, error) {
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

	return &ProductDTO{
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
