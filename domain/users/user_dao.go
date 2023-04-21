package users

import (
	"database/sql"
	"fmt"
	"github.com/KodjoTouglo/bookstore_users/storage/postgres/users_db"
	"github.com/KodjoTouglo/bookstore_users/utils/date_utils"
	"github.com/KodjoTouglo/bookstore_users/utils/errors"
	"strings"
)

const queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES($1, $2, $3, $4) " +
	"RETURNING id;"

var usersDB = make(map[int64]*User)

func (user *User) Save() *errors.APIError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
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
		if strings.Contains(err.Error(), "users_email_key") {
			return errors.BadRequestError(fmt.Sprintf("email `%s` already exist", user.Email))
		}
		return errors.InternalServerError(err.Error())
	}
	user.Id = userId
	return nil
}

func (user *User) Get() *errors.APIError {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
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
