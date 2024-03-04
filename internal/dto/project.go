package dto

type ProjectCreateDto struct {
	Order       int    `json:"order,omitempty"`
	Title       string `json:"title" validate:"required"`
	Link        string `json:"link" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type ProjectUpdateDto struct {
	Order       int    `json:"order,omitempty"`
	Title       string `json:"title" validate:"required"`
	Link        string `json:"link" validate:"required"`
	Description string `json:"description" validate:"required"`
}
