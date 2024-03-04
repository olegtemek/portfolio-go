package dto

type StackCreateDto struct {
	Order int    `json:"order,omitempty"`
	Svg   string `json:"svg" validate:"required"`
	Title string `json:"title" validate:"required"`
}
