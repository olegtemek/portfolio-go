package model

type Project struct {
	Id          int    `json:"id"`
	Order       int    `json:"order"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
}
