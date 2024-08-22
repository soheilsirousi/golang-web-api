package router

import (
	"github.com/gin-gonic/gin"
	handler "github.com/soheilsirousi/golang-web-api/src/api/handlers"
)

func UserRouter(r *gin.RouterGroup) {
	r.GET("/all", handler.AllUsersHandler)
	r.GET("/username/:username", handler.GetUserByUsernameHandler)
	r.GET("/id/:id", handler.GetUserByIdHandler)
	r.GET("/id", handler.GetUserByQuery)
	r.POST("/body", handler.GetRequestBody)
	r.POST("/create", handler.CreateUserHandler)
}
