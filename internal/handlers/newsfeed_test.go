package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"newsletter/internal/configuration"
	"newsletter/internal/database"
	"newsletter/internal/handlers"
	"newsletter/internal/news"
	"newsletter/internal/news/model"
	"os"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	cfg *viper.Viper
	c   *http.Client
}

// CustomValidator API custom validator for all inputs
type CustomValidator struct {
	validator *validator.Validate
}

// Validate validates all API calls
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func (s *TestSuite) SetupSuite() {
	cfg := configuration.ReadConf(os.Getenv("CONFIG_FILE_TEST_LOCATION"), os.Getenv("CONFIG_FILE_TEST_NAME"))
	s.cfg = cfg
	s.c = http.DefaultClient

	dsn := database.BuildDSN(cfg.GetString("database.user"), cfg.GetString("database.pass"), os.Getenv("DATABASE_HOST"), cfg.GetInt("database.port"), cfg.GetString("database.dbName"))
	dbx := database.ConnectDB(dsn)

	newsDao := news.NewDao(dbx)
	newsService := news.NewService(newsDao)
	handler := handlers.NewHandler(newsService)

	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/v1/GetNews", handler.GetNews)
	e.GET("/v1/GetFilteredNews", handler.GetFilteredNews)

	go e.Start(cfg.GetString("httpServer.port"))
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) TestGetNews() {
	resp := new([]model.News)
	testParams := []byte(`{"page":0,"pageSize":10}`)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s%s/v1/GetNews", s.cfg.GetString("httpServer.host"), s.cfg.GetString("httpServer.port")), bytes.NewBuffer(testParams))
	if err != nil {
		s.Suite.T().Errorf("Error %v", err)
	}

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	response, err := s.c.Do(req)
	if err != nil {
		s.T().Logf("Error response: %v", err)
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		s.T().Logf("Error decoding: %v", err)
	}

	s.Equal(10, len(*resp), "Expected 10 got %d", len(*resp))
	s.Equal(200, response.StatusCode, "Expected 200 got %d", response.StatusCode)
}

func (s *TestSuite) TestFilterNewsByCategoryAndProvider() {
	resp := []model.News{}
	testParams := []byte(`{"page":0,"pageSize":5,"category":["uk"],"provider":["skynews"]}`)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s%s/v1/GetFilteredNews", s.cfg.GetString("httpServer.host"), s.cfg.GetString("httpServer.port")), bytes.NewBuffer(testParams))
	if err != nil {
		s.Suite.T().Logf("Error %v", err)
	}

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	response, err := s.c.Do(req)
	if err != nil {
		s.T().Logf("Error response: %v", err)
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		s.T().Logf("Error decoding: %v", err)
	}

	s.Equal(3, len(resp), "Expected 3 got %d", len(resp))
	s.Equal(200, response.StatusCode, "Expected 200 got %d", response.StatusCode)

	for _, item := range resp {
		s.Equal("uk", item.Category, "Expected Technology got %s", item.Category)
	}
}
