package model

type Stack struct {
	Id    int    `json:"id"`
	Order int    `json:"order"`
	Svg   string `json:"svg"`
	Title string `json:"title"`
}
