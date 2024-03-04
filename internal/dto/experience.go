package dto

type ExperienceCreateDto struct {
	Order       int    `json:"order,omitempty"`
	Title       string `json:"title" validate:"required"`
	Position    string `json:"position" validate:"required"`
	Period      string `json:"period" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type ExperienceUpdateDto struct {
	Order       int    `json:"order"`
	Title       string `json:"title" validate:"required"`
	Position    string `json:"position" validate:"required"`
	Period      string `json:"period" validate:"required"`
	Description string `json:"description" validate:"required"`
}
