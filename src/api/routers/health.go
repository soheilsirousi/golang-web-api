package router

import (
	"github.com/gin-gonic/gin"
	handler "github.com/soheilsirousi/golang-web-api/src/api/handlers"
)

func HealthRouter(r *gin.RouterGroup) {
	healthHandler := handler.NewHealthHandler()
	r.GET("/health", healthHandler.Health)
}
