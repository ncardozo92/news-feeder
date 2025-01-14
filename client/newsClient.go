package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	NEWS_API_URL_ENV                    = "NEWS_API_URL"
	NEWS_API_KEY                        = "NEWS_API_KEY"
	MESSAGE_UNSUCCESSFULL_NEWS_RESPONSE = "unsuccessfulL response from news API, HTTP status was %d"
)

var NewsClient HttpClient

func init() {
	NewsClient = &http.Client{}
}

func ExecNewsRequest(newsCh chan<- NewsResponseDTO, errCh chan<- string) {

	responseDTO := NewsResponseDTO{}
	newsUrl := os.Getenv(NEWS_API_URL_ENV)
	newsApiKey := os.Getenv(NEWS_API_KEY)

	// for testing purposes
	if newsUrl == "" {
		newsUrl = "http://localhost:8089?api_key=%s"
	}

	fullApiUrl := fmt.Sprintf(newsUrl, newsApiKey)

	request, requestErr := http.NewRequest(http.MethodGet, fullApiUrl, bytes.NewReader([]byte("")))

	if requestErr != nil {
		errCh <- "Error generating requets to get the latetest news"
		return
	}

	response, fetchingErr := NewsClient.Do(request)

	if fetchingErr != nil {
		errCh <- "error fetching data from news API"
		return
	}

	if response.StatusCode >= MAX_STATUS_CODE_SUCCESS {
		errCh <- fmt.Sprintf(MESSAGE_UNSUCCESSFULL_NEWS_RESPONSE, response.StatusCode)
		return
	}

	responseBody, readingResponseBodyErr := io.ReadAll(response.Body)

	if readingResponseBodyErr != nil {
		errCh <- readingResponseBodyErr.Error()
		return
	}

	if unmarshalingErr := json.Unmarshal(responseBody, &responseDTO); unmarshalingErr != nil {
		errCh <- unmarshalingErr.Error()
		return
	}

	removeDeletedNews(&responseDTO)

	newsCh <- responseDTO
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
