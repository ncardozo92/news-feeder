package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/labstack/gommon/log"
)

const (
	WEATHER_API_URL_ENV            = "WEATHER_API_URL"
	MESSAGE_UNSUCCESSFULL_RESPONSE = "unsuccessfulL response from weather API, HTTP status was %d"
)

var weatherClient HttpClient

func Init() {
	weatherClient = &http.Client{}
}

func ExecWeatherRequest(latitude, longitude, GTMZone string) (WeatherResponseDTO, error) {

	queryString := url.Values{
		"latitude":  []string{latitude},
		"longitude": []string{longitude},
		"timezone":  []string{GTMZone},
	}

	responseDTO := WeatherResponseDTO{}

	fullApiUrl := fmt.Sprintf("%s?%s", os.Getenv(WEATHER_API_URL_ENV), queryString.Encode())

	request, requestErr := http.NewRequest(http.MethodGet, fullApiUrl, bytes.NewReader([]byte("")))

	if requestErr != nil {
		log.Errorf("Error")
		return responseDTO, requestErr
	}

	response, fetchingErr := weatherClient.Do(request)

	if fetchingErr != nil {
		return responseDTO, fetchingErr
	} else if response.StatusCode > MAX_STATUS_CODE_SUCCESS {
		return responseDTO, fmt.Errorf(MESSAGE_UNSUCCESSFULL_RESPONSE, response.StatusCode)
	}

	defer response.Body.Close()

	readResponseBody, readingResponseBodyErr := io.ReadAll(response.Body)

	if readingResponseBodyErr != nil {
		return WeatherResponseDTO{}, readingResponseBodyErr
	}

	if unmarshallingErr := json.Unmarshal(readResponseBody, &responseDTO); unmarshallingErr != nil {
		return WeatherResponseDTO{}, unmarshallingErr
	}

	return responseDTO, nil
}
