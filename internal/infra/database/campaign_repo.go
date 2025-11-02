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

func (r *CampaignRepository) Get() []campaign.Campaign {
	return r.campaigns
}

func (r *CampaignRepository) GetByID(id string) (*campaign.Campaign, error) {

	return nil, nil
}
