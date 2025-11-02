package campaign

import (
	"github.com/edmilsonmedeiross/emailn/internal/contract"
	internalerrors "github.com/edmilsonmedeiross/emailn/internal/domain/internal-errors"
)

type Service interface {
	Create(campaign *contract.NewCampaignDTO) (string, error)
	Get() []Campaign
	GetByID(id string) (*contract.GetByCampaignResponseDTO, error)
}

type ServiceImpl struct {
	Repository Repository
}

func (s *ServiceImpl) Create(campaign *contract.NewCampaignDTO) (string, error) {
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

func (s *ServiceImpl) Get() []Campaign {
	campaigns := s.Repository.Get()

	return campaigns
}

func (s *ServiceImpl) GetByID(id string) (*contract.GetByCampaignResponseDTO, error) {
	cp, err := s.Repository.GetByID(id)

	if err != nil {
		return nil, internalerrors.ErrSaveCampaignFailed
	}

	return &contract.GetByCampaignResponseDTO{
		CampaignID:   cp.ID,
		CampaignName: cp.Name,
		Status:       cp.Status,
	}, nil
}
