package db

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis() {
	ctx := context.Background()
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis-server:6379",
		Password: "",
		DB:       0,
	})

	err := RedisClient.Ping(ctx).Err()
	if err != nil {
		fmt.Println("Error connecting to Redis : ", err.Error())
		return
	}

	fmt.Println("Connected to Redis...")
}

func Close() {
	if RedisClient != nil {
		RedisClient.Close()
	}
}
