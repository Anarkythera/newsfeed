package news

import (
	"sync"
	"ziglunewsletter/internal/news/model"
)

type newsCache struct {
	mu   sync.RWMutex
	news []*model.News
}

func newNewsCache() (*newsCache, error) {

	return nil, nil
}
