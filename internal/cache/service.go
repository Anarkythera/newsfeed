package newscache

import (
	"sync"
	"ziglunewsletter/internal/news"
	"ziglunewsletter/internal/news/model"
)

type Service interface {
}

type service struct {
	mu   sync.RWMutex
	news []*model.Item
	dao  *news.Dao
}
