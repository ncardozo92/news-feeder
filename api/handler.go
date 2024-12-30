package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	PATH_FEED = "/feed"
)

func GetFeed(c echo.Context) error {
	return c.NoContent(http.StatusNotImplemented)
}
