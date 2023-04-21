package users

import (
	"github.com/KodjoTouglo/bookstore_users/utils/errors"
	"strings"
	"time"
)

type User struct {
	Id          int64     `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	DateCreated time.Time `json:"date_created"`
}

func (user *User) ValidateEmail() *errors.APIError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.BadRequestError("Invalid email address.")
	}
	return nil
}
