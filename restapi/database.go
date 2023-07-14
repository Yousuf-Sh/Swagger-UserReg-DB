package restapi

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	// Replace the database connection details with your own configuration
	db, err := sql.Open("postgres", "postgres://postgres:ys192001@localhost:5432/user_registeration?sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping the database: %w", err)
	}

	return db, nil
}
