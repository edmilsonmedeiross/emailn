package campaign

import (
	"errors"
	"testing"

	"github.com/edmilsonmedeiross/emailn/internal/contract"
	internalerrors "github.com/edmilsonmedeiross/emailn/internal/domain/internal-errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	newCampaign = contract.NewCampaignDTO{
		Name:    "teste",
		Content: "any content",
		Emails:  []string{"example@example.com"},
	}

	service = Service{}
)

func TestCreateCampaign(t *testing.T) {
	assert := assert.New(t)
	repo := new(repositoryMock)
	repo.On("Save", mock.Anything).Return(nil)

	service.Repository = repo

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_SaveCampaign(t *testing.T) {
	repo := new(repositoryMock)

	repo.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name || campaign.Content != newCampaign.Content || len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}
		return true
	})).Return(nil)

	service := Service{Repository: repo}
	id, err := service.Create(newCampaign)

	assert.NotNil(t, id)
	assert.Nil(t, err)
	repo.AssertCalled(t, "Save", mock.Anything)
	repo.AssertExpectations(t)
}

func Test_ErrorOnSaveCampaign(t *testing.T) {
	assert := assert.New(t)
	repo := new(repositoryMock)
	repo.On("Save", mock.Anything).Return(errors.New("internal server error"))

	service := Service{Repository: repo}

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(err, internalerrors.ErrSaveCampaignFailed))
}
