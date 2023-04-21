package pq_error

import (
	"github.com/KodjoTouglo/bookstore_users/utils/errors"
	"github.com/lib/pq"
	"strings"
)

const errorNoRows = "no rows in result set"

func ParseError(err error) *errors.APIError {
	sqlErr, ok := err.(*pq.Error)
	if ok {
		switch sqlErr.Code {
		case "23505":
			return errors.BadRequestError("data already exist")
		default:
			return errors.InternalServerError("error executing SQL statement", sqlErr.Error())
		}
	} else if strings.Contains(err.Error(), errorNoRows) {
		return errors.NotFoundError("no record matching given id")
	} else {
		return errors.InternalServerError("error parsing database response")
	}
}
