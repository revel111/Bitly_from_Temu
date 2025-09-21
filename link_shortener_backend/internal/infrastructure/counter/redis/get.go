package redis

import (
	"context"
	"strconv"
)

func (repo CounterRepository) Get(key string) (int64, error) {
	res, err := repo.CounterClient.Get(context.Background(), key).Result()

	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(res, 10, 64)
}
