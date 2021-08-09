package repositories

import (
	"api/src/models"
	"database/sql"
)

// Posts represents a posts repository
type Posts struct {
	db *sql.DB
}

// NewPostsRepository create a new posts repository
func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

// Create insert a new post in the database
func (p Posts) Create(post models.Post) (uint64, error) {
	lastInsertID := 0
	err := p.db.QueryRow(
		`INSERT INTO posts (title, content, author_id) VALUES ($1, $2, $3) RETURNING id`,
		post.Title,
		post.Content,
		post.AuthorID,
	).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}
	defer p.db.Close()

	return uint64(lastInsertID), nil
}
