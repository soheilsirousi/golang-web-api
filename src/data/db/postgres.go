package db

import (
	"fmt"

	config "github.com/soheilsirousi/golang-web-api/src/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgresClient *gorm.DB

func InitPostgres(cnf *config.Config) error {
	cnn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Tehran",
		cnf.Postgres.Host, cnf.Postgres.User, cnf.Postgres.Password, cnf.Postgres.DbName, cnf.Postgres.Port, cnf.Postgres.SslMode)

	postgresClient, err := gorm.Open(postgres.Open(cnn), &gorm.Config{})

	if err != nil {
		return err
	}

	db, _ := postgresClient.DB()
	err = db.Ping()
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(cnf.Postgres.MaxOpenConn)
	db.SetMaxIdleConns(cnf.Postgres.MaxIdleConn)
	db.SetConnMaxLifetime(cnf.Postgres.MaxConnLifeTime)
	db.SetConnMaxIdleTime(cnf.Postgres.MaxConnIdleTime)
	return nil
}

func ClosePostgres() {
	db, _ := postgresClient.DB()
	db.Close()
}

func GetPostgres() *gorm.DB {
	return postgresClient
}
