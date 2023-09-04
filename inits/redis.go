package inits

import (
	"fmt"
	"gologin/config"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var RedisClient *redis.Client

func RedisInit() {
	addr := fmt.Sprintf("%s:%d", config.Env.Redis.Host, config.Env.Redis.Port)

	fmt.Println(123, addr)
	RedisClient = redis.NewClient(&redis.Options{
		Addr: addr,
		// Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		logrus.Error(err)
	}
}
