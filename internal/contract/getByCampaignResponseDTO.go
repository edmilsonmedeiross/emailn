package contract

type GetByCampaignResponseDTO struct {
	CampaignID   string `json:"campaign_id"`
	CampaignName string `json:"campaign_name"`
	Status       string `json:"status"`
}
