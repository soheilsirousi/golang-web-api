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

// All User
// @Summary all user
// @Description show all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} helper.BaseResponse "Success"
// @Failure 400 {object} helper.BaseResponse "Failed"
// @Router /v1/user/all [get]
func AllUsersHandler(c *gin.Context) {
	c.Header("content-type", "application/json")
	c.JSON(http.StatusOK, helper.NewBaseResponse("all users here!", true, 0))
}

func CreateUserHandler(c *gin.Context) {
	c.JSON(http.StatusCreated, helper.NewBaseResponse("user created", true, 0))
}

// Get user
// @Summary get user
// @Description get user by username
// @Tags users
// @Accept json
// @Produce json
// @Param username path string true "user username"
// @Success 200 {object} helper.BaseResponse "Success"
// @Failure 400 {object} helper.BaseResponse "Failed"
// @Router /v1/user/username/{username} [get]
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

// Create user
// @Summary create user
// @Description create user
// @Tags users
// @Accept json
// @Produce json
// @Param user body User true "User data"
// @Success 200 {object} helper.BaseResponse "Success"
// @Failure 400 {object} helper.BaseResponse "Failed"
// @Router /v1/user/create [post]
func GetRequestBody(c *gin.Context) {
	var user User
	err := c.ShouldBindBodyWithJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.NewBaseResponseWithValidationError("failed to bind request body", false, -1, err))
		return
	}
	c.JSON(http.StatusOK, helper.NewBaseResponse(user, true, 0))
}
