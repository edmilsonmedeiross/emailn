package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Test Campaign"
	content  = "This is a test campaign."
	contacts = []string{"email@one.com", "email@two.com"}
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Equal(len(contacts), len(campaign.Contacts))

	for i, contact := range campaign.Contacts {
		assert.Equal(contacts[i], contact.Email)
	}
}

func TestNewCampaignID(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(t, campaign.ID)
}

func Test_NewCampaign_MustValidateName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)
	assert.Equal("name is required", err.Error())
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)
	assert.Equal("content is required", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})
	assert.Equal("contacts is required", err.Error())
}

func Test_NewCampaign_MustCreateCampaignWithCreatedOn(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.CreatedOn)
	assert.False(campaign.CreatedOn.IsZero())
}
