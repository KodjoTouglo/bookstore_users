package users

import (
	"fmt"
	"github.com/KodjoTouglo/bookstore_users/utils/date_utils"
	"github.com/KodjoTouglo/bookstore_users/utils/errors"
)

var usersDB = make(map[int64]*User)

func (user *User) Save() *errors.APIError {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.BadRequestError(fmt.Sprintf("Email `%s` already registered.", user.Email))
		}
		return errors.BadRequestError(fmt.Sprintf("User `%d` already exist.", user.Id))
	}
	user.DateCreated = date_utils.GetNowString()
	usersDB[user.Id] = user
	return nil
}

func (user *User) Get() *errors.APIError {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NotFoundError(fmt.Sprintf("User %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}
