package database

import (
	"github.com/edmilsonmedeiross/emailn/internal/domain/campaign"
)

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (r *CampaignRepository) Save(campaign *campaign.Campaign) error {
	r.campaigns = append(r.campaigns, *campaign)
	return nil
}
