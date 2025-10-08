package main

import (
	"net/http"

	"github.com/edmilsonmedeiross/emailn/internal/domain/campaign"
	"github.com/edmilsonmedeiross/emailn/internal/endpoints"
	"github.com/edmilsonmedeiross/emailn/internal/infra/database"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)

	repository := &database.CampaignRepository{}
	service := &campaign.Service{Repository: repository}
	handler := &endpoints.Handler{Service: service}

	r.Post("/campaigns", handler.CampaignPost)
	r.Get("/campaigns", handler.CampaignsGet)

	http.ListenAndServe(":8080", r)
}
