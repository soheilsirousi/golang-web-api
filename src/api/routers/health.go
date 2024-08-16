package router

import (
	"github.com/gin-gonic/gin"
	handler "github.com/soheilsirousi/golang-web-api/src/api/handlers"
)

func Router(r *gin.RouterGroup) {
	handler := handler.NewHealthHandler()
	r.GET("health/", handler.Health)
}
