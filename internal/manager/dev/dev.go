package dev

import (
	"github.com/pinpt/agent.next/internal/graphql"
	"github.com/pinpt/agent.next/sdk"
	"github.com/pinpt/go-common/log"
)

type devManager struct {
	logger log.Logger
}

var _ sdk.Manager = (*devManager)(nil)

// GraphQLManager returns a graphql manager instance
func (m *devManager) GraphQLManager() sdk.GraphQLClientManager {
	return graphql.New()
}

// CreateWebHook is used by the integration to create a webhook on behalf of the integration for a given customer and refid
func (m *devManager) CreateWebHook(customerID string, refID string) error {
	log.Error(m.logger, "cannot create a webhook in dev mode")
	return nil
}

// New will create a new dev sdk.Manager
func New(logger log.Logger) sdk.Manager {
	return &devManager{logger}
}