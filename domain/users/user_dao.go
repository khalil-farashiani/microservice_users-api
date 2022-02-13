package users

import (
	"fmt"
	"github.com/khalil-farashiani/microservice_users-api/utils/date_utils"
	"github.com/khalil-farashiani/microservice_users-api/utils/errors"
)

var UserDB = make(map[int64]*User)

func (user *User) Get() *errors.RestErr {
	result := UserDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user with %d id not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	if current := UserDB[user.Id]; current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user with %d already exists", user.Id))
	}
	user.DateCreated = date_utils.GetNowString()
	UserDB[user.Id] = user
	return nil
}
