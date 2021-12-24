package handlers_test

import (
	"os"
	"testing"
	"ziglunewsletter/internal/configuration"
	"ziglunewsletter/internal/database"
	"ziglunewsletter/internal/handlers"
	"ziglunewsletter/internal/news"

	"github.com/labstack/echo/v4"
	"github.com/mmcdole/gofeed"
)

func TestGetNews(t *testing.T) {

	cfg := configuration.ReadConf()
	file, err := os.Open("../../testdata/testfeed.xml")
	if err != nil {
		t.Errorf("Error opening the testdata file %v", err)
	}

	dsn := database.BuildDSN(cfg.GetString("database.user"), cfg.GetString("database.pass"), cfg.GetString("database.host"), cfg.GetInt("database.port"), cfg.GetString("database.dbName"))
	dbx := database.ConnectDB(dsn)

	newsDao := news.NewDao(dbx)
	newsService := news.NewService(newsDao)
	handler := handlers.NewHandler(newsService)

	e := echo.New()

	fp := gofeed.NewParser()
	feed, _ := fp.Parse(file)

}

func TestGetFilteredNews(t *testing.T) {

}
