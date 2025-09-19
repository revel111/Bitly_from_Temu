package service

import (
	"context"
	"errors"
	"fmt"
	"linkShortener/pkg"
	"log"
	"strconv"

	"linkShortener/configs"
	"linkShortener/internal/database"
	"linkShortener/internal/httperrors"
	"linkShortener/internal/model"

	"github.com/redis/go-redis/v9"
)

// todo: service layer why http errors?
func CreateLink(longLink string) (shortedLink string, err *httperrors.HttpError) {
	var linkModel model.Link
	//todo: databse infra layer
	result := database.DB.First(&linkModel, "long_url = ?", longLink)

	if result.RowsAffected == 1 {
		return
	}

	linkModel.LongUrl = longLink
	if err := database.DB.Create(&linkModel).Error; err != nil {
		return "", httperrors.NewHttpError(503, "Service unavailable")
	}

	linkModel.ShortUrl = pkg.EncodeBase62(linkModel.Id)
	if err := database.DB.Save(&linkModel).Error; err != nil {
		return "", httperrors.NewHttpError(503, "Service unavailable")
	}

	return fmt.Sprintf("%s/api/v1/links/%s", configs.BaseUrl.GetValue(), linkModel.ShortUrl), nil
}

// todo: separate usecases
func GetLink(shortLink string) (longLink string, err *httperrors.HttpError) {
	var linkModel model.Link
	result := database.DB.First(&linkModel, "short_url = ?", shortLink)

	if result.RowsAffected == 0 {
		return "", httperrors.NewHttpError(404, "Link not found")
	}

	//todo: infra layer
	_, incErr := database.CounterClient.Incr(context.Background(), shortLink).Result()
	if incErr != nil && !errors.Is(incErr, redis.Nil) {
		log.Println("Error incrementing counter: ", incErr)
	}

	return linkModel.LongUrl, nil
}

func GetCounter(shortLink string) (uint64, *httperrors.HttpError) {
	counterValue, getErr := database.CounterClient.Get(context.Background(), shortLink).Result()
	if errors.Is(getErr, redis.Nil) {
		return 0, httperrors.NewHttpError(404, "Counter not found")
	}

	counter, parseErr := strconv.ParseUint(counterValue, 10, 64)
	if parseErr != nil {
		return 0, httperrors.NewHttpError(400, "Invalid counter value")
	}

	return counter, nil
}
