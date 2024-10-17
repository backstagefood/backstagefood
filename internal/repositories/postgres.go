package db

import (
	"database/sql"
	"fmt"
	"github.com/backstagefood/backstagefood/internal/domain"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
)

type ApplicationDatabase struct {
	SqlClient *sql.DB
}

func New() *ApplicationDatabase {
	connStr := fmt.Sprintf(
		"%s://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	client, err := sql.Open(os.Getenv("DB_DRIVER"), connStr)
	if err != nil {
		slog.Error("error connect with database", err.Error(), err)
		panic(err)
	}

	if err = client.Ping(); err != nil {
		slog.Error("error ping database", err.Error(), err)
		panic(err)
	}

	return &ApplicationDatabase{SqlClient: client}
}

func (s *ApplicationDatabase) ListProducts(description string) ([]*domain.Product, error) {
	query := "SELECT a.id, a.id_category, a.description, a.ingredients, a.price, a.created_at, a.updated_at, b.id, b.description FROM products a, product_categories b where a.id_category = b.id AND a.description ILIKE '%' || $1 || '%'"
	stmt, err := s.SqlClient.Prepare(query)
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

func (s *ApplicationDatabase) FindProductById(id string) (*domain.Product, error) {
	query := "SELECT a.id, a.id_category, a.description, a.ingredients, a.price, a.created_at, a.updated_at, b.id, b.description FROM products a, product_categories b where a.id_category = b.id AND a.id = $1"
	stmt, err := s.SqlClient.Prepare(query)
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

func (s *ApplicationDatabase) CreateProduct(product *domain.Product) (*domain.Product, error) {
	query := "INSERT INTO products (id, id_category, description, ingredients, created_at, updated_at, price) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	stmt, err := s.SqlClient.Prepare(query)
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

// Get CategoryID returns the category ID based on the category name
func (s *ApplicationDatabase) GetCategoryID(categoryDescription string) (string, error) {
	query := "SELECT id FROM product_categories WHERE description = $1"
	stmt, err := s.SqlClient.Prepare(query)
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
