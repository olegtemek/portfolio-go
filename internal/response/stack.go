package response

import "github.com/olegtemek/portfolio-go/internal/model"

type StackCreateResponse struct {
	*model.Stack `json:"stack"`
	Status       int `json:"status"`
}

type StackGetAllResponse struct {
	Stacks []*model.Stack `json:"stacks"`
	Status int            `json:"status"`
}

type StackDeleteResponse struct {
	*model.Stack `json:"stack"`
	Status       int `json:"status"`
}
