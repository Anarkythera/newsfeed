package news

import (
	"database/sql"
	"fmt"
	"sort"
	"ziglunewsletter/internal/news/model"

	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
)

type Service interface {
	GetNewsFromAllSources(page, pageSize int)
	NewsSourceFilter(page, pageSize int, source string) ([]model.News, error)
}

type service struct {
	dao Dao
}

func NewService(dao Dao) Service {
	return &service{
		dao: dao,
	}
}

func (s *service) GetNewsFromAllSources(page, pageSize int) {
	sources, err := s.dao.GetNewsSources()
	if err != nil {
		fmt.Println(err)
	}

	feedItemsList := []*gofeed.Item{}
	//feedItemsList := []*model.News{}
	fp := gofeed.NewParser()

	for _, source := range sources {
		feed, _ := fp.ParseURL(source.Url)
		sort.Sort(feed)
		feedItemsList = append(feedItemsList, feed.Items...)
		for _, item := range feed.Items {
			//fmt.Println("source: ", feed.Title)
			fmt.Println("Title: ", item.Title)
			fmt.Println("Date", item.PublishedParsed)
		}
	}

	//Descent sort of news based on date
	sort.Slice(feedItemsList, func(i, j int) bool { return feedItemsList[i].Published > feedItemsList[j].Published })

	fmt.Println("After sort")

	for _, item := range feedItemsList {
		fmt.Println(item.Title)
		fmt.Println(item.PublishedParsed)
	}

	s.saveNews(feedItemsList)
}

func (s *service) NewsSourceFilter(page, pageSize int, source string) ([]model.News, error) {
	newsFilter, err := s.dao.FilterNewsBySource(page, pageSize, source)
	if err != nil {
		return nil, errors.Wrapf(err, "error filtering news for source %s", source)
	}

	for _, item := range newsFilter {
		fmt.Println(item.Title)
		fmt.Println(item.Url)
	}

	return newsFilter, err
}

func (s *service) saveNews(items []*gofeed.Item) error {
	fmt.Println("SAVING NEWS")

	mostRecent, err := s.dao.GetMostRecentNewsDate()

	if errors.Is(err, sql.ErrNoRows) {
		//LOG WARN HERE
		fmt.Println(err)
		for _, item := range items {
			tmp := model.News{
				Url:         item.GUID,
				Title:       item.Title,
				Source:      "TBD",
				PublishDate: *item.PublishedParsed,
			}
			s.dao.InsertNews(tmp)
		}

		return nil

	} else if err != nil {
		return errors.Wrapf(err, "Couldn't get latest news publish date")
	}

	fmt.Println("MOST RECENT: ", mostRecent)

	for i := len(items) - 1; i >= 0; i-- {
		fmt.Println("item date: ", items[i].PublishedParsed)
		if items[i].PublishedParsed.After(mostRecent) {
			fmt.Println("INSERTING NEWS")
			tmp := model.News{
				Url:         items[i].GUID,
				Title:       items[i].Title,
				Source:      "TBD",
				PublishDate: *items[i].PublishedParsed,
			}
			s.dao.InsertNews(tmp)
		}
	}

	return nil
}
