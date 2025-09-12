package model

import "time"

type Link struct {
	Id        uint64    `gorm:"primaryKey" json:"id;autoIncrement"`
	ShortUrl  string    `gorm:"uniqueIndex" json:"short_url"`
	LongUrl   string    `gorm:"uniqueIndex" json:"long_url;not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at;not null"`
}
