package service

import (
	"linkShortener/internal/database"
	"linkShortener/internal/errors"
	"linkShortener/internal/model"
)

func CreateLink(longLink string) (shortedLink string, err *errors.HttpError) {
	var linkModel *model.Link
	result := database.DB.First(&linkModel, "long_url = ?", longLink)

	if result.RowsAffected == 1 {
		return "", errors.NewHttpError(400, "Link already exists")
	}

	var shortLink = Shorten(longLink)
	database.DB.Create(&model.Link{ShortUrl: shortLink, LongUrl: longLink})

	return "shortedLink", nil
}

func GetLink(shortLink string) (longLink string, err *errors.HttpError) {
	var linkModel model.Link
	result := database.DB.First(&linkModel, "short_url = ?", shortLink)

	if result.RowsAffected == 0 {
		return "", errors.NewHttpError(404, "Link not found")
	}

	return linkModel.LongUrl, nil
}
