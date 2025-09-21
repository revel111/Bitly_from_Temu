package redis

import (
	"linkShortener/internal/domain/counter"

	"github.com/redis/go-redis/v9"
)

type CounterRepository struct {
	CounterClient *redis.Client
}

func NewCounterRepository(CounterClient *redis.Client) counter.Repository {
	return &CounterRepository{CounterClient: CounterClient}
}
