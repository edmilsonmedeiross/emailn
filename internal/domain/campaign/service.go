package campaign

import (
	"github.com/edmilsonmedeiross/emailn/internal/contract"
	internalerrors "github.com/edmilsonmedeiross/emailn/internal/domain/internal-errors"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(campaign contract.NewCampaignDTO) (string, error) {
	newCampaign, err := NewCampaign(campaign.Name, campaign.Content, campaign.Emails)
	if err != nil {
		return "", err
	}

	err = s.Repository.Save(newCampaign)
	if err != nil {
		return "", internalerrors.ErrSaveCampaignFailed
	}

	return newCampaign.ID, nil
}
