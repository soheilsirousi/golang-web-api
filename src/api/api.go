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
)

func InitServer(cnf *config.Config) {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middlewares.Limiter())
	SetValidation()
	SetRouter(r)

	r.Run(fmt.Sprintf(":%s", cnf.Server.Port))
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
