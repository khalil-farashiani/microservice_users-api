package mysql_utils

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/khalil-farashiani/microservice_users-api/utils/errors"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParsError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)

	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(
				fmt.Sprintf("no recored matching given id"),
			)
		}
		return errors.NewInternalServerError("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error parsing database request")
}
