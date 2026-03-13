package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// OpenDB centraliza la apertura de conexión para evitar acoplar handlers
// o servicios de aplicación a detalles de infraestructura.
func OpenDB() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// DSN local por defecto para facilitar ejecución del ejemplo.
		dsn = "postgres://postgres:postgres@localhost:5432/aliado_ddd?sslmode=disable"
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open postgres: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	return db, nil
}
