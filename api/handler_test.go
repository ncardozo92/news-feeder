package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/ncardozo92/news-feeder/client"
	"github.com/stretchr/testify/assert"
)

type mockWeatherClient struct {
}

func (m *mockWeatherClient) Do(req *http.Request) (*http.Response, error) {
	return getWeatherDoFunc(req)
}

type mockNewsClient struct {
}

func (m *mockNewsClient) Do(req *http.Request) (*http.Response, error) {
	return getNewsDoFunc(req)
}

var getWeatherDoFunc func(req *http.Request) (*http.Response, error)
var getNewsDoFunc func(req *http.Request) (*http.Response, error)
var weatherResponseBodySuccess string = `{
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

var newsResponseSuccess = `
{
	"status": "ok",
	"totalResults": 39123,
	"articles": [
		{
			"source": {
				"id": null,
				"name": "[Removed]"
			},
			"author": null,
			"title": "[Removed]",
			"description": "[Removed]",
			"url": "https://removed.com",
			"urlToImage": null,
			"publishedAt": "2024-12-20T13:00:06Z",
			"content": "[Removed]"
		},
		{
			"source": {
				"id": "wired",
				"name": "Wired"
			},
			"author": "Paresh Dave",
			"title": "This Website Shows How Much Google’s AI Can Glean From Your Photos",
			"description": "A photo sharing startup founded by an ex-Google engineer found a clever way to turn Google’s tech against itself.",
			"url": "https://www.wired.com/story/website-google-ai-photos-ente/",
			"urlToImage": "https://media.wired.com/photos/6747781f0dabf1e9f09fed7e/191:100/w_1280,c_limit/AI-Photo-Information-Scan-Business-955510024.mp4",
			"publishedAt": "2024-12-02T11:30:00Z",
			"content": "Software engineer Vishnu Mohandas decided he would quit Google in more ways than one when he learned the tech giant had briefly helped the US military develop AI to study drone footage. In 2020, he l… [+3180 chars]"
		},
		{
			"source": {
				"id": null,
				"name": "[Removed]"
			},
			"author": null,
			"title": "[Removed]",
			"description": "[Removed]",
			"url": "https://removed.com",
			"urlToImage": null,
			"publishedAt": "2024-12-01T18:15:17Z",
			"content": "[Removed]"
		}
	]
}
`

func init() {
	client.WeatherClient = &mockWeatherClient{}
	client.NewsClient = &mockNewsClient{}
}

func TestGetFeedSuccess(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	getWeatherDoFunc = func(req *http.Request) (*http.Response, error) {
		res := http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader([]byte(weatherResponseBodySuccess)))}

		return &res, nil
	}

	getNewsDoFunc = func(req *http.Request) (*http.Response, error) {
		response := http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader([]byte(newsResponseSuccess)))}

		return &response, nil
	}

	handlerErr := GetFeed(c)

	fmt.Println(rec.Body.String())

	assert.NoError(t, handlerErr)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), "1.7"))                                                                // response contains temperature
	assert.True(t, strings.Contains(rec.Body.String(), "This Website Shows How Much Google’s AI Can Glean From Your Photos")) // response contains temperature
}
