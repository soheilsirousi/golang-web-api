package main

import (
	"log"

	"github.com/soheilsirousi/golang-web-api/src/api"
	config "github.com/soheilsirousi/golang-web-api/src/configs"
	"github.com/soheilsirousi/golang-web-api/src/data/cache"
	"github.com/soheilsirousi/golang-web-api/src/data/db"
)

func main() {
	cnf := config.GetConfig()
	err := cache.InitRedis(cnf)
	if err != nil {
		log.Fatal(err)
	}
	defer cache.CloseRedis()
	err = db.InitPostgres(cnf)
	if err != nil {
		log.Fatal(err)
	}
	defer db.ClosePostgres()
	api.InitServer(cnf)
}
