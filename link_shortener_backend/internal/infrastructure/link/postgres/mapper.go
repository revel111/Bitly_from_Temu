package postgres

import domain "linkShortener/internal/domain/link"

func ToDb(link *domain.Link) *DbLink {
	if link == nil {
		return nil
	}

	return &DbLink{
		Id:        link.Id,
		ShortUrl:  link.ShortUrl,
		LongUrl:   link.LongUrl,
		CreatedAt: link.CreatedAt,
	}
}

func ToDomain(dbLink *DbLink) *domain.Link {
	if dbLink == nil {
		return nil
	}

	return &domain.Link{
		Id:        dbLink.Id,
		ShortUrl:  dbLink.ShortUrl,
		LongUrl:   dbLink.LongUrl,
		CreatedAt: dbLink.CreatedAt,
	}
}
