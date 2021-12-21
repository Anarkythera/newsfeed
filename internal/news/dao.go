package news

import (
	"ziglunewsletter/internal/news/model"

	"github.com/jmoiron/sqlx"
)

type Dao interface {
	GetNewsFromAllSources() ([]model.Sources, error)
}

type dao struct {
	db *sqlx.DB
}

func NewDao(db *sqlx.DB) Dao {
	return &dao{
		db: db,
	}
}

func (d *dao) GetNewsFromAllSources() ([]model.Sources, error) {
	sourceURL := []model.Sources{}

	query := "SELECT * FROM sources"
	err := d.db.Select(&sourceURL, query)

	return sourceURL, err
}
