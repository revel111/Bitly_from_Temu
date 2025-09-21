package link

import "time"

type Link struct {
	Id        uint64
	ShortUrl  string
	LongUrl   string
	CreatedAt time.Time
}
