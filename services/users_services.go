package services

import (
	"github.com/Shuto-san/go-rest-api-mvc-sample/domain/users"
	"github.com/Shuto-san/go-rest-api-mvc-sample/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	if userId <= 0 {
		return nil, errors.NewBadRequestError("invalid user id")
	}
	user := &users.User{}
	if err := user.Get(userId); err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	// depend on users package
	// if err := users.Validate(&user); err !=nill {
	//     return err
	// }

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
