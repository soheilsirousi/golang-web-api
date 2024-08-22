package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	router "github.com/soheilsirousi/golang-web-api/src/api/routers"
	"github.com/soheilsirousi/golang-web-api/src/api/validations"
	config "github.com/soheilsirousi/golang-web-api/src/configs"
)

func InitServer() {
	cnf := config.GetConfig()
	r := gin.New()
	val, ok := binding.Validator.Engine().(*validator.Validate)
	r.Use(gin.Logger(), gin.Recovery())

	if ok {
		err := val.RegisterValidation("password", validations.PasswordValidation, true)
		if err != nil {
			log.Printf("error while register tag")
		}
	}

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		user := v1.Group("/user")
		router.UserRouter(user)
		router.HealthRouter(v1)
	}

	r.Run(fmt.Sprintf(":%s", cnf.Server.Port))
}
