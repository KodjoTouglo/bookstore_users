package users

import (
	"database/sql"
	"github.com/KodjoTouglo/bookstore_users/storage/postgres/users_db"
	"github.com/KodjoTouglo/bookstore_users/utils/date_utils"
	"github.com/KodjoTouglo/bookstore_users/utils/errors"
	"github.com/KodjoTouglo/bookstore_users/utils/pq_error"
)

const (
	//errorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES($1, $2, $3, $4) RETURNING id;"
	queryGetUserById = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = $1;"
	queryUpdateUser  = "UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE id = $4;"
	queryDeleteUser  = "DELETE FROM users WHERE id = $1;"
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
		return pq_error.ParseError(err)
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
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return pq_error.ParseError(getErr)
	}
	return nil
}

func (user *User) Update() *errors.APIError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return pq_error.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.APIError {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)
	if _, err = stmt.Exec(user.Id); err != nil {
		return pq_error.ParseError(err)
	}
	return nil
}
