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
// @Tags experience
// @Accept json
// @Produce json
// @Param input body dto.ExperienceCreateDto true "Experience"
// @Success 200 {object} response.ExperienceCreateResponse
// @Router /experience [post]
func (h *Handler) CreateExperience(w http.ResponseWriter, r *http.Request) {
	h.Log = h.Log.With("Source", "ExperienceHandler:createExperience")

	var req dto.ExperienceCreateDto

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

	experience, err := h.Services.Experience.Create(&req)

	if err != nil {
		h.Log.Error("ERROR", err)
		response.New(&w, r, fmt.Errorf("something went wrong: %s", err), 400)
		return
	}

	render.JSON(w, r, &response.ExperienceCreateResponse{
		Experience: experience,
		Status:     200,
	})
}

// @Summary update
// @Security BasicAuth
// @Tags experience
// @Accept json
// @Produce json
// @Param id path int true "Experience"
// @Param input body dto.ExperienceUpdateDto true "Experience"
// @Success 200 {object} response.ExperienceUpdateResponse
// @Router /experience/{id} [patch]
func (h *Handler) UpdateExperience(w http.ResponseWriter, r *http.Request) {
	h.Log = h.Log.With("Source", "ExperienceHandler:updateExperience")

	var req dto.ExperienceUpdateDto

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

	experience, err := h.Services.Experience.Update(idInt, &req)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			response.New(&w, r, fmt.Errorf("experience not found: %s", err), 404)
			return
		}

		response.New(&w, r, fmt.Errorf("something went wrong: %s", err), 400)
		return
	}

	render.JSON(w, r, &response.ExperienceUpdateResponse{
		Experience: experience,
		Status:     200,
	})
}

// @Summary delete
// @Security BasicAuth
// @Tags experience
// @Accept json
// @Produce json
// @Param id path int true "Experience id"
// @Success 200 {object} response.ExperienceDeleteResponse
// @Router /experience/{id} [delete]
func (h *Handler) DeleteExperience(w http.ResponseWriter, r *http.Request) {
	h.Log = h.Log.With("Source", "ExperienceHandler:deleteExperience")
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.Log.Error("ERROR", err)
		response.New(&w, r, fmt.Errorf("cannot parse id: %s", err), 400)
		return
	}

	experience, err := h.Services.Experience.Delete(idInt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			response.New(&w, r, fmt.Errorf("experience not found: %s", err), 404)
			return
		}

		response.New(&w, r, fmt.Errorf("something went wrong: %s", err), 400)
		return
	}

	render.JSON(w, r, &response.ExperienceDeleteResponse{
		Experience: experience,
		Status:     200,
	})
}
