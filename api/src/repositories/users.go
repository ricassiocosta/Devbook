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
	lasInsertID := 0
	err := u.db.QueryRow(
		`insert into users (name, username, email, password) values ($1, $2, $3, $4) returning id`,
		user.Name,
		user.Username,
		user.Email,
		user.Password,
	).Scan(&lasInsertID)
	if err != nil {
		return 0, err
	}
	defer u.db.Close()

	return uint64(lasInsertID), nil
}
