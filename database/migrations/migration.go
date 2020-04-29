package main

import (
	"github.com/Shuto-san/go-rest-api-mvc-sample/database/configs"
	"github.com/Shuto-san/go-rest-api-mvc-sample/domain/users"
)

func main() {
	db := configs.Client
	defer db.Close()

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&users.User{})
}
