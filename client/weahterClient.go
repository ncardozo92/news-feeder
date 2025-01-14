package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const (
	WEATHER_API_URL_ENV                    = "WEATHER_API_URL"
	MESSAGE_UNSUCCESSFULL_WEATHER_RESPONSE = "unsuccessfulL response from weather API, HTTP status was %d"
)

var WeatherClient HttpClient

func init() {
	WeatherClient = &http.Client{}
}

func ExecWeatherRequest(latitude, longitude, GTMZone string, weatherCh chan<- WeatherResponseDTO, errCh chan<- string) {

	queryString := url.Values{
		"latitude":  []string{latitude},
		"longitude": []string{longitude},
		"timezone":  []string{GTMZone},
	}

	responseDTO := WeatherResponseDTO{}

	fullApiUrl := fmt.Sprintf("%s?%s", os.Getenv(WEATHER_API_URL_ENV), queryString.Encode())

	request, requestErr := http.NewRequest(http.MethodGet, fullApiUrl, bytes.NewReader([]byte{}))

	if requestErr != nil {
		errCh <- requestErr.Error()
		return
	}

	response, fetchingErr := WeatherClient.Do(request)

	if fetchingErr != nil {
		errCh <- fetchingErr.Error()
		return
	} else if response.StatusCode > MAX_STATUS_CODE_SUCCESS {
		errCh <- fmt.Sprintf(MESSAGE_UNSUCCESSFULL_WEATHER_RESPONSE, response.StatusCode)
		return
	}

	defer response.Body.Close()

	readResponseBody, readingResponseBodyErr := io.ReadAll(response.Body)

	if readingResponseBodyErr != nil {
		errCh <- readingResponseBodyErr.Error()
		return
	}

	if unmarshallingErr := json.Unmarshal(readResponseBody, &responseDTO); unmarshallingErr != nil {
		errCh <- unmarshallingErr.Error()
		return
	}

	// Everything is OK
	weatherCh <- responseDTO
}
