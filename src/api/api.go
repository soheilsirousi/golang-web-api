package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/soheilsirousi/golang-web-api/src/api/middlewares"
	router "github.com/soheilsirousi/golang-web-api/src/api/routers"
	"github.com/soheilsirousi/golang-web-api/src/api/validations"
	config "github.com/soheilsirousi/golang-web-api/src/configs"
	"github.com/soheilsirousi/golang-web-api/src/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middlewares.Limiter())
	SetValidation()
	SetRouter(r)
	SetSwagger(r, cfg)
	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func SetRouter(r *gin.Engine) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		user := v1.Group("/user")
		router.UserRouter(user)
		router.HealthRouter(v1)
	}
}

func SetValidation() {
	val, ok := binding.Validator.Engine().(*validator.Validate)

	if ok {
		err := val.RegisterValidation("password", validations.PasswordValidation, true)
		if err != nil {
			log.Printf("error while register tag")
		}
	}
}

func SetSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
