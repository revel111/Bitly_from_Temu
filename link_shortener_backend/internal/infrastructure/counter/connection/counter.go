package connection

import (
	"context"
	"fmt"
	"linkShortener/configs"
	"log"

	"github.com/redis/go-redis/v9"
)

func ConnectCounter(data configs.ConfigData) (*redis.Client, func(), error) {
	counterClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", data.RedisHost, data.RedisPort),
		Password: data.RedisPass,
		DB:       0,
	})

	response, err := counterClient.Ping(context.Background()).Result()

	if err != nil || response != "PONG" {
		return nil, nil, err
	}

	return counterClient, closeDb(counterClient), nil
}

func closeDb(counterClient *redis.Client) func() {
	return func() {
		if err := counterClient.Close(); err != nil {
			log.Printf("Error closing Redis client: %v", err)
		}
	}
}
