package main

import (
	"os"
	"time"
	"ziglunewsletter/internal/configuration"
	"ziglunewsletter/internal/database"
	"ziglunewsletter/internal/handlers"
	"ziglunewsletter/internal/news"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// CustomValidator API custom validator for all inputs
type CustomValidator struct {
	validator *validator.Validate
}

// Validate validates all API calls
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// First version of newsapp
func main() {
	cfg := configuration.ReadConf()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	log.Logger = log.With().Caller().Stack().Logger()

	dsn := database.BuildDSN(cfg.GetString("database.user"), cfg.GetString("database.pass"), cfg.GetString("database.host"), cfg.GetInt("database.port"), cfg.GetString("database.dbName"))
	dbx := database.ConnectDB(dsn)

	newsDao := news.NewDao(dbx)
	newsService := news.NewService(newsDao)

	handler := handlers.NewHandler(newsService)

	e := echo.New()
	e.StdLogger.SetOutput(log.Logger)
	e.StdLogger.SetPrefix("ECHO: ")
	e.Logger.SetOutput(log.Logger)
	e.Logger.SetPrefix("ECHO: ")

	log.Info().Msg("Starting news app server")

	e.Validator = &CustomValidator{validator: validator.New()}

	//newsService.GetNewsFromAllSources(0, 10)
	//newsService.NewsSourceFilter(0, 5, "TBD")

	e.GET("/GetNews", handler.GetNews)
	e.GET("/GetFilteredNews", handler.GetFilteredNews)

	// Get news and cache them
	if err := e.Start(cfg.GetString("httpServer.port")); err != nil {
		log.Fatal().Err(err).Msg("Error starting http server")
	}
}
