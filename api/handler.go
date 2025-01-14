package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ncardozo92/news-feeder/client"
)

const (
	PATH_FEED             = "/feed"
	QUERY_PARAM_LATITUDE  = "latitude"
	QUERY_PARAM_LONGITUDE = "longitude"
	GTM_ZONE_AUTO         = "auto"
)

func GetFeed(c echo.Context) error {

	latitude := c.QueryParam(QUERY_PARAM_LATITUDE)
	longitude := c.QueryParam(QUERY_PARAM_LONGITUDE)

	responseDTO := FeedResponseDTO{}

	weatherCh := make(chan client.WeatherResponseDTO)
	newsCh := make(chan client.NewsResponseDTO)
	errCh := make(chan string)

	go client.ExecWeatherRequest(latitude, longitude, GTM_ZONE_AUTO, weatherCh, errCh)
	go client.ExecNewsRequest(newsCh, errCh)

	select {
	case weatherResponse := <-weatherCh:
		mapWeather(&responseDTO, weatherResponse)
	case err := <-errCh:
		fmt.Println(err)
		c.Logger().Warn(err)
	}

	select {
	case newsResponse := <-newsCh:
		mapNews(&responseDTO, newsResponse)
	case err := <-errCh:
		c.Logger().Warn(err)
	}

	return c.JSON(http.StatusOK, responseDTO)
}

func mapNews(responseDTO *FeedResponseDTO, newsResponse client.NewsResponseDTO) {
	newsList := []NewsDTO{}

	for _, article := range newsResponse.Articles {
		news := NewsDTO{
			Author:      article.Author,
			Title:       article.Title,
			Description: article.Description,
			Url:         article.Url,
			PublishedAt: article.PublishedAt,
			Content:     article.Content,
		}

		newsList = append(newsList, news)
	}

	responseDTO.NewsList = newsList
}

func mapWeather(responseDTO *FeedResponseDTO, weatherResponse client.WeatherResponseDTO) {
	if len(weatherResponse.Hourly.Temperatures) > 0 {
		responseDTO.Temperature = weatherResponse.Hourly.Temperatures[0]
	}
}
