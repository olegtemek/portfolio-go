package response

import "github.com/olegtemek/portfolio-go/internal/model"

type ProjectCreateResponse struct {
	*model.Project `json:"project"`
	Status         int `json:"status"`
}

type ProjectGetAllResponse struct {
	Projects []*model.Project `json:"projects"`
	Status   int              `json:"status"`
}

type ProjectUpdateResponse struct {
	Project *model.Project `json:"project"`
	Status  int            `json:"status"`
}

type ProjectDeleteResponse struct {
	*model.Project `json:"project"`
	Status         int `json:"status"`
}

type ProjectGetOneResponse struct {
	*model.Project `json:"project"`
	Status         int `json:"status"`
}
