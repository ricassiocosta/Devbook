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
		`INSERT INTO users (name, username, email, password) VALUES ($1, $2, $3, $4) RETURNING id`,
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
		"SELECT id, name, username, email, created_at FROM users WHERE name LIKE $1 OR username LIKE $2",
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

//Show returns a user from database
func (u Users) Show(userID uint64) (models.User, error) {
	lines, err := u.db.Query(
		"SELECT id, name, username, email, created_at FROM users WHERE id = $1",
		userID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (u Users) Update(ID uint64, user models.User) error {
	statement, err := u.db.Query(
		"UPDATE users SET name = $1, username = $2, email = $3 WHERE id = $4",
		user.Name,
		user.Username,
		user.Email,
		ID,
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	return nil
}

func (u Users) Delete(ID uint64) error {
	statement, err := u.db.Query(
		"DELETE FROM users WHERE id = $1",
		ID,
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	return nil
}
