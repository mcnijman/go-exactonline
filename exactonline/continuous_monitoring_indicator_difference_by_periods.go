// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// ContinuousMonitoringIndicatorDifferenceByPeriodsService is responsible for communicating with
// the IndicatorDifferenceByPeriods endpoint of the ContinuousMonitoring service.
type ContinuousMonitoringIndicatorDifferenceByPeriodsService service

// ContinuousMonitoringIndicatorDifferenceByPeriods:
// Service: ContinuousMonitoring
// Entity: IndicatorDifferenceByPeriods
// URL: /api/v1/beta/{division}/continuousmonitoring/IndicatorDifferenceByPeriods
// HasWebhook: false
// IsInBeta: true
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ContinuousMonitoringIndicatorDifferenceByPeriods
type ContinuousMonitoringIndicatorDifferenceByPeriods struct {
	// ID: Primary key
	ID *GUID `json:"ID,omitempty"`

	// AccountType: Indicates GL Account types (1 = Revenue, 2 = Cost, 3 = Result)
	AccountType *int `json:"AccountType,omitempty"`

	// Active: Indicates if this indicator is active or inactive
	Active *byte `json:"Active,omitempty"`

	// AfterEntry: Indicates whether it will be validated immediately after an entry
	AfterEntry *byte `json:"AfterEntry,omitempty"`

	// BudgetScenario: ID of Budget scenario (This property is only used when FinPeriod = 3)
	BudgetScenario *GUID `json:"BudgetScenario,omitempty"`

	// Classification: Indicator classification (1 = Quality, 2 = Advice). Default = 1
	Classification *int `json:"Classification,omitempty"`

	// CompareWith: Compare with (1 = Last year, 2 = Last year until period, 3 = Budget)
	CompareWith *int `json:"CompareWith,omitempty"`

	// Created: Creation date
	Created *Date `json:"Created,omitempty"`

	// CreateSignal: Indicates whether a signal is created
	CreateSignal *byte `json:"CreateSignal,omitempty"`

	// Creator: User ID of creator
	Creator *GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description of indicator
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// ExternalCode: External code
	ExternalCode *string `json:"ExternalCode,omitempty"`

	// Modified: Last modified date
	Modified *Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Operator: Operator to be used in conjunction with journal (1 = Equal to, 2 = Unequal to, 3 = Greater than, 4 = Greater than or equal to, 5 = Less than, 6 = Less than or equal to, 7 = Between)
	Operator *int `json:"Operator,omitempty"`

	// Severity: Severity of the indicators (1 = Low, 2 = Medium, 3 = High, 4 = Critical)
	Severity *int `json:"Severity,omitempty"`

	// Type: Indicator type (1 = Balance G/L account per financial year, 2 = Usage of journals, 3 = Deviating amount entered, 4 = Liquidity, 5 = VAT Return deadline, 6 = Difference result in percentage, 7 = Different VAT code used)
	Type *int `json:"Type,omitempty"`

	// ValueFrom: Value from/Value. Default value is 0. This field should be used together with any choice of operator.
	ValueFrom *float64 `json:"ValueFrom,omitempty"`

	// ValueTo: Value to. Default value is 0.
	ValueTo *float64 `json:"ValueTo,omitempty"`
}

func (s *ContinuousMonitoringIndicatorDifferenceByPeriods) GetIdentifier() GUID {
	return *s.ID
}

// List the IndicatorDifferenceByPeriods entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ContinuousMonitoringIndicatorDifferenceByPeriodsService) List(ctx context.Context, division int, all bool) ([]*ContinuousMonitoringIndicatorDifferenceByPeriods, error) {
	var entities []*ContinuousMonitoringIndicatorDifferenceByPeriods
	u, err := s.client.ResolvePathWithDivision("/api/v1/beta/{division}/continuousmonitoring/IndicatorDifferenceByPeriods?$select=*", division)
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

/* // Get the IndicatorDifferenceByPeriods enitity, by ID.
func (s *ContinuousMonitoringIndicatorDifferenceByPeriodsService) Get(ctx context.Context, division int, id GUID) (*ContinuousMonitoringIndicatorDifferenceByPeriods, error) {
	var entities []*ContinuousMonitoringIndicatorDifferenceByPeriods
	u, err := s.client.ResolvePathWithDivision("/api/v1/beta/{division}/continuousmonitoring/IndicatorDifferenceByPeriods?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d IndicatorDifferenceByPeriods entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
