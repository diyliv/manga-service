package models

type Manga struct {
	Name        string `json:"manga_name"`
	Link        string `json:"manga_link"`
	Genre       string `json:"manga_genre"`
	Chapters    string `json:"manga_chapters"`
	Description string `json:"manga_description"`
	Year        string `json:"manga_year_start"`
	Rate        string `json:"manga_rationg"`
}
