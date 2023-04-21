package users

import (
	"github.com/KodjoTouglo/bookstore_users/domain/users"
	"github.com/KodjoTouglo/bookstore_users/services"
	"github.com/KodjoTouglo/bookstore_users/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getUserId(userIdParam string) (int64, *errors.APIError) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.BadRequestError("Invalid user `id`, id should be a number")
	}
	return userId, nil
}

func GetUser(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.StatusCode, idErr)
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

func UpdateUser(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.StatusCode, idErr)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		APIErr := errors.BadRequestError("Invalid json body.")
		c.JSON(APIErr.StatusCode, APIErr)
		return
	}
	user.Id = userId
	isPartial := c.Request.Method == http.MethodPatch
	result, updateErr := services.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.StatusCode, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.StatusCode, idErr)
		return
	}
	if err := services.DeleteUser(userId); err != nil {
		c.JSON(idErr.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
