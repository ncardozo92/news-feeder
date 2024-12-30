package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/ncardozo92/news-feeder/api"
)

const flagDev = "--dev"

func main() {
	e := echo.New()

	if flagEnvIsPresent(os.Args) {
		if err := godotenv.Load("./.env"); err != nil {
			e.Logger.Fatal("Could not load .env file")
		}

		e.Logger.Info("Environment set to develop")
	}

	e.GET(api.PATH_FEED, api.GetFeed)

}

func flagEnvIsPresent(args []string) bool {
	var result bool

	for _, arg := range args {
		if arg == flagDev {
			result = true
			break
		}
	}

	return result
}
