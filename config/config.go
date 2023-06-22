package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type Storage struct {
	conn *sql.DB
}

type Scanner interface {
	Scan(dest ...interface{}) error
}

func (s *Storage) Ping() error {
	return s.conn.Ping()
}

func dsn() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
}

func NewDB() (*Storage, error) {

	// Capture connection properties.
	conn, err := sql.Open("mysql", dsn())
	if err != nil {
		return nil, fmt.Errorf("could not open sql: %w", err)
	}

	return &Storage{
		conn: conn,
	}, nil
}
