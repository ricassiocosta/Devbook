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
func (u *User) Prepare(step string) error {
	if err := u.validate(step); err != nil {
		return err
	}

	u.format()
	return nil
}

func (u *User) validate(step string) error {
	if u.Name == "" {
		return errors.New("user name is mandatory")
	}

	if u.Username == "" {
		return errors.New("username is mandatory")
	}

	if u.Email == "" {
		return errors.New("user email is mandatory")
	}

	if step == "register" && u.Password == "" {
		return errors.New("user password is mandatory")
	}

	return nil
}

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Username = strings.TrimSpace(u.Username)
	u.Email = strings.TrimSpace(u.Email)
}
