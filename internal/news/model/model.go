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
	URL         string         `db:"url" json:"url"`
	Title       string         `db:"title" json:"title"`
	Provider    string         `db:"provider" json:"provider"`
	Category    string         `db:"category" json:"category"`
	PublishDate time.Time      `db:"publish_date" json:"publishDate"`
	Thumbnail   sql.NullString `db:"thumbnail" json:"thumbnail"`
}

type RssItem struct {
	Item      *gofeed.Item
	RssSource string
}

type NewsRequest struct {
	Page     *int `query:"page" validate:"required"`
	PageSize *int `query:"pageSize" validate:"required"`
}

type NewsRequestSourceFilter struct {
	Page     *int     `query:"page" validate:"required"`
	PageSize *int     `query:"pageSize" validate:"required"`
	Category []string `query:"category" validate:"required"`
	Provider []string `query:"provider" validate:"required"`
}
