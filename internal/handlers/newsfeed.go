package handlers

import (
	"net/http"
	"ziglunewsletter/internal/news"
	"ziglunewsletter/internal/news/model"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type NewsHandler interface {
	GetNews(c echo.Context) error
	GetFilteredNews(c echo.Context) error
}

type newsHandler struct {
	news news.Service
}

func NewHandler(svc news.Service) NewsHandler {
	return &newsHandler{
		news: svc,
	}
}

func (h *newsHandler) GetNews(c echo.Context) error {
	params := new(model.NewsRequest)

	if err := (&echo.DefaultBinder{}).Bind(params, c); err != nil {
		log.Error().Err(err).Msg("Failed to bind request to project struct")
		return c.JSON(http.StatusInternalServerError, "Wrong types of project fields")
	}

	if err := c.Validate(params); err != nil {
		log.Error().Err(err).Msg("One or more required fields are not in the request")
		return c.JSON(http.StatusInternalServerError, "Wrong types of project fields")
	}

	news, err := h.news.GetNewsFromAllSources(*params.Page, *params.PageSize)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to get news. %v", err)

		return c.JSON(http.StatusInternalServerError, "Failed to get news")
	}

	return c.JSON(http.StatusOK, news)
}

func (h *newsHandler) GetFilteredNews(c echo.Context) error {
	params := new(model.NewsRequestSourceFilter)

	if err := (&echo.DefaultBinder{}).Bind(params, c); err != nil {
		log.Error().Err(err).Msg("Failed to bind request to project struct")

		return c.JSON(http.StatusInternalServerError, "Wrong types of project fields")
	}

	if err := c.Validate(params); err != nil {
		log.Error().Err(err).Msg("One or more required fields are not in the request")

		return c.JSON(http.StatusInternalServerError, "Wrong types of project fields")
	}

	news, err := h.news.NewsSourceFilter(*params.Page, *params.PageSize, params.NewsSource)
	if err != nil {
		log.Error().Err(err).Msg("Error while filtering news")

		return c.JSON(http.StatusInternalServerError, "Failed to get filtered news")
	}

	return c.JSON(http.StatusOK, news)
}
