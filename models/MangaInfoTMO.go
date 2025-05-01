package models

type MangaInfoTMO struct {
	Title string `json:"title"`

	Image string `json:"image"`

	Type string `json:"type"`

	Score string `json:"score"`

	Demography string `json:"demography"`

	Synopsis string `json:"synopsis"`

	Status string `json:"status"`

	Genres []string `json:"genres"`

	Chapters []Chapter `json:"chapters"`
}

type Chapter struct {
	Title string
}
