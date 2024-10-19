package model

type AddSong struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type Song struct {
	Id          string `json:"id"`
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type QueryParam struct {
	Field string `json:"field,omitempty"`
	Sort  string `json:"sort,omitempty"`
	Page  int    `json:"page,omitempty"`
	Limit int    `json:"limit,omitempty"`
}
