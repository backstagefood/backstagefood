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

func (p *ProductService) GetProductById(id string) (*domain.Product, error) {
	product, err := p.productRepository.FindProductById(id)
	if err != nil {
		slog.Error("error:", err)
		return nil, fmt.Errorf("product with id: %s not found", id)
	}

	return &domain.Product{
		ID:              product.ID,
		Description:     product.Description,
		Ingredients:     product.Ingredients,
		Price:           product.Price,
		IDCategory:      product.ProductCategory.ID,
		ProductCategory: product.ProductCategory,
	}, nil
}

func (p *ProductService) GetProducts(description string) ([]*domain.Product, error) {
	productList, err := p.productRepository.ListProducts(description)
	if err != nil {
		slog.Error("error:", err)
		return []*domain.Product{}, fmt.Errorf("products not found")
	}

	var output []*domain.Product
	for _, product := range productList {
		output = append(output, &domain.Product{
			ID:              product.ID,
			Description:     product.Description,
			Ingredients:     product.Ingredients,
			Price:           product.Price,
			IDCategory:      product.ProductCategory.ID,
			ProductCategory: product.ProductCategory,
		})
	}
	return output, nil
}

func (p *ProductService) CreateProduct(productDTO *portService.ProductDTO) (*domain.Product, error) {
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

	return &domain.Product{
		ID:          createdProduct.ID,
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

func (p *ProductService) UpdateProduct(productDTO *portService.ProductDTO) (*domain.Product, error) {
	existingProduct, err := p.productRepository.FindProductById(productDTO.Id)
	if err != nil {
		return nil, err
	}

	productToUpdate := &domain.Product{
		ID:          productDTO.Id,
		IDCategory:  existingProduct.IDCategory,
		Description: existingProduct.Description,
		Ingredients: existingProduct.Ingredients,
		Price:       existingProduct.Price,
		UpdatedAt:   domain.GetNow(),
	}

	if productDTO.Description != "" {
		productToUpdate.Description = productDTO.Description
	}

	if productDTO.Ingredients != "" {
		productToUpdate.Ingredients = productDTO.Ingredients
	}

	if productDTO.Price != 0 {
		productToUpdate.Price = productDTO.Price
	}

	if productDTO.Category != "" {
		categoryID, err := p.GetCategoryID(productDTO.Category)
		if err != nil {
			return nil, err
		}
		productToUpdate.IDCategory = categoryID
	}

	//Validate Product
	if err := domain.ValidateProduct(productToUpdate); err != nil {
		return nil, err
	}

	updatedProduct, err := p.productRepository.UpdateProduct(productToUpdate)
	if err != nil {
		slog.Error("error:", err)
		return nil, fmt.Errorf("error updating product")
	}

	return &domain.Product{
		ID:          updatedProduct.ID,
		Description: updatedProduct.Description,
		Ingredients: updatedProduct.Ingredients,
		Price:       updatedProduct.Price,
	}, nil
}

// DeleteProduct deletes a product from the database
func (p *ProductService) DeleteProduct(productID string) error {
	if err := p.productRepository.DeleteProduct(productID); err != nil {
		return err
	}
	return nil
}

// GetCategories returns a list of all product categories
func (p *ProductService) GetCategories() ([]*domain.ProductCategory, error) {
	categories, err := p.productRepository.GetCategories()
	if err != nil {
		return []*domain.ProductCategory{}, err
	}

	return categories, nil
}
