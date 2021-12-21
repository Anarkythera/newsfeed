package main

import (
	"fmt"
	"os"
	"time"
	"ziglunewsletter/internal/configuration"
	"ziglunewsletter/internal/database"
	"ziglunewsletter/internal/handlers"
	"ziglunewsletter/internal/news"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/mmcdole/gofeed"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// First version of newsapp

func main() {
	cfg := configuration.ReadConf()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	log.Logger = log.With().Caller().Stack().Logger()

	dsn := database.BuildDSN(cfg.GetString("database.user"), cfg.GetString("database.pass"), cfg.GetString("database.host"), cfg.GetInt("database.port"), cfg.GetString("database.dbName"))
	dbx := database.ConnectDB(dsn)

	newsDao := news.NewDao(dbx)
	newsService := news.NewService(newsDao)

	handler := handlers.NewHandler(newsService)

	e := echo.New()
	//e.StdLogger.SetOutput(log.Logger)
	//e.StdLogger.SetPrefix("ECHO: ")
	//e.Logger.SetOutput(log.Logger)
	//e.Logger.SetPrefix("ECHO: ")

	log.Info().Msg("Starting news app server")

	newsService.GetNewsFromSources(0, 10)

	e.GET("/GetNews", handler.GetNews)

	// Get news and cache them
	if err := e.Start(cfg.GetString("httpServer.port")); err != nil {
		log.Fatal().Err(err).Msg("Error starting http server")
	}
}

func getNews() {
	feedItemsList := []*gofeed.Item{}
	client := resty.New()
	resp, _ := client.R().Get("http://feeds.skynews.com/feeds/rss/uk.xml")
	fmt.Println("Response Info:")
	fmt.Println("  Body       :\n", string(resp.Body()))
	fmt.Println()

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("http://feeds.skynews.com/feeds/rss/uk.xml")
	fmt.Println("Channel: ", feed.Title)

	feedItemsList = append(feedItemsList, feed.Items...)

	fp2 := gofeed.NewParser()
	feed2, _ := fp2.ParseURL("http://feeds.bbci.co.uk/news/uk/rss.xml")

	fmt.Println("Channel: ", feed2.Title)

	feedItemsList = append(feedItemsList, feed2.Items...)

	for _, item := range feedItemsList {
		fmt.Println(item.Title)
	}
}
