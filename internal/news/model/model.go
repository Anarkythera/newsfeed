package model

import (
	"database/sql"
	"time"

	"github.com/mmcdole/gofeed"
)

type Sources struct {
	ID        int64     `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	URL       string    `db:"url"`
	Provider  string    `db:"provider"`
	Category  string    `db:"category"`
}

type News struct {
	ID          int64          `db:"id"`
	CreatedAt   time.Time      `db:"created_at"`
	URL         string         `db:"url"`
	Title       string         `db:"title"`
	Provider    string         `db:"provider"`
	Category    string         `db:"category"`
	PublishDate time.Time      `db:"publish_date"`
	Thumbnail   sql.NullString `db:"thumbnail"`
}

type RssItem struct {
	Item      *gofeed.Item
	RssSource string
}

type NewsRequest struct {
	Page     *int `json:"page" validate:"required"`
	PageSize *int `json:"pageSize" validate:"required"`
}

type NewsRequestSourceFilter struct {
	Page       *int     `json:"page" validate:"required"`
	PageSize   *int     `json:"pageSize" validate:"required"`
	NewsSource []string `json:"newsSource" validate:"required"`
}
