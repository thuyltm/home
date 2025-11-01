package repository

import (
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByID(id int) (string, error) {
	var name string
	err := r.db.QueryRow("SELECT name FROM users WHERE id = ?", id).Scan(&name)
	if err != nil {
		return "", fmt.Errorf("failed to get user: %w", err)
	}
	return name, nil
}
