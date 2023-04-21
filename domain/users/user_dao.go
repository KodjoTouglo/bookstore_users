package users

import (
	"database/sql"
	"fmt"
	"github.com/KodjoTouglo/bookstore_users/storage/postgres/users_db"
	"github.com/KodjoTouglo/bookstore_users/utils/date_utils"
	"github.com/KodjoTouglo/bookstore_users/utils/errors"
	"github.com/lib/pq"
	"strings"
)

const (
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES($1, $2, $3, $4) RETURNING id;"
	queryGetUserById = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = $1"
)

func (user *User) Save() *errors.APIError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.InternalServerError("", err.Error())
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)
	user.DateCreated = date_utils.GetNow()
	var userId int64
	err = stmt.QueryRow(user.FirstName, user.LastName, user.Email, user.DateCreated).Scan(&userId)
	if err != nil {
		sqlErr, ok := err.(*pq.Error)
		if ok {
			switch sqlErr.Code {
			case "23505":
				return errors.BadRequestError(fmt.Sprintf("email `%s` already exist", user.Email))
			default:
				return errors.InternalServerError("error executing SQL statement", sqlErr.Error())
			}
		} else {
			return errors.InternalServerError(err.Error(), "error creating user")
		}
	}
	user.Id = userId
	return nil
}

func (user *User) Get() *errors.APIError {
	stmt, err := users_db.Client.Prepare(queryGetUserById)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.NotFoundError(fmt.Sprintf("User `%d` not found", user.Id), err.Error())
		}
		return errors.InternalServerError(fmt.Sprintf("error when trying to get user `%d`", user.Id), err.Error())
	}
	return nil
}
