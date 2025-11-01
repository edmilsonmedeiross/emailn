package endpoints

import (
	"net/http"

	"github.com/edmilsonmedeiross/emailn/internal/contract"
	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contract.NewCampaignDTO
	render.DecodeJSON(r.Body, &request)

	id, err := h.Service.Create(request)

	return map[string]interface{}{"id": id}, http.StatusCreated, err
}
