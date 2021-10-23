package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/sufimalek/bookstore_users-api/domain/users"
	"github.com/sufimalek/bookstore_users-api/services"
	"github.com/sufimalek/bookstore_users-api/utils/errors"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	//Error handle
	// 	return
	// }

	// //To populate given struct wit json format // One liner to get error and check if condition
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())
	// 	//TODO: handle json error
	// 	return
	// }

	//This will replace above 2 error calls in one
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO
		restErr := errors.NewBadRequestError("Invalide Json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		//TODO: handle error related to create user
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id shoould be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		//TODO: handle error related to create user
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id shoould be a number")
		c.JSON(err.Status, err)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO
		restErr := errors.NewBadRequestError("Invalide Json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "Implement!")
// }
