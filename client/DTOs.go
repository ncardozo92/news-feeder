package client

type WeatherResponseDTO struct {
	Hourly HourlyDTO `json:"hourly"`
}

type HourlyDTO struct {
	Times        []string  `json:"time"`
	Temperatures []float32 `json:"temperature_2m"`
}
