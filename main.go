package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("./.env"); err != nil {
		fmt.Println("Could not load .env file")
		return
	}

	fmt.Printf("Weather api key is %s", os.Getenv("WEATHER_API_KEY"))
}
