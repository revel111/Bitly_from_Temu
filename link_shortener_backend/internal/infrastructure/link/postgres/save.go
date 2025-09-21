package postgres

import "linkShortener/internal/domain/link"

func (repo LinkRepository) Save(link *link.Link) (*link.Link, error) {
	linkModel := ToDb(link)

	if err := repo.db.Save(&linkModel).Error; err != nil {
		return nil, err
	}

	return ToDomain(linkModel), nil
}
