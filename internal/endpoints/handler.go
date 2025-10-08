package endpoints

import "github.com/edmilsonmedeiross/emailn/internal/domain/campaign"

type Handler struct {
	Service *campaign.Service
}
