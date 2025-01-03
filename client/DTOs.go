package client

type WeatherResponseDTO struct {
	Hourly HourlyDTO `json:"hourly"`
}

type HourlyDTO struct {
	Times        []string  `json:"time"`
	Temperatures []float32 `json:"temperature_2m"`
}

type NewsResponseDTO struct {
	Articles []Article `json:"articles"`
}

type Article struct {
	Source      Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	ImageUrl    string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
	/*
		{
			"source": {
				"id": "wired",
				"name": "Wired"
			},
			"author": "Arunima Kar",
			"title": "Can Artificial Rain, Drones, or Satellites Clean Toxic Air?",
			"description": "India’s capital has turned to tech to fight its worst air pollution in eight years.",
			"url": "https://www.wired.com/story/artificial-rain-drones-and-satellites-can-tech-clean-indias-toxic-air/",
			"urlToImage": "https://media.wired.com/photos/6734cff01daede74a78d6818/191:100/w_1280,c_limit/AP24306187920631.jpg",
			"publishedAt": "2024-12-02T11:30:00Z",
			"content": "Amid all of these concerns, the city has been turning to drones to monitor pollution hotspots, in addition to those spraying water to suppress PM2.5. Drones are useful for accessing areas that are ha… [+3243 chars]"
		}
	*/
}

type Source struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
