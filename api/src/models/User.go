package models

import (
	"errors"
	"strings"
	"time"
)

// User defines the user's data stored in the database
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Prepare will validate and format the received user
func (u *User) Prepare() error {
	if err := u.validate(); err != nil {
		return err
	}

	u.format()
	return nil
}

func (u *User) validate() error {
	if u.Name == "" {
		return errors.New("user name is mandatory")
	}

	if u.Username == "" {
		return errors.New("username is mandatory")
	}

	if u.Email == "" {
		return errors.New("user email is mandatory")
	}

	return nil
}

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Name = strings.TrimSpace(u.Username)
	u.Name = strings.TrimSpace(u.Email)
}
