package users

import (
	"github.com/KodjoTouglo/bookstore_users/domain/users"
	"github.com/KodjoTouglo/bookstore_users/services"
	"github.com/KodjoTouglo/bookstore_users/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.BadRequestError("Invalid user `id`, id should be a number")
		c.JSON(err.StatusCode, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.StatusCode, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		APIErr := errors.BadRequestError("Invalid json body.")
		c.JSON(APIErr.StatusCode, APIErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.StatusCode, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
