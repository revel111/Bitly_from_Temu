package postgres

import "linkShortener/internal/domain/link"

func (repo LinkRepository) GetByShortUrl(shortUrl string) (*link.Link, error) {
	var linkModel DbLink
	if err := repo.db.First(&linkModel, "short_url = ?", shortUrl).Error; err != nil {
		return nil, err
	}

	return ToDomain(&linkModel), nil
}
