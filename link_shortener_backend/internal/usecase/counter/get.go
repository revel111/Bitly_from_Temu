package counter

import (
	"errors"
	"linkShortener/internal/domain/counter"

	"github.com/redis/go-redis/v9"
)

type GetCounterUseCase interface {
	Execute(shortLink string) (int64, error)
}

type getCounterUseCase struct {
	counterRepo counter.Repository
}

func NewGetCounterUseCase(counterRepo counter.Repository) GetCounterUseCase {
	return &getCounterUseCase{counterRepo: counterRepo}
}

func (c getCounterUseCase) Execute(shortLink string) (int64, error) {
	counterValue, getErr := c.counterRepo.Get(shortLink)

	if errors.Is(getErr, redis.Nil) {
		return 0, counter.NewNotFoundError()
	}

	return counterValue, nil
}
