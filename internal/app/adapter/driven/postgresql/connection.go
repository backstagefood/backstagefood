package postgresql

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
)

type sqlDb struct {
	SqlClient *sql.DB
}

func New() *sqlDb {
	connStr := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@" + os.Getenv("DB_HOST") + "/" + os.Getenv("DB_NAME") + "?sslmode=disable"

	client, err := sql.Open(os.Getenv("DB_DRIVER"), connStr)
	if err != nil {
		slog.Error("error connect with database", err.Error(), err)
		panic(err)
	}

	if err = client.Ping(); err != nil {
		slog.Error("error ping database", err.Error(), err)
		panic(err)
	}

	return &sqlDb{SqlClient: client}
}
