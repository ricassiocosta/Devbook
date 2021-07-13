package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

// Search gets all users that matches with an given filter
func (u Users) Search(nameOrUsername string) ([]models.User, error) {
	nameOrUsername = fmt.Sprintf("%%%s%%", nameOrUsername)

	lines, err := u.db.Query(
		"select id, name, username, email, created_at from users where name LIKE $1 or username LIKE $2",
		nameOrUsername,
		nameOrUsername,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
