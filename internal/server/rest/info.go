package rest

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/olegtemek/portfolio-go/internal/dto"
	"github.com/olegtemek/portfolio-go/internal/response"
)

// @Summary createOrUpdate
// @Security BasicAuth
// @Tags info
// @Accept json
// @Produce json
// @Param input body dto.InfoCreateOrUpdateDto true "info"
// @Success 200 {object} response.InfoCreateOrUpdateResponse
// @Router /info [post]
func (h *Handler) CreateOrUpdateInfo(w http.ResponseWriter, r *http.Request) {
	h.Log = h.Log.With("Source", "InfoHandler:createOrUpdateInfo")

	var req dto.InfoCreateOrUpdateDto

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

	info, err := h.Services.Info.CreateOrUpdate(&req)
	if err != nil {
		h.Log.Error("ERROR", err)
		response.New(&w, r, fmt.Errorf("something went wrong: %s", err), 400)
		return
	}

	render.JSON(w, r, &response.InfoCreateOrUpdateResponse{
		Info:   info,
		Status: 200,
	})

}
