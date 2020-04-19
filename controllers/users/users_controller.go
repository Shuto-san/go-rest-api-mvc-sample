package users

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "github.com/Shuto-san/go-rest-api-mvc-sample/domain/users"
  "github.com/Shuto-san/go-rest-api-mvc-sample/services"
  "github.com/Shuto-san/go-rest-api-mvc-sample/utils/errors"
  // "io/ioutil"
  // "fmt"
  // "encoding/json"
)


func GetUser(c *gin.Context) {
  c.String(http.StatusNotImplemented, "implement me")
}

func SearchUser(c *gin.Context) {
  c.String(http.StatusNotImplemented, "implement me")
}

func CreateUser(c *gin.Context) {
  var user users.User
  // bytes, err := ioutil.ReadAll(c.Request.Body)
  // if err != nil {
  //   return
  // }
  // if err := json.Unmarshal(bytes, &user); err != nil {
  //     fmt.Println(err.Error())
  //     return
  // }
  //
  // ↓↓↓↓↓↓↓↓↓↓replace
  if err := c.ShouldBindJSON(&user); err != nil {
    restErr := errors.NewBadRequestError("invalid json body")
    c.JSON(restErr.Status, restErr)
    return
  }

  result, saveErr := services.CreateUser(user)
  if saveErr != nil {
    c.JSON(saveErr.Status, saveErr)
    return
  }
  c.JSON(http.StatusCreated, result)
}
