package model

import (
	"time"

	"github.com/mmcdole/gofeed"
)

type News struct {
	Item *gofeed.Item
}

type Sources struct {
	ID         int64     `db:"id"`
	CreatedAt  time.Time `db:"created_at"`
	Url        string    `db:"url"`
	SourceName string    `db:"source_name"`
}
