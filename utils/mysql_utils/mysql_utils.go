package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/sufimalek/bookstore_users-api/utils/errors"
)

var (
	noRowsExists = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), noRowsExists) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewIntervalServerError("error in db")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewIntervalServerError("error processing request")
}
