package link

import (
	"errors"
	"linkShortener/internal/domain/link"
	"linkShortener/pkg"
	"net/url"

	"gorm.io/gorm"
)

type CreateLinkUseCase interface {
	Execute(longLink string) (string, error)
}

type createLinkUseCase struct {
	repo link.Repository
}

func NewCreateLinkUseCase(repo link.Repository) CreateLinkUseCase {
	return &createLinkUseCase{repo: repo}
}

func (c createLinkUseCase) Execute(longLink string) (string, error) {
	_, err := url.ParseRequestURI(longLink)
	if err != nil {
		return "", link.NewInvalidUrlError()
	}

	linkModel, err := c.repo.GetByLongUrl(longLink)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", link.NewServiceUnavailableError()
	} else if linkModel != nil {
		return "", link.NewExistsError()
	}

	linkModel = &link.Link{LongUrl: longLink}
	if linkModel, err = c.repo.Save(linkModel); err != nil {
		return "", link.NewServiceUnavailableError()
	}

	linkModel.ShortUrl = pkg.EncodeBase62(linkModel.Id)
	if linkModel, err = c.repo.Save(linkModel); err != nil {
		return "", link.NewServiceUnavailableError()
	}

	return linkModel.ShortUrl, nil
}
