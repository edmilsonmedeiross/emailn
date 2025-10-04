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

	campaign := NewCampaign(name, content, contacts)

	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Equal(len(contacts), len(campaign.Contacts))

	for i, contact := range campaign.Contacts {
		assert.Equal(contacts[i], contact.Email)
	}
}

func TestNewCampaignID(t *testing.T) {
	assert := assert.New(t)

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(t, campaign.ID)
}
