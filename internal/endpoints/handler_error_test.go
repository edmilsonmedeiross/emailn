package endpoints

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	internalerrors "github.com/edmilsonmedeiross/emailn/internal/domain/internal-errors"
	"github.com/stretchr/testify/assert"
)

func TestHandlerErrorShouldReturnInternalServerError(t *testing.T) {
	assert := assert.New(t)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/some-endpoint", nil)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, http.StatusInternalServerError, internalerrors.ErrSaveCampaignFailed
	}

	handlerFunc := HandlerError(endpoint)

	handlerFunc.ServeHTTP(w, r)

	assert.Equal(http.StatusInternalServerError, w.Code)
	assert.Contains(w.Body.String(), internalerrors.ErrSaveCampaignFailed.Error())
}

func TestHandlerErrorShouldReturnBadRequest(t *testing.T) {
	assert := assert.New(t)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/endpoint", nil)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, http.StatusBadRequest, errors.New("bad request")
	}

	handler := HandlerError(endpoint)
	handler.ServeHTTP(w, r)

	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Contains(w.Body.String(), "bad request")
}
