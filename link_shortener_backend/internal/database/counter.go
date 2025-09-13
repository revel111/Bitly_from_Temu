package database

import (
	"context"
	"fmt"
	"linkShortener/configs"
	"log"

	"github.com/redis/go-redis/v9"
)

var CounterClient *redis.Client

func ConnectCounter() *redis.Client {
	CounterClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", configs.RedisHost.GetValue(), configs.RedisPort.GetValue()),
		Password: configs.RedisPass.GetValue(),
		DB:       0,
	})

	response, err := CounterClient.Ping(context.Background()).Result()
	if err != nil || response != "PONG" {
		log.Fatalf("Error connecting to redis counter: %v", err)
	}

	return CounterClient
}
