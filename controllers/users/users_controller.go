package users

import (
	"net/http"
	"strconv"

	"github.com/Shuto-san/go-rest-api-mvc-sample/domain/users"
	"github.com/Shuto-san/go-rest-api-mvc-sample/services"
	"github.com/Shuto-san/go-rest-api-mvc-sample/utils/errors"
	"github.com/gin-gonic/gin"
	// "io/ioutil"
	// "fmt"
	// "encoding/json"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
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
		// restErr := errors.RestErr{
		//   Message: "invalid json body",
		//   Status: http.StatusBadRequest,
		//   Error: "bad request"
		// }
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
