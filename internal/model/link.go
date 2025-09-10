package model

import "time"

type Link struct {
	ShortUrl  string    `gorm:"primaryKey" json:"short_url,omitempty"`
	LongUrl   string    `gorm:"uniqueIndex" json:"long_url,omitempty"`
	CreatedAt time.Time `json:"created_at" json:"created_at,omitempty"`
}
