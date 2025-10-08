package endpoints

import (
	"errors"
	"net/http"

	"github.com/edmilsonmedeiross/emailn/internal/contract"
	internalerrors "github.com/edmilsonmedeiross/emailn/internal/domain/internal-errors"
	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) {
	var request contract.NewCampaignDTO
	render.DecodeJSON(r.Body, &request)

	id, err := h.Service.Create(request)
	if err != nil {
		if errors.Is(err, internalerrors.ErrSaveCampaignFailed) {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{"error: ": err.Error()})
			return
		} else {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]string{"id": id})
}
