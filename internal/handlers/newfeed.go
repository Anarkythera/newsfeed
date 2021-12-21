package handlers

import (
	"net/http"
	"ziglunewsletter/internal/news"

	"github.com/labstack/echo/v4"
)

type NewsHandler interface {
	GetNews(c echo.Context) error
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
	return c.XML(http.StatusOK, "hello")
}
