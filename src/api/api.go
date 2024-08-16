package api

import (
	"github.com/gin-gonic/gin"
	router "github.com/soheilsirousi/golang-web-api/src/api/routers"
)

func InitServer() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1/")
	{
		router.Router(v1)
	}

	r.Run(":5050")
}
