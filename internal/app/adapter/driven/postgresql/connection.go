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

func (s *SqlDb) ListAll() ([]*domain.Product, error) {
	return []*domain.Product{}, nil
}
