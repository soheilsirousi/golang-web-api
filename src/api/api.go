package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	router "github.com/soheilsirousi/golang-web-api/src/api/routers"
	config "github.com/soheilsirousi/golang-web-api/src/configs"
)

func InitServer() {
	cnf := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		user := v1.Group("/user")
		router.UserRouter(user)
		router.HealthRouter(v1)
	}

	r.Run(fmt.Sprintf(":%s", cnf.Server.Port))
}
