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

// GetByID get a single post from database
func (p Posts) GetByID(postID uint64) (models.Post, error) {
	line, err := p.db.Query(
		`
		SELECT p.*, u.username FROM
		posts AS p INNER JOIN users AS u
		ON u.id = p.author_id WHERE p.id = $1
		`,
		postID,
	)
	if err != nil {
		return models.Post{}, err
	}
	defer line.Close()

	var post models.Post

	if line.Next() {
		if err = line.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorUsername,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

// GetPosts gets the followers posts
func (p Posts) GetPosts(userID uint64) ([]models.Post, error) {
	lines, err := p.db.Query(
		`
		SELECT DISTINCT p.*, u.username 
		FROM posts AS p INNER JOIN users AS u
		ON u.id = p.author_id 
		INNER JOIN followers AS f ON p.author_id = f.user_id
		WHERE u.id = $1 OR f.follower_id = $2
		ORDER BY 1 DESC
		`,
		userID,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorUsername,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
