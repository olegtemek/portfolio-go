package response

import "github.com/olegtemek/portfolio-go/internal/model"

type InfoCreateOrUpdateResponse struct {
	*model.Info `json:"info"`
	Status      int `json:"status"`
}

type InfoGetResponse struct {
	*model.Info `json:"info"`
	Status      int `json:"status"`
}
