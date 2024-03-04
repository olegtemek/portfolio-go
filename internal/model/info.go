package model

type Info struct {
	Id       int    `json:"id,omitempty"`
	Email    string `json:"email"`
	Github   string `json:"github"`
	Telegram string `json:"telegram"`
	Location string `json:"location"`
	Position string `json:"position"`
	About    string `json:"about"`
}
