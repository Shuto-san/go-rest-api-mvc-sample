package services

import (
  "github.com/Shuto-san/go-rest-api-mvc-sample/domain/users"
  "github.com/Shuto-san/go-rest-api-mvc-sample/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

  return &user, nil
}
