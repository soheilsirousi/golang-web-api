package main

import (
	"github.com/soheilsirousi/golang-web-api/src/api"
	config "github.com/soheilsirousi/golang-web-api/src/configs"
	data "github.com/soheilsirousi/golang-web-api/src/data/cache"
)

func main() {
	cnf := config.GetConfig()
	data.InitRedis(cnf)
	api.InitServer(cnf)
}
