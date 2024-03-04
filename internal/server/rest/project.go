package rest

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/olegtemek/portfolio-go/internal/dto"
	"github.com/olegtemek/portfolio-go/internal/response"
)

// @Summary create
// @Security BasicAuth
// @Tags project
// @Accept json
// @Produce json
// @Param input body dto.ProjectCreateDto true "Project"
// @Success 200 {object} response.ProjectCreateResponse
// @Router /project [post]
func (h *Handler) CreateProject(w http.ResponseWriter, r *http.Request) {
	h.Log = h.Log.With("Source", "ProjectHandler:createProject")

	var req dto.ProjectCreateDto

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		h.Log.Error("ERROR", err)
		response.New(&w, r, fmt.Errorf("failed to decode request body: %s", err), 400)
		return
	}

	if err := validator.New().Struct(req); err != nil {
		validateErr := err.(validator.ValidationErrors)
		render.JSON(w, r, response.ValidationError(validateErr))
		return
	}

	project, err := h.Services.Project.Create(&req)

	if err != nil {
		h.Log.Error("ERROR", err)
		response.New(&w, r, fmt.Errorf("something went wrong: %s", err), 400)
		return
	}

	render.JSON(w, r, &response.ProjectCreateResponse{
		Project: project,
		Status:  200,
	})
}

// @Summary update
// @Security BasicAuth
// @Tags project
// @Accept json
// @Produce json
// @Param id path int true "Project"
// @Param input body dto.ProjectUpdateDto true "Project"
// @Success 200 {object} response.ProjectUpdateResponse
// @Router /project/{id} [patch]
func (h *Handler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	h.Log = h.Log.With("Source", "ProjectHandler:updateProject")

	var req dto.ProjectUpdateDto

	id := chi.URLParam(r, "id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.Log.Error("ERROR", err)
		response.New(&w, r, fmt.Errorf("cannot parse id: %s", err), 400)
		return
	}

	err = render.DecodeJSON(r.Body, &req)
	if err != nil {
		h.Log.Error("ERROR", err)
		response.New(&w, r, fmt.Errorf("failed to decode request body: %s", err), 400)
		return
	}

	if err := validator.New().Struct(req); err != nil {
		validateErr := err.(validator.ValidationErrors)
		render.JSON(w, r, response.ValidationError(validateErr))
		return
	}

	project, err := h.Services.Project.Update(idInt, &req)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			response.New(&w, r, fmt.Errorf("project not found: %s", err), 404)
			return
		}

		response.New(&w, r, fmt.Errorf("something went wrong: %s", err), 400)
		return
	}

	render.JSON(w, r, &response.ProjectUpdateResponse{
		Project: project,
		Status:  200,
	})
}

// @Summary delete
// @Security BasicAuth
// @Tags project
// @Accept json
// @Produce json
// @Param id path int true "Project id"
// @Success 200 {object} response.ProjectDeleteResponse
// @Router /project/{id} [delete]
func (h *Handler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	h.Log = h.Log.With("Source", "ProjectHandler:deleteProject")
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.Log.Error("ERROR", err)
		response.New(&w, r, fmt.Errorf("cannot parse id: %s", err), 400)
		return
	}

	project, err := h.Services.Project.Delete(idInt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			response.New(&w, r, fmt.Errorf("project not found: %s", err), 404)
			return
		}

		response.New(&w, r, fmt.Errorf("something went wrong: %s", err), 400)
		return
	}

	render.JSON(w, r, &response.ProjectDeleteResponse{
		Project: project,
		Status:  200,
	})
}
