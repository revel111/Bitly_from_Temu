package link

import (
	"errors"
	"linkShortener/internal/domain/counter"
	"linkShortener/internal/domain/link"
	"log"

	"github.com/redis/go-redis/v9"
)

type GetLinkUseCase interface {
	Execute(shortLink string) (string, error)
}

type getLinkUseCase struct {
	linkRepo    link.Repository
	counterRepo counter.Repository
}

func NewGetLinkUseCase(repo link.Repository, counterRepo counter.Repository) GetLinkUseCase {
	return &getLinkUseCase{linkRepo: repo, counterRepo: counterRepo}
}

func (c getLinkUseCase) Execute(shortLink string) (string, error) {
	linkModel, err := c.linkRepo.GetByShortUrl(shortLink)

	if err != nil {
		return "", link.NewNotFoundError()
	}

	_, incErr := c.counterRepo.Increment(shortLink)
	if incErr != nil && !errors.Is(incErr, redis.Nil) {
		log.Println("Error incrementing counter: ", incErr)
	}

	return linkModel.LongUrl, nil
}
