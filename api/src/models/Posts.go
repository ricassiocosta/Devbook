package models

import (
	"errors"
	"strings"
	"time"
)

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

var (
	errMissingPostTitle   = errors.New("post title cannot be empty")
	errMissingPostContent = errors.New("post content cannot be empty")
)

// Prepare validates and format the post fields
func (p *Post) Prepare() error {
	p.Title = strings.TrimSpace(p.Title)
	if len(p.Title) == 0 {
		return errMissingPostTitle
	}

	p.Content = strings.TrimSpace(p.Content)
	if len(p.Content) == 0 {
		return errMissingPostContent
	}
	return nil
}
