package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/backstagefood/backstagefood/internal/app/domain"
	_ "github.com/lib/pq"
)

type SqlDb struct {
	SqlClient *sql.DB
}

func New() *SqlDb {
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

	return &SqlDb{SqlClient: client}
}

func (s *SqlDb) ListAllProducts() ([]*domain.Product, error) {
	statement := "SELECT id, id_category, description, ingredients, price, created_at, updated_at FROM products"
	rows, err := s.SqlClient.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := make([]domain.Product, 0)
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.IDCategory, &product.Description, &product.Ingredients, &product.Price, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return []*domain.Product{}, nil
}

func (s *SqlDb) FindProductById(id string) (*domain.Product, error) {
	stmt, err := s.SqlClient.Prepare("SELECT id, id_category, description, ingredients, price, created_at, updated_at FROM products WHERE id = $1")
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	var product domain.Product
	err = stmt.QueryRow(id).Scan(&product.ID, &product.IDCategory, &product.Description, &product.Ingredients, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
