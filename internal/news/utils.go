package news

import (
	"database/sql"
	"ziglunewsletter/internal/news/model"

	"github.com/mmcdole/gofeed"
)

func createThumbnail() error {

	return nil
}

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

	/*if thumbnail != nil {
		news.Thumbnail = sql.NullString{
			String: "thumbnail",
			Valid:  true,
		}
	} else {
		news.Thumbnail = sql.NullString{
			String: "No Thumbnail",
			Valid:  true,
		}
	}
	*/
}
