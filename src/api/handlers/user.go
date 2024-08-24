package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soheilsirousi/golang-web-api/src/api/helper"
)

type User struct {
	Username string `binding:"required" json:"username"`
	Password string `binding:"required,password" json:"password"`
}

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func AllUsersHandler(c *gin.Context) {
	c.Header("content-type", "application/json")
	c.JSON(http.StatusOK, helper.NewBaseResponse("all users here!", true, 0))
}

func CreateUserHandler(c *gin.Context) {
	c.JSON(http.StatusCreated, helper.NewBaseResponse("user created", true, 0))
}

func GetUserByUsernameHandler(c *gin.Context) {
	username, _ := c.Params.Get("username")
	c.JSON(http.StatusOK, helper.NewBaseResponse(fmt.Sprintf("%s is here!", username), true, 0))
}

func GetUserByIdHandler(c *gin.Context) {
	id, _ := c.Params.Get("id")
	c.JSON(http.StatusOK, helper.NewBaseResponse(fmt.Sprintf("user: %s is here!", id), true, 0))
}

func GetUserByQuery(c *gin.Context) {
	id, _ := c.GetQuery("id")
	c.JSON(http.StatusOK, helper.NewBaseResponse(fmt.Sprintf("%s from query is here!", id), true, 0))
}

func GetRequestBody(c *gin.Context) {
	var body User
	err := c.ShouldBindBodyWithJSON(&body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.NewBaseResponseWithValidationError("faild to bind request body", false, -1, err))
		return
	}
	c.JSON(http.StatusOK, helper.NewBaseResponse(body, true, 0))
}
