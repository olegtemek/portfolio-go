package response

import "github.com/olegtemek/portfolio-go/internal/model"

type ExperienceCreateResponse struct {
	*model.Experience `json:"experience"`
	Status            int `json:"status"`
}

type ExperienceGetAllResponse struct {
	Experiences []*model.Experience `json:"experiences"`
	Status      int                 `json:"status"`
}

type ExperienceUpdateResponse struct {
	Experience *model.Experience `json:"experience"`
	Status     int               `json:"status"`
}

type ExperienceDeleteResponse struct {
	*model.Experience `json:"experience"`
	Status            int `json:"status"`
}

type ExperienceGetOneResponse struct {
	*model.Experience `json:"experience"`
	Status            int `json:"status"`
}
