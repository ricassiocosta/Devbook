package models

import (
	"api/src/utils/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

var (
	errMissingName     = errors.New("name is a mandatory field and can't be blank")
	errMissingUsername = errors.New("username is a mandatory field and can't be blank")
	errMissingEmail    = errors.New("email is a mandatory field and can't be blank")
	errMissingPassword = errors.New("password is a mandatory field and can't be blank")
	errInvalidEmail    = errors.New("email must have a valid format")
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

	if err := u.format(step); err != nil {
		return err
	}

	return nil
}

func (u *User) validate(step string) error {
	if u.Name == "" {
		return errMissingName
	}

	if u.Username == "" {
		return errMissingUsername
	}

	if u.Email == "" {
		return errMissingEmail
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errInvalidEmail
	}

	if step == "register" && u.Password == "" {
		return errMissingPassword
	}

	return nil
}

func (u *User) format(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Username = strings.TrimSpace(u.Username)
	u.Email = strings.TrimSpace(u.Email)

	if step == "register" {
		hashedPassword, err := security.Hash(u.Password)
		if err != nil {
			return err
		}

		u.Password = string(hashedPassword)
	}

	return nil
}
