package postgresdb

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	POSTGRES_DB      = "POSTGRES_DB"
	POSTGRES_USER    = "POSTGRES_USER"
	POSTGRES_PASS    = "POSTGRES_PASS"
	POSTGRES_PORT    = "POSTGRES_PORT"
	POSTGRES_SSLMODE = "POSTGRES_SSLMODE"
)

func NewPostgresDBConnection() (*sql.DB, error) {
	postgresUser := os.Getenv(POSTGRES_USER)
	postgresPassword := os.Getenv(POSTGRES_PASS)
	postgresPort := os.Getenv(POSTGRES_PORT)
	postgresDB := os.Getenv(POSTGRES_DB)
	postgresSSLMode := os.Getenv(POSTGRES_SSLMODE)

	connStr := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=%s", postgresUser, postgresPassword, postgresPort, postgresDB, postgresSSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
