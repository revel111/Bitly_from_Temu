package postgres

import (
	"linkShortener/internal/domain/link"
)

func (repo LinkRepository) GetByLongUrl(longUrl string) (*link.Link, error) {
	var linkModel *DbLink
	if err := repo.db.First(&linkModel, "long_url = ?", longUrl).Error; err != nil {
		return nil, err
	}

	return ToDomain(linkModel), nil
}
