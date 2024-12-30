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
	//doFunc func(*http.Client) (*http.Response, error)
}

func (m *mockWeatherClient) Do(req *http.Request) (*http.Response, error) {
	// coming soon!
	return getDoFunc(req)
}

var (
	// GetDoFunc fetches the mock client's `Do` func
	getDoFunc           func(req *http.Request) (*http.Response, error)
	responseBodySuccess string = `{
    "latitude": 52.52,
    "longitude": 13.419998,
    "generationtime_ms": 0.030040740966796875,
    "utc_offset_seconds": 0,
    "timezone": "GMT",
    "timezone_abbreviation": "GMT",
    "elevation": 38.0,
    "hourly_units": {
        "time": "iso8601",
        "temperature_2m": "°C"
    },
    "hourly": {
        "time": [
            "2024-12-12T00:00",
            "2024-12-12T01:00",
            "2024-12-12T02:00"
        ],
        "temperature_2m": [
            1.7,
            1.7,
            1.6
        ]
    }
}`
)

// We set the mock to use in the tests
func init() {
	weatherClient = &mockWeatherClient{}
}

func TestFetchSuccess(t *testing.T) {

	getDoFunc = func(req *http.Request) (*http.Response, error) {
		res := http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader([]byte(responseBodySuccess)))}

		return &res, nil
	}

	response, err := ExecWeatherRequest("-34.6183919", "-58.442937", "auto")

	assert.NoError(t, err)
	assert.NotEmpty(t, response.Hourly.Temperatures)
	assert.NotEmpty(t, response.Hourly.Times)

}

func TestFetchWithStatus500Fails(t *testing.T) {
	getDoFunc = func(req *http.Request) (*http.Response, error) {
		res := http.Response{StatusCode: 500, Body: io.NopCloser(bytes.NewReader([]byte("")))}

		return &res, nil
	}

	_, err := ExecWeatherRequest("-34.6183919", "-58.442937", "auto")

	assert.Error(t, err)
	assert.Equal(t, fmt.Sprintf(MESSAGE_UNSUCCESSFULL_RESPONSE, 500), err.Error())
}
