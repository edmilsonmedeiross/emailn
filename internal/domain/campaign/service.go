package campaign

import "github.com/edmilsonmedeiross/emailn/internal/contract"

type Service struct {
	Repository Repository
}

func (s *Service) Create(campaign contract.NewCampaignDTO) (string, error) {
	newCampaign, _ := NewCampaign(campaign.Name, campaign.Content, campaign.Emails)
	s.Repository.Save(newCampaign)

	return newCampaign.ID, nil
}
