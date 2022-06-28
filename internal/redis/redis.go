package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func NewRedis() *redis.Client {

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "",
		DB:       0,
	})
	return redisClient
}
