// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// ContinuousMonitoringIndicatorGLAccountsService is responsible for communicating with
// the IndicatorGLAccounts endpoint of the ContinuousMonitoring service.
type ContinuousMonitoringIndicatorGLAccountsService service

// ContinuousMonitoringIndicatorGLAccounts:
// Service: ContinuousMonitoring
// Entity: IndicatorGLAccounts
// URL: /api/v1/beta/{division}/continuousmonitoring/IndicatorGLAccounts
// HasWebhook: false
// IsInBeta: true
// Methods: GET POST DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ContinuousMonitoringIndicatorGLAccounts
type ContinuousMonitoringIndicatorGLAccounts struct {
	// ID: Primary key
	ID *GUID `json:"ID,omitempty"`

	// GLAccount: ID of GLAccount
	GLAccount *GUID `json:"GLAccount,omitempty"`

	// GLAccountCode: GL account code
	GLAccountCode *string `json:"GLAccountCode,omitempty"`

	// GLAccountDescription: Description of GLAccount
	GLAccountDescription *string `json:"GLAccountDescription,omitempty"`

	// Indicator: ID of Indicators
	Indicator *GUID `json:"Indicator,omitempty"`
}

func (s *ContinuousMonitoringIndicatorGLAccounts) GetIdentifier() GUID {
	return *s.ID
}

// List the IndicatorGLAccounts entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ContinuousMonitoringIndicatorGLAccountsService) List(ctx context.Context, division int, all bool) ([]*ContinuousMonitoringIndicatorGLAccounts, error) {
	var entities []*ContinuousMonitoringIndicatorGLAccounts
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

/* // Get the IndicatorGLAccounts enitity, by ID.
func (s *ContinuousMonitoringIndicatorGLAccountsService) Get(ctx context.Context, division int, id GUID) (*ContinuousMonitoringIndicatorGLAccounts, error) {
	var entities []*ContinuousMonitoringIndicatorGLAccounts
	u, err := s.client.ResolvePathWithDivision("/api/v1/beta/{division}/continuousmonitoring/IndicatorGLAccounts?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d IndicatorGLAccounts entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
