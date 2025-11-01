package endpoints

import (
	"net/http"
)

func (h *Handler) CampaignsGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaigns := h.Service.Repository.Get()

	return map[string]interface{}{"campaigns": campaigns}, http.StatusOK, nil
}
