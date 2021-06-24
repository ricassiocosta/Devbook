package repositories

import (
	"api/src/models"
	"database/sql"
)

// Users represents an user repository
type Users struct {
	db *sql.DB
}

// NewUsersRepositories
func NewUsersRepositories(db *sql.DB) *Users {
	return &Users{db}
}

// Create insert an user in the database
func (u Users) Create(user models.User) (uint64, error) {
	return 0, nil
}
