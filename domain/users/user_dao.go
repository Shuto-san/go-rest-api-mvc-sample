package users

import (
	"github.com/Shuto-san/go-rest-api-mvc-sample/database/configs"
	"github.com/Shuto-san/go-rest-api-mvc-sample/utils/errors"
)

func (user *User) Get(userId int64) *errors.RestErr {
	db := configs.Client
	defer db.Close()

	if err := db.First(user, userId).Error; err != nil {
		return errors.NewInternalServerError("record not found")
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	db := configs.Client
	defer db.Close()

	if err := db.Create(&user).Error; err != nil {
		return errors.NewInternalServerError("save error")
	}
	return nil
}
