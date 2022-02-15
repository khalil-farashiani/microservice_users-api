package users

import (
	"github.com/khalil-farashiani/microservice_users-api/datasources/mysql/users_db"
	"github.com/khalil-farashiani/microservice_users-api/utils/date_utils"
	"github.com/khalil-farashiani/microservice_users-api/utils/errors"
	"github.com/khalil-farashiani/microservice_users-api/utils/mysql_utils"
)

const (
	errorNoRows     = "no rows in result set"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(&user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		mysql_utils.ParsError(getErr)
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParsError(saveErr)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParsError(err)
	}

	user.Id = userId
	return nil
}
