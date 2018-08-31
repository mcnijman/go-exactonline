// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// ContinuousMonitoringIndicatorStatesService is responsible for communicating with
// the IndicatorStates endpoint of the ContinuousMonitoring service.
type ContinuousMonitoringIndicatorStatesService service

// ContinuousMonitoringIndicatorStates:
// Service: ContinuousMonitoring
// Entity: IndicatorStates
// URL: /api/v1/beta/{division}/continuousmonitoring/IndicatorStates
// HasWebhook: false
// IsInBeta: true
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ContinuousMonitoringIndicatorStates
type ContinuousMonitoringIndicatorStates struct {
	// ID: Primary key
	ID *GUID `json:",omitempty"`

	// Count: To store the number (e.g. 2) of occurrences of an indicator (e.g. Number of deviating entries: 2)
	Count *int `json:",omitempty"`

	// Created: Creation date
	Created *Date `json:",omitempty"`

	// Creator: User ID of creator
	Creator *GUID `json:",omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:",omitempty"`

	// Division: Division code
	Division *int `json:",omitempty"`

	// Indicator: ID of Indicators
	Indicator *GUID `json:",omitempty"`

	// IndicatorDescription: Indicator type description
	IndicatorDescription *string `json:",omitempty"`

	// IndicatorType: Indicator type (1 = Balance G/L account per financial year, 2 = Usage of journals, 3 = Deviating amount entered, 4 = Liquidity, 5 = VAT Return deadline, 6 = Difference result in percentage, 7 = Different VAT code used)
	IndicatorType *int `json:",omitempty"`

	// LastUpdated: Last update date
	LastUpdated *Date `json:",omitempty"`

	// Modified: Last modified date
	Modified *Date `json:",omitempty"`

	// Modifier: User ID of modifier
	Modifier *GUID `json:",omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:",omitempty"`

	// ReportingYear: Financial year
	ReportingYear *int `json:",omitempty"`

	// Status: Indicator status (1 = OK, 2 = Warning, 3 = Exception)
	Status *int `json:",omitempty"`

	// Value: To store a value (e.g. -1234.56) that will be used by the indicators current situation (e.g. Lowest expected balance of liquid assets will be: -1,234.56)
	Value *float64 `json:",omitempty"`
}

func (s *ContinuousMonitoringIndicatorStates) GetIdentifier() GUID {
	return *s.ID
}

// List the IndicatorStates entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ContinuousMonitoringIndicatorStatesService) List(ctx context.Context, division int, all bool) ([]*ContinuousMonitoringIndicatorStates, error) {
	var entities []*ContinuousMonitoringIndicatorStates
	u, err := s.client.ResolvePathWithDivision("/api/v1/beta/{division}/continuousmonitoring/IndicatorStates?$select=*", division)
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