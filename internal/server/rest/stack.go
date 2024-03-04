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
// @Tags stack
// @Accept json
// @Produce json
// @Param input body dto.StackCreateDto true "stack body"
// @Success 200 {object} response.StackCreateResponse
// @Router /stack [post]
func (h *Handler) CreateStack(w http.ResponseWriter, r *http.Request) {
	h.Log = h.Log.With("Source", "StackHandler:create")

	var req dto.StackCreateDto

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

	stack, err := h.Services.Stack.Create(&req)

	if err != nil {
		h.Log.Error("ERROR", err)
		response.New(&w, r, fmt.Errorf("something went wrong: %s", err), 400)
		return
	}

	render.JSON(w, r, &response.StackCreateResponse{
		Stack:  stack,
		Status: 200,
	})
}

// @Summary create
// @Security BasicAuth
// @Tags stack
// @Accept json
// @Produce json
// @Param id path int true "stack"
// @Success 200 {object} response.StackDeleteResponse
// @Router /stack/{id} [delete]
func (h *Handler) DeleteStack(w http.ResponseWriter, r *http.Request) {
	h.Log = h.Log.With("Source", "StackHandler:deleteStack")
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.Log.Error("ERROR", err)
		response.New(&w, r, fmt.Errorf("cannot parse id: %s", err), 400)
		return
	}

	stack, err := h.Services.Stack.Delete(idInt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			response.New(&w, r, fmt.Errorf("stack not found: %s", err), 404)
			return
		}

		response.New(&w, r, fmt.Errorf("something went wrong: %s", err), 400)
		return
	}

	render.JSON(w, r, &response.StackDeleteResponse{
		Stack:  stack,
		Status: 200,
	})
}
