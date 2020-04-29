package configs

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	mysql_username = "mysql_username"
	mysql_password = "mysql_password"
	mysql_host     = "mysql_host"
	mysql_dbname   = "mysql_dbname"
)

var (
	Client *gorm.DB

	username = os.Getenv(mysql_username)
	password = os.Getenv(mysql_password)
	host     = os.Getenv(mysql_host)
	dbname   = os.Getenv(mysql_dbname)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, dbname,
	)

	var err error
	Client, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	Client.LogMode(true)
}
