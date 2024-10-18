package repositories

import (
	"database/sql"
	"github.com/backstagefood/backstagefood/internal/core/domain"
	portRepository "github.com/backstagefood/backstagefood/internal/core/ports/repositories"
)

type productRepository struct {
	sqlClient *sql.DB
}

func NewProductRepository(database *ApplicationDatabase) portRepository.Product {
	return &productRepository{
		sqlClient: database.sqlClient,
	}
}

func (s *productRepository) ListProducts(description string) ([]*domain.Product, error) {
	query := "SELECT a.id, a.id_category, a.description, a.ingredients, a.price, a.created_at, a.updated_at, b.id, b.description FROM products a, product_categories b where a.id_category = b.id AND a.description ILIKE '%' || $1 || '%'"
	stmt, err := s.sqlClient.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(description)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := make([]*domain.Product, 0)
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(
			&product.ID,
			&product.IDCategory,
			&product.Description,
			&product.Ingredients,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.ProductCategory.ID,
			&product.ProductCategory.Description); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (s *productRepository) FindProductById(id string) (*domain.Product, error) {
	query := "SELECT a.id, a.id_category, a.description, a.ingredients, a.price, a.created_at, a.updated_at, b.id, b.description FROM products a, product_categories b where a.id_category = b.id AND a.id = $1"
	stmt, err := s.sqlClient.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	var product domain.Product
	err = stmt.QueryRow(id).Scan(
		&product.ID,
		&product.IDCategory,
		&product.Description,
		&product.Ingredients,
		&product.Price,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.ProductCategory.ID,
		&product.ProductCategory.Description)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (s *productRepository) CreateProduct(product *domain.Product) (*domain.Product, error) {
	query := "INSERT INTO products (id, id_category, description, ingredients, created_at, updated_at, price) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	stmt, err := s.sqlClient.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(
		product.ID,
		product.IDCategory,
		product.Description,
		product.Ingredients,
		product.CreatedAt,
		product.UpdatedAt,
		product.Price).Scan(&product.ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productRepository) GetCategoryID(categoryDescription string) (string, error) {
	query := "SELECT id FROM product_categories WHERE description = $1"
	stmt, err := s.sqlClient.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return "", err
	}
	var categoryID string
	err = stmt.QueryRow(categoryDescription).Scan(&categoryID)
	if err != nil {
		return "", err
	}
	return categoryID, nil
}
