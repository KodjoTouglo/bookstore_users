package services

import (
	"github.com/KodjoTouglo/bookstore_users/domain/users"
	"github.com/KodjoTouglo/bookstore_users/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.APIError) {
	if err := user.ValidateEmail(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.APIError) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, nil
	}
	return result, nil
}
