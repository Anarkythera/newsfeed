package news

import (
	"database/sql"
	"fmt"
	"ziglunewsletter/internal/news/model"

	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type Service interface {
	GetNewsFromAllSources(page, pageSize int) ([]model.News, error)
	NewsSourceFilter(page, pageSize int, source []string) ([]model.News, error)
}

type service struct {
	dao Dao
}

func NewService(dao Dao) Service {
	return &service{
		dao: dao,
	}
}

func (s *service) GetNewsFromAllSources(page, pageSize int) ([]model.News, error) {
	sources, err := s.dao.GetNewsSources()
	if err != nil {
		fmt.Println(err)
	}

	feedItemsList := []*model.News{}
	fp := gofeed.NewParser()

	for _, source := range sources {
		feed, _ := fp.ParseURL(source.URL)
		for _, item := range feed.Items {
			fmt.Println("Title: ", item.Title)
			fmt.Println("Date", item.PublishedParsed)
			tmp := &model.News{
				URL:         item.GUID,
				Title:       item.Title,
				Provider:    source.Provider,
				Category:    source.Category,
				PublishDate: *item.PublishedParsed,
			}
			checkForThumbnail(tmp, item)
			fmt.Println("NEWS thumb: ", tmp.Thumbnail)
			feedItemsList = append(feedItemsList, tmp)
		}
	}

	//maybe is not needed? Currently already sorting in db
	//Descent sort of news based on date
	/*sort.Slice(feedItemsList, func(i, j int) bool {
		return feedItemsList[i].Item.PublishedParsed.Before(*feedItemsList[j].Item.PublishedParsed)
	})*/

	if err := s.saveNews(feedItemsList); err != nil {
		log.Error().Msgf("Error while saving new news into db %v", err)
	}

	return s.dao.GetNewsFromAllSources(page, pageSize)
}

func (s *service) NewsSourceFilter(page, pageSize int, source []string) ([]model.News, error) {
	newsFilter, err := s.dao.FilterNewsBySource(page, pageSize, source)
	if err != nil {
		return nil, errors.Wrapf(err, "error filtering news for source %s", source)
	}

	for _, item := range newsFilter {
		fmt.Println(item.Title)
		fmt.Println(item.URL)
	}

	return newsFilter, err
}

func (s *service) saveNews(items []*model.News) error {
	fmt.Println("SAVING NEWS")

	mostRecent, err := s.dao.GetMostRecentNewsDate()

	if errors.Is(err, sql.ErrNoRows) {
		//LOG WARN HERE
		log.Warn().Msgf("No columns exist in table news %v", err)

		for _, item := range items {
			err := s.dao.InsertNews(*item)
			if err != nil {
				return err
			}
		}

		return nil
	} else if err != nil {
		return errors.Wrapf(err, "Couldn't get latest news publish date")
	}

	for i := len(items) - 1; i >= 0; i-- {
		fmt.Println("item date: ", items[i].PublishDate)

		if items[i].PublishDate.After(mostRecent) {
			log.Info().Msgf("Inserting new with title %s from feed %s into DB", items[i].Title, items[i].Provider)
			fmt.Println("INSERTING NEWS")
			err := s.dao.InsertNews(*items[i])

			if err != nil {
				return err
			}
		}
	}

	return nil
}
