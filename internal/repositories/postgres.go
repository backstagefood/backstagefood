package repositories

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
)

type ApplicationDatabase struct {
	sqlClient *sql.DB
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

	return &ApplicationDatabase{sqlClient: client}
}

func (s *ApplicationDatabase) Client() *sql.DB {
	return s.sqlClient
}

func (s *ApplicationDatabase) DataBaseHeatlh() error {
	return s.sqlClient.Ping()
}
