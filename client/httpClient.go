package client

import "net/http"

const MAX_STATUS_CODE_SUCCESS = 399

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
