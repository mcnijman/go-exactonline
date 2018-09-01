// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package continuousmonitoring

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// IndicatorGLAccountsEndpoint is responsible for communicating with
// the IndicatorGLAccounts endpoint of the ContinuousMonitoring service.
type IndicatorGLAccountsEndpoint service

// IndicatorGLAccounts:
// Service: ContinuousMonitoring
// Entity: IndicatorGLAccounts
// URL: /api/v1/beta/{division}/continuousmonitoring/IndicatorGLAccounts
// HasWebhook: false
// IsInBeta: true
// Methods: GET POST DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ContinuousMonitoringIndicatorGLAccounts
type IndicatorGLAccounts struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// GLAccount: ID of GLAccount
	GLAccount *types.GUID `json:"GLAccount,omitempty"`

	// GLAccountCode: GL account code
	GLAccountCode *string `json:"GLAccountCode,omitempty"`

	// GLAccountDescription: Description of GLAccount
	GLAccountDescription *string `json:"GLAccountDescription,omitempty"`

	// Indicator: ID of Indicators
	Indicator *types.GUID `json:"Indicator,omitempty"`
}

func (s *IndicatorGLAccounts) GetIdentifier() types.GUID {
	return *s.ID
}

// List the IndicatorGLAccounts entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *IndicatorGLAccountsEndpoint) List(ctx context.Context, division int, all bool) ([]*IndicatorGLAccounts, error) {
	var entities []*IndicatorGLAccounts
	u, err := s.client.ResolvePathWithDivision("/api/v1/beta/{division}/continuousmonitoring/IndicatorGLAccounts?$select=*", division)
	if err != nil {
		return nil, err
	}
	if all {
		err = s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err = s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
