package dev

import (
	"fmt"
	"net/url"

	"github.com/pinpt/agent.next/internal/graphql"
	"github.com/pinpt/agent.next/internal/http"
	"github.com/pinpt/agent.next/sdk"
	"github.com/pinpt/go-common/v10/api"
	"github.com/pinpt/go-common/v10/log"
)

type devManager struct {
	logger  log.Logger
	channel string
}

var _ sdk.Manager = (*devManager)(nil)

// GraphQLManager returns a graphql manager instance
func (m *devManager) GraphQLManager() sdk.GraphQLClientManager {
	return graphql.New()
}

// HTTPManager returns a HTTP manager instance
func (m *devManager) HTTPManager() sdk.HTTPClientManager {
	return http.New()
}

// CreateWebHook is used by the integration to create a webhook on behalf of the integration for a given customer and refid
func (m *devManager) CreateWebHook(customerID string, refType string, refID string) (string, error) {
	log.Error(m.logger, "cannot create a webhook in dev mode")
	return "", nil
}

// RefreshOAuth2Token will refresh the OAuth2 access token using the provided refreshToken and return a new access token
func (m *devManager) RefreshOAuth2Token(refType string, refreshToken string) (string, error) {
	theurl := api.BackendURL(api.AuthService, m.channel)
	theurl += fmt.Sprintf("oauth/%s/refresh/%s", refType, url.PathEscape(refreshToken))
	var res struct {
		AccessToken string `json:"access_token"`
	}
	client := http.New().New(theurl, map[string]string{"Content-Type": "application/json"})
	_, err := client.Get(&res)
	if err != nil {
		return "", err
	}
	return res.AccessToken, nil
}

// New will create a new dev sdk.Manager
func New(logger log.Logger, channel string) sdk.Manager {
	return &devManager{logger, channel}
}
