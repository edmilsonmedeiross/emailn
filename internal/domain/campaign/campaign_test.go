package campaign

import (
	"testing"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "Test Campaign"
	content  = "This is a test campaign."
	contacts = []string{"email@one.com", "email@two.com"}
	fake     = faker.New()
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

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)
	assert.Equal("the field name must have a minimum value of 5", err.Error())
}

func Test_ValidateCampaignNameMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts)
	assert.Equal("the field name must have a maximum value of 24", err.Error())
}

func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)
	assert.Equal("the field content must have a minimum value of 5", err.Error())
}

func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(1040), contacts)
	assert.Equal("the field content must have a maximum value of 1024", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})
	assert.Equal("the field contacts must have a minimum value of 1", err.Error())
}

func Test_NewCampaign_MustCreateCampaignWithCreatedOn(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.CreatedOn)
	assert.False(campaign.CreatedOn.IsZero())
}
