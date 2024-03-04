package dto

type InfoCreateOrUpdateDto struct {
	Email    string `json:"email" validate:"required"`
	Github   string `json:"github" validate:"required"`
	Telegram string `json:"telegram" validate:"required"`
	Location string `json:"location" validate:"required"`
	Position string `json:"position" validate:"required"`
	About    string `json:"about" validate:"required"`
}
