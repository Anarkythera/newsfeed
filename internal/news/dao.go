package news

import (
	"time"
	"ziglunewsletter/internal/news/model"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Dao interface {
	GetNewsSources() ([]model.Sources, error)
	InsertNews(model.News) error
	GetNewsFromAllSources(page, pageSize int) ([]model.News, error)
	FilterNewsBySource(page, pageSize int, newsSource []string) ([]model.News, error)
	GetMostRecentNewsDate() (time.Time, error)
}

type dao struct {
	db *sqlx.DB
}

func NewDao(db *sqlx.DB) Dao {
	return &dao{
		db: db,
	}
}

func (d *dao) GetNewsSources() ([]model.Sources, error) {
	sourceURL := []model.Sources{}

	query := "SELECT * FROM sources"
	err := d.db.Select(&sourceURL, query)

	return sourceURL, errors.Wrapf(err, "Error retrieving news sources from the database")
}

func (d *dao) GetNewsFromAllSources(page, pageSize int) ([]model.News, error) {
	news := []model.News{}

	query := "SELECT * FROM news ORDER BY publish_date LIMIT ? OFFSET ?"
	err := d.db.Select(&news, d.db.Rebind(query), pageSize, page)

	return news, errors.Wrapf(err, "Error fetching news from all sources from DB")
}

func (d *dao) InsertNews(news model.News) error {
	stmt, err := d.db.Preparex("INSERT INTO news (url, title, provider, category, publish_date, thumbnail) VALUES ($1,$2,$3,$4,$5,$6)")
	if err != nil {
		return errors.Wrapf(err, "Error preparing the sql statement")
	}

	defer stmt.Close()
	_, err = stmt.Exec(news.URL, news.Title, news.Provider, news.Category, news.PublishDate, news.Thumbnail)

	return errors.Wrapf(err, "Error inserting the record in the database")
}

func (d *dao) FilterNewsBySource(page, pageSize int, newsSource []string) ([]model.News, error) {
	news := []model.News{}

	query, queryArgs, _ := sqlx.In("SELECT * FROM news WHERE rss_source IN (?) ORDER BY publish_date DESC LIMIT ? OFFSET ? ", newsSource, pageSize, page)

	err := d.db.Select(&news, d.db.Rebind(query), queryArgs...)

	return news, errors.Wrapf(err, "Error while filtering the news from the DB")
}

func (d *dao) GetMostRecentNewsDate() (time.Time, error) {
	news := model.News{}

	query := "SELECT * FROM news ORDER BY publish_date DESC LIMIT 1"
	err := d.db.Get(&news, query)

	return news.PublishDate, errors.Wrapf(err, "Error fetching most recent news")
}
