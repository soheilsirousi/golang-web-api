package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func AllUsersHandler(c *gin.Context) {
	c.Header("content-type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  "all users here!",
		"errors":  "",
	})
}

func CreateUserHandler(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"result":  "user created!",
		"errors":  "",
	})
}

func GetUserByUsernameHandler(c *gin.Context) {
	username, _ := c.Params.Get("username")
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  fmt.Sprintf("%s is here!", username),
		"errors":  "",
	})
}

func GetUserByIdHandler(c *gin.Context) {
	id, _ := c.Params.Get("id")
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  fmt.Sprintf("user %s is here!", id),
		"errors":  "",
	})
}
