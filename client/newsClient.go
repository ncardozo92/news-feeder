package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/gommon/log"
)

const (
	NEWS_API_URL_ENV = "NEWS_API_URL"
	NEWS_API_KEY     = "NEWS_API_KEY"
)

var newsClient HttpClient

func ExecNewsRequest() (NewsResponseDTO, error) {

	responseDTO := NewsResponseDTO{}
	newsUrl := os.Getenv(NEWS_API_URL_ENV)
	newsApiKey := os.Getenv(NEWS_API_KEY)

	// for testing purposes
	if newsUrl == "" {
		newsUrl = "http://localhost:8089?api_jey=%s"
	}

	fullApiUrl := fmt.Sprintf(newsUrl, newsApiKey)

	request, requestErr := http.NewRequest(http.MethodGet, fullApiUrl, bytes.NewReader([]byte("")))

	if requestErr != nil {
		log.Errorf("Error")
		return responseDTO, requestErr
	}

	response, fetchingErr := newsClient.Do(request)

	if fetchingErr != nil {
		return responseDTO, fetchingErr
	}

	if response.StatusCode >= MAX_STATUS_CODE_SUCCESS {
		return responseDTO, fmt.Errorf("error fetching data from news API")
	}

	responseBody, readingResponseBodyErr := io.ReadAll(response.Body)

	if readingResponseBodyErr != nil {
		return NewsResponseDTO{}, readingResponseBodyErr
	}

	if unmarshalingErr := json.Unmarshal(responseBody, &responseDTO); unmarshalingErr != nil {
		return NewsResponseDTO{}, unmarshalingErr
	}

	removeDeletedNews(&responseDTO)

	return responseDTO, nil
}

func removeDeletedNews(dto *NewsResponseDTO) {
	filteredArticles := []Article{}

	for _, article := range dto.Articles {
		if article.Source.Id != "" {
			filteredArticles = append(filteredArticles, article)
		}
	}

	dto.Articles = filteredArticles
}
