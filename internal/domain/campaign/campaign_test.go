package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
	name := "Test Campaign"
	content := "This is a test campaign."
	contacts := []string{"email@one.com", "email@two.com"}
	
	campaign := NewCampaign(name, content, contacts);

	if campaign.ID != "1" {
		t.Errorf("Expected campaign ID to be 1, but got %s", campaign.ID)
	}

	if campaign.Name != name {
		t.Errorf("Expected campaign name to be %s, but got %s", name, campaign.Name)
	}

	if campaign.Content != content {
		t.Errorf("Expected campaign content to be %s, but got %s", content, campaign.Content)
	}

	if len(campaign.Contacts) != len(contacts) {
		t.Errorf("Expected %d contacts, but got %d", len(contacts), len(campaign.Contacts))
	}

	for i, contact := range campaign.Contacts {
		if contact.Email != contacts[i] {
			t.Errorf("Expected contact email to be %s, but got %s", contacts[i], contact.Email)
		}
	}
}	