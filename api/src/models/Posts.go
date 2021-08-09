package models

import "time"

// Post represents a post created by an user
type Post struct {
	ID             uint64    `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorID       uint64    `json:"author_id,omitempty"`
	AuthorUsername string    `json:"author_username,omitempty"`
	Likes          int64     `json:"likes"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}
