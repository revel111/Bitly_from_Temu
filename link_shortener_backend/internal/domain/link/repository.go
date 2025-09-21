package link

type Repository interface {
	Save(link *Link) (*Link, error)
	GetByShortUrl(shortUrl string) (*Link, error)
	GetByLongUrl(longUrl string) (*Link, error)
}
