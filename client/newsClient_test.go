package client

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	newsResponseSuccess = `
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
)

type mockNewsClient struct {
}

func (m *mockNewsClient) Do(req *http.Request) (*http.Response, error) {
	return getDoFunc(req)
}

// We set the mock to use in the tests
func init() {
	newsClient = &mockNewsClient{}
}

func TestFetchNewsSuccess(t *testing.T) {

	getDoFunc = func(req *http.Request) (*http.Response, error) {
		response := http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader([]byte(newsResponseSuccess)))}

		return &response, nil
	}

	response, fetchingErr := ExecNewsRequest()

	assert.NoError(t, fetchingErr)
	assert.Equal(t, 1, len(response.Articles))
}

func TestFetchNewsFails(t *testing.T) {

	getDoFunc = func(req *http.Request) (*http.Response, error) {
		response := http.Response{StatusCode: 500}

		return &response, nil
	}

	_, fetchingErr := ExecNewsRequest()

	assert.Error(t, fetchingErr)
}
