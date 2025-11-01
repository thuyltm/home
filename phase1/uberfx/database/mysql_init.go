package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
	"go.uber.org/fx"
)

// NewMySQLDB creates a new MySQL database connection.
func NewMySQLDB(lifecycle fx.Lifecycle) (*sql.DB, error) {
	// Replace with your actual database connection string
	dataSourceName := "me:home@tcp(127.0.0.1:3306)/mydb?parseTime=true"
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Configure connection pool settings
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Add lifecycle hooks for graceful shutdown
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return db.PingContext(ctx) // Verify connection on startup
		},
		OnStop: func(ctx context.Context) error {
			return db.Close() // Close connection on shutdown
		},
	})

	return db, nil
}
