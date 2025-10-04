package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "Test Campaign"
	content := "This is a test campaign."
	contacts := []string{"email@one.com", "email@two.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Equal(len(contacts), len(campaign.Contacts))

	for i, contact := range campaign.Contacts {
		assert.Equal(contacts[i], contact.Email)
	}
}
