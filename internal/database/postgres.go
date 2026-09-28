package database

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/url"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func OpenPostgres() (*sql.DB, error) {
	dsn, err := connectionString()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("abrir conexão PostgreSQL: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("validar conexão PostgreSQL: %w", err)
	}

	return db, nil
}

func connectionString() (string, error) {
	if dsn := os.Getenv("DATABASE_URL"); dsn != "" {
		return dsn, nil
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DB")
	if user == "" || password == "" || databaseName == "" {
		return "", fmt.Errorf("configure DATABASE_URL ou POSTGRES_USER, POSTGRES_PASSWORD e POSTGRES_DB")
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	query := url.Values{}
	query.Set("sslmode", "disable")
	return (&url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(user, password),
		Host:     net.JoinHostPort(host, port),
		Path:     databaseName,
		RawQuery: query.Encode(),
	}).String(), nil
}
