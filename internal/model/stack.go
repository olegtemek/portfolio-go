package model

import "html/template"

type Stack struct {
	Id    int           `json:"id"`
	Order int           `json:"order"`
	Svg   template.HTML `json:"svg"`
	Title string        `json:"title"`
}
