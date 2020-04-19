package app

import (
  "github.com/Shuto-san/go-rest-api-mvc-sample/controllers/ping"
  "github.com/Shuto-san/go-rest-api-mvc-sample/controllers/users"
)

func mapUrls() {
  router.GET("/ping", ping.Ping)
  router.GET("/users/:user_id", users.GetUser)
  router.GET("/users", users.SearchUser)
  router.POST("/users", users.CreateUser)
}
