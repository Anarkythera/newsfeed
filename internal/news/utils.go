package news

import (
	"database/sql"
	"newsletter/internal/news/model"

	"github.com/mmcdole/gofeed"
)

func checkForThumbnail(news *model.News, item *gofeed.Item) {
	thumbnail, ok := item.Extensions["media"]["thumbnail"]

	if ok {
		news.Thumbnail = sql.NullString{
			String: thumbnail[0].Attrs["url"],
			Valid:  true,
		}
	} else {
		news.Thumbnail = sql.NullString{
			String: "No Thumbnail",
			Valid:  true,
		}
	}
}
