package campaign

import (
	"testing"

	"github.com/edmilsonmedeiross/emailn/internal/contract"
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

func TestCreateCampaign(t *testing.T) {
	assert := assert.New(t)
	service := Service{}
	newCampaign := contract.NewCampaignDTO{
		Name:    "test",
		Content: "any content",
		Emails:  []string{"example@example.com"},
	}

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_SaveCampaign(t *testing.T) {
	repo := new(repositoryMock)
	newCampaign := contract.NewCampaignDTO{Name: "Test Campaign", Content: "Test Description", Emails: []string{"example@example.com"}}

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
