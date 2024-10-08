package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	config "github.com/soheilsirousi/golang-web-api/src/configs"
)

var redisClient *redis.Client

func InitRedis(cnf *config.Config) error {
	ctx := context.Background()
	redisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cnf.Redis.Host, cnf.Redis.Port),
		DB:           cnf.Redis.Db,
		Password:     cnf.Redis.Password,
		DialTimeout:  cnf.Redis.DialTimeout * time.Second,
		ReadTimeout:  cnf.Redis.ReadTimeout * time.Second,
		WriteTimeout: cnf.Redis.WriteTimeout * time.Second,
		PoolSize:     cnf.Redis.PoolSize,
		PoolTimeout:  cnf.Redis.PoolTimeout * time.Second,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}

func CloseRedis() {
	redisClient.Close()
}

func GetRedis() *redis.Client {
	return redisClient
}
