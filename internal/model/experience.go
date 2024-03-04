package model

type Experience struct {
	Id          int    `json:"id"`
	Order       int    `json:"order"`
	Title       string `json:"title"`
	Position    string `json:"position"`
	Period      string `json:"period"`
	Description string `json:"description"`
}
