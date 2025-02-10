package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock for testing purposes
type mockWeatherClient struct {
}

func (m *mockWeatherClient) Do(req *http.Request) (*http.Response, error) {
	return getDoFunc(req)
}

var (
	// GetDoFunc fetches the mock client's `Do` func
	getDoFunc                  func(req *http.Request) (*http.Response, error)
	weatherResponseBodySuccess string = `{
    "latitude": 52.52,
    "longitude": 13.419998,
    "generationtime_ms": 0.022292137145996094,
    "utc_offset_seconds": 3600,
    "timezone": "Europe/Berlin",
    "timezone_abbreviation": "GMT+1",
    "elevation": 38.0,
    "current_units": {
        "time": "iso8601",
        "interval": "seconds",
        "temperature_2m": "°C"
    },
    "current": {
        "time": "2025-01-22T20:00",
        "interval": 900,
        "temperature_2m": 1.2
    }
}`
)

// We set the mock to use in the tests
func init() {
	WeatherClient = &mockWeatherClient{}
}

func TestFetchWeatherSuccess(t *testing.T) {

	responseCh := make(chan WeatherResponseDTO)
	errCh := make(chan string)

	getDoFunc = func(req *http.Request) (*http.Response, error) {
		res := http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader([]byte(weatherResponseBodySuccess)))}

		return &res, nil
	}

	go ExecWeatherRequest("-34.6183919", "-58.442937", "auto", responseCh, errCh)

	select {
	case response := <-responseCh:
		assert.Equal(t, float32(1.2), response.CurrentWeather.Temperature)
	case err := <-errCh:
		t.Error(err)
	}

}

func TestFetchWithStatus500Fails(t *testing.T) {

	responseCh := make(chan WeatherResponseDTO)
	errCh := make(chan string)

	getDoFunc = func(req *http.Request) (*http.Response, error) {
		res := http.Response{StatusCode: 500, Body: io.NopCloser(bytes.NewReader([]byte("")))}

		return &res, nil
	}

	go ExecWeatherRequest("-34.6183919", "-58.442937", "auto", responseCh, errCh)

	err := <-errCh
	assert.Equal(t, fmt.Sprintf(MESSAGE_UNSUCCESSFULL_WEATHER_RESPONSE, 500), err)
}
