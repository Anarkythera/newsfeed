package news

import (
	"database/sql"
	"newsletter/internal/news/model"

	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type Service interface {
	GetNewsFromAllSources(page, pageSize int) ([]model.News, error)
	NewsSourceFilter(page, pageSize int, source []string, category []string) ([]model.News, error)
}

type service struct {
	dao Dao
}

func NewService(dao Dao) Service {
	svc := &service{
		dao: dao,
	}

	if err := svc.updateDatabase(); err != nil {
		log.Warn().Msgf("Couldn't update the database with new news %v", err)
	}

	return svc
}

func (s *service) GetNewsFromAllSources(page, pageSize int) ([]model.News, error) {
	return s.dao.GetNewsFromAllSources(page, pageSize)
}

func (s *service) NewsSourceFilter(page, pageSize int, source []string, category []string) ([]model.News, error) {
	newsFilter, err := s.dao.FilterNewsBySource(page, pageSize, source, category)
	if err != nil {
		return nil, errors.Wrapf(err, "error filtering news for source %s", source)
	}

	return newsFilter, err
}

func (s *service) saveNews(items []*model.News) error {
	mostRecent, err := s.dao.GetMostRecentNewsDate()

	if errors.Is(err, sql.ErrNoRows) {
		log.Warn().Msgf("No columns exist in table news %v", err)

		for _, item := range items {
			log.Info().Msgf("Inserting news %s from feed %s into the db", item.Title, item.Provider)
			err := s.dao.InsertNews(*item)

			if err != nil {
				log.Warn().Msgf("Error while inserting the news %s from feed %s into DB. %v", item.Title, item.Provider, err)
			}
		}

		return nil
	} else if err != nil {
		return errors.Wrapf(err, "Couldn't get latest news publish date")
	}

	for _, item := range items {
		if item.PublishDate.After(mostRecent) {
			log.Info().Msgf("Inserting news %s from feed %s into the db", item.Title, item.Provider)
			err := s.dao.InsertNews(*item)

			if err != nil {
				log.Warn().Msgf("Error while inserting the news %s from feed %s into DB. %v", item.Title, item.Provider, err)
			}
		}
	}

	return nil
}

func (s *service) updateDatabase() error {
	sources, err := s.dao.GetNewsSources()
	if err != nil {
		return err
	}

	feedItemsList := []*model.News{}
	fp := gofeed.NewParser()

	for _, source := range sources {
		feed, _ := fp.ParseURL(source.URL)
		for _, item := range feed.Items {
			tmp := &model.News{
				URL:         item.GUID,
				Title:       item.Title,
				Provider:    source.Provider,
				Category:    source.Category,
				PublishDate: *item.PublishedParsed,
			}
			checkForThumbnail(tmp, item)
			feedItemsList = append(feedItemsList, tmp)
		}
	}

	if err := s.saveNews(feedItemsList); err != nil {
		log.Error().Msgf("Error while saving new news into db %v", err)

		return err
	}

	return nil
}
