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

// NewUsersRepository
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create insert an user in the database
func (u Users) Create(user models.User) (uint64, error) {
	lastInsertID := 0
	err := u.db.QueryRow(
		`INSERT INTO users (name, username, email, password) VALUES ($1, $2, $3, $4) RETURNING id`,
		user.Name,
		user.Username,
		user.Email,
		user.Password,
	).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}
	defer u.db.Close()

	return uint64(lastInsertID), nil
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

// Update enable users update they informations
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

// Delete enable users to delete their accounts
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

// GetByEmail returns an user who email matches to a given one
func (u Users) GetByEmail(email string) (models.User, error) {
	line, err := u.db.Query(
		"SELECT id, password FROM users WHERE email = $1",
		email,
	)
	if err != nil {
		return models.User{}, err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Follow enable users to follow anothers
func (u Users) Follow(userID, followerID uint64) error {
	statement, err := u.db.Query(
		"INSERT INTO followers (user_id, follower_id) VALUES ($1, $2) ON CONFLICT DO NOTHING",
		userID,
		followerID,
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	return nil
}

// Follow enable users to stop follow anothers
func (u Users) Unfollow(userID, followerID uint64) error {
	statement, err := u.db.Query(
		"DELETE FROM followers WHERE user_id = $1 AND follower_id = $2",
		userID,
		followerID,
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	return nil
}

// GetFollowers returns all user followers
func (u Users) GetFollowers(userID uint64) ([]models.User, error) {
	lines, err := u.db.Query(
		`SELECT u.id, u.name, u.username, u.email, u.created_at
			FROM users AS u INNER JOIN followers AS f
			ON u.id = f.follower_id
			WHERE f.user_id = $1
		`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err := lines.Scan(
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

// GetFollowing returns who an user is following
func (u Users) GetFollowing(userID uint64) ([]models.User, error) {
	lines, err := u.db.Query(
		`SELECT u.id, u.name, u.username, u.email, u.created_at
			FROM users AS u INNER JOIN followers AS f
			ON u.id = f.user_id
			WHERE f.follower_id = $1
		`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err := lines.Scan(
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

// GetPassword returns user's password
func (u Users) GetPassword(userID uint64) (string, error) {
	line, err := u.db.Query(
		"SELECT password FROM users WHERE id = $1",
		userID,
	)
	if err != nil {
		return "", err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

// UpdatePassword updates user's password
func (u Users) UpdatePassword(userID uint64, password string) error {
	statement, err := u.db.Query(
		"UPDATE users SET password = $1 WHERE id = $2",
		password,
		userID,
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	return nil
}
