package main

import (
	"net/http"

	"github.com/edmilsonmedeiross/emailn/internal/contract"
	"github.com/edmilsonmedeiross/emailn/internal/domain/campaign"
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

	service := campaign.Service{}

	r.Post("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		var request contract.NewCampaignDTO
		render.DecodeJSON(r.Body, &request)

		id, err := service.Create(request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, map[string]string{"id": id})
	})

	http.ListenAndServe(":8080", r)
}
