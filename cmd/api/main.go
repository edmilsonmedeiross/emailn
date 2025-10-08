package main

import (
	"errors"
	"net/http"

	"github.com/edmilsonmedeiross/emailn/internal/contract"
	"github.com/edmilsonmedeiross/emailn/internal/domain/campaign"
	internalerrors "github.com/edmilsonmedeiross/emailn/internal/domain/internal-errors"
	"github.com/edmilsonmedeiross/emailn/internal/infra/database"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)

	repository := &database.CampaignRepository{}
	service := campaign.Service{Repository: repository}

	r.Post("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		var request contract.NewCampaignDTO
		render.DecodeJSON(r.Body, &request)

		id, err := service.Create(request)
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
	})

	http.ListenAndServe(":8080", r)
}
