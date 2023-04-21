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

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.APIError) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(userId int64) *errors.APIError {
	user := &users.User{Id: userId}
	return user.Delete()
}
