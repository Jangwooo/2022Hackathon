package redis

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

func Connection() (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_host"),
		Password: os.Getenv("redis_pwd"),
		DB:       2,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		log.Panic(err.Error())
	}

	return client
}
