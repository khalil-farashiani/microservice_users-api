package services

import (
	"github.com/khalil-farashiani/microservice_users-api/domain/users"
	"github.com/khalil-farashiani/microservice_users-api/utils/crypto_utils"
	"github.com/khalil-farashiani/microservice_users-api/utils/date_utils"
	"github.com/khalil-farashiani/microservice_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMD5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
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

func DeleteUSer(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func Search(status string) ([]users.User, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
