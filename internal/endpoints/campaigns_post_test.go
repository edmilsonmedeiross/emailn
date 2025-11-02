package endpoints

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edmilsonmedeiross/emailn/internal/contract"
	"github.com/edmilsonmedeiross/emailn/internal/domain/campaign"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockCampaignService struct {
	mock.Mock
}

func (m *mockCampaignService) Create(cp *contract.NewCampaignDTO) (string, error) {
	args := m.Called(cp)
	return args.String(0), args.Error(1)
}

func (m *mockCampaignService) Get() []campaign.Campaign {
	args := m.Called()
	return args.Get(0).([]campaign.Campaign)
}

func (m *mockCampaignService) GetByID(id string) (*contract.GetByCampaignResponseDTO, error) {
	// args := m.Called(id)
	return nil, nil
}

func TestCampaignsPost(t *testing.T) {

	assert := assert.New(t)
	service := new(mockCampaignService)
	handler := Handler{Service: service}

	t.Run("should return 201 on successful campaign creation", func(t *testing.T) {

		body := contract.NewCampaignDTO{
			Name:    "Test Campaign",
			Content: "This is the content of the test campaign",
			Emails:  []string{"test1@example.com", "test2@example.com"},
		}

		service.On("Create", &body).Return("12345", nil)

		reqBody, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		w := httptest.NewRecorder()

		response, status, err := handler.CampaignPost(w, req)

		assert.NoError(err)
		assert.Equal(http.StatusCreated, status)
		assert.Equal(map[string]interface{}{"id": "12345"}, response)
	})

	t.Run("should handle errors gracefully", func(t *testing.T) {
		body := contract.NewCampaignDTO{
			Name:    "",
			Content: "",
			Emails:  []string{},
		}

		service.On("Create", &body).Return("", errors.New("invalid input"))

		reqBody, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		w := httptest.NewRecorder()

		response, _, err := handler.CampaignPost(w, req)

		assert.Error(err)
		assert.Equal(response, map[string]interface{}{"id": ""})
	})
}
