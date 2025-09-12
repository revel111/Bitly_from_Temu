package service

import (
	"fmt"
	"linkShortener/internal/database"
	"linkShortener/internal/errors"
	"linkShortener/internal/model"
	"linkShortener/internal/utils"
)

func CreateLink(longLink string) (shortedLink string, err *errors.HttpError) {
	var linkModel model.Link
	result := database.DB.First(&linkModel, "long_url = ?", longLink)

	if result.RowsAffected == 1 {
		return linkModel.LongUrl, errors.NewHttpError(409, "Link already exists")
	}

	linkModel.LongUrl = longLink
	if err := database.DB.Create(&linkModel).Error; err != nil {
		return "", errors.NewHttpError(503, "Service unavailable")
	}

	linkModel.ShortUrl = utils.Encode(linkModel.Id)
	if err := database.DB.Save(&linkModel).Error; err != nil {
		return "", errors.NewHttpError(503, "Service unavailable")
	}

	return fmt.Sprintf("http://localhost:8080/api/v1/links/%s", linkModel.ShortUrl), nil
}

func GetLink(shortLink string) (longLink string, err *errors.HttpError) {
	var linkModel model.Link
	result := database.DB.First(&linkModel, "id = ?", shortLink)

	if result.RowsAffected == 0 {
		return "", errors.NewHttpError(404, "Link not found")
	}

	return linkModel.LongUrl, nil
}
