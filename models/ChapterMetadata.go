package models

type ChapterData struct {
	Title   string `json:"title"`
	UrlRead string `json:"urlRead"`
	Scan    string `json:"scan"`
	Date    string `json:"date"`
}

type ChapterMetadata struct {
	StatusCode int           `json:"statusCode"`
	Data       []ChapterData `json:"data"`
}
