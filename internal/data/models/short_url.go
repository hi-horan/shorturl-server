package models

import (
	"time"
)

type ShortUrl struct {
	ShortKey    string     `gorm:"column:short_key;primaryKey"`
	OriginalUrl string     `gorm:"column:original_url"`
	ExpireTime  *time.Time `gorm:"column:expire_time"`

	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

func (*ShortUrl) TableName() string {
	return "short_urls"
}
