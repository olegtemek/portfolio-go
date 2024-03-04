package rest

import (
	"fmt"
	"html/template"
	"net/http"
	"path"

	"github.com/olegtemek/portfolio-go/internal/model"
	"github.com/olegtemek/portfolio-go/internal/response"
)

type IndexTemplate struct {
	Info       *model.Info
	Experience []*model.Experience
	Project    []*model.Project
	Stack      []template.HTML
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("web", "main.html"))
	if err != nil {
		response.New(&w, r, fmt.Errorf("something went wrong, try again later %s", err), 400)
		return
	}

	info, err := h.Services.Info.Get()
	if err != nil {
		response.New(&w, r, fmt.Errorf("something went wrong, try again later %s", err), 400)
		return
	}
	projects, err := h.Services.Project.GetAll()
	if err != nil {
		response.New(&w, r, fmt.Errorf("something went wrong, try again later %s", err), 400)
		return
	}
	experiences, err := h.Services.Experience.GetAll()
	if err != nil {
		response.New(&w, r, fmt.Errorf("something went wrong, try again later %s", err), 400)
		return
	}
	stacks, err := h.Services.Stack.GetAll()
	if err != nil {
		response.New(&w, r, fmt.Errorf("something went wrong, try again later %s", err), 400)
		return
	}

	var stacksValues []template.HTML

	for _, n := range stacks {
		htmlEncapsulate := template.HTML(n.Svg)
		stacksValues = append(stacksValues, htmlEncapsulate)
	}

	data := IndexTemplate{
		Info:       info,
		Project:    projects,
		Experience: experiences,
		Stack:      stacksValues,
	}

	if err := tmpl.Execute(w, data); err != nil {
		response.New(&w, r, fmt.Errorf("something went wrong, try again later %s", err), 400)
	}
}
