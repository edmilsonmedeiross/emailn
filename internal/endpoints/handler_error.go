package endpoints

import (
	"errors"
	"net/http"

	internalerrors "github.com/edmilsonmedeiross/emailn/internal/domain/internal-errors"
	"github.com/go-chi/render"
)

type EndpointFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

// Tirar duvidas com will
func HandlerError(endpointFunc EndpointFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, status, err := endpointFunc(w, r)

		if err != nil {
			if errors.Is(err, internalerrors.ErrSaveCampaignFailed) {
				render.Status(r, http.StatusInternalServerError)
			} else {
				render.Status(r, http.StatusBadRequest)
			}
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}

		render.Status(r, status)
		if obj != nil {
			render.JSON(w, r, obj)
		}
	})
}
