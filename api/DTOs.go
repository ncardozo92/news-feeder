package api

type FeedResponseDTO struct {
	Temperature float32   `json:"temperature"`
	NewsList    []NewsDTO `json:"news"`
}

type NewsDTO struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"source_url"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}
