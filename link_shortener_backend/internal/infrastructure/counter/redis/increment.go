package redis

import "context"

func (repo CounterRepository) Increment(key string) (int64, error) {
	return repo.CounterClient.Incr(context.Background(), key).Result()
}
