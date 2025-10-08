package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignsGet(w http.ResponseWriter, r *http.Request) {
	campaigns := h.Service.Repository.Get()
	render.Status(r, http.StatusOK)
	render.JSON(w, r, campaigns)
}
