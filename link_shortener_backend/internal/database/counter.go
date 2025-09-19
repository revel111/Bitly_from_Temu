package database

import (
	"context"
	"fmt"
	"linkShortener/configs"
	"log"

	"github.com/redis/go-redis/v9"
)

func ConnectCounter(data configs.ConfigData) (*redis.Client, func(), error) {
	CounterClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", data.RedisHost, data.RedisPort),
		Password: data.RedisPass,
		DB:       0,
	})

	response, err := CounterClient.Ping(context.Background()).Result()

	if err != nil || response != "PONG" {
		return nil, nil, err
	}

	closer := func() {
		if err := CounterClient.Close(); err != nil {
			log.Printf("Error closing Redis client: %v", err)
		}
	}

	return CounterClient, closer, nil
}
