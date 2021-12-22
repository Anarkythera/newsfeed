package model

import (
	"time"
)

type Sources struct {
	ID         int64     `db:"id"`
	CreatedAt  time.Time `db:"created_at"`
	Url        string    `db:"url"`
	SourceName string    `db:"source_name"`
}

type News struct {
	ID          int64     `db:"id"`
	CreatedAt   time.Time `db:"created_at"`
	Url         string    `db:"url"`
	Title       string    `db:"title"`
	Source      string    `db:"rss_source"`
	PublishDate time.Time `db:"publish_date"`
}
