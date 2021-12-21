package news

import (
	"fmt"
	"sort"

	"github.com/mmcdole/gofeed"
)

type Service interface {
	GetNewsFromSources(page, pagesize int)
}

type service struct {
	dao Dao
}

func NewService(dao Dao) Service {
	return &service{
		dao: dao,
	}
}

func (s *service) GetNewsFromSources(page, pagesize int) {
	sources, err := s.dao.GetNewsFromAllSources()
	if err != nil {
		fmt.Println(err)
	}

	feedItemsList := []*gofeed.Item{}
	fp := gofeed.NewParser()

	for _, source := range sources {
		feed, _ := fp.ParseURL(source.Url)
		feedItemsList = append(feedItemsList, feed.Items...)
	}

	for _, item := range feedItemsList {
		fmt.Println(item.Title)
		fmt.Println(item.Published)
	}

	//Descent sort of news based on date
	sort.Slice(feedItemsList, func(i, j int) bool { return feedItemsList[i].Published > feedItemsList[j].Published })

	fmt.Println("After sort")

	for _, item := range feedItemsList {
		fmt.Println(item.Title)
		fmt.Println(item.Published)
	}
}
