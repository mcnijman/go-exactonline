// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// PayrollTaxEmploymentEndFlexCodesService is responsible for communicating with
// the TaxEmploymentEndFlexCodes endpoint of the Payroll service.
type PayrollTaxEmploymentEndFlexCodesService service

// PayrollTaxEmploymentEndFlexCodes:
// Service: Payroll
// Entity: TaxEmploymentEndFlexCodes
// URL: /api/v1/{division}/payroll/TaxEmploymentEndFlexCodes
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=PayrollTaxEmploymentEndFlexCodes
type PayrollTaxEmploymentEndFlexCodes struct {
	// ID: Primary key
	ID *GUID `json:"ID,omitempty"`

	// Code: Code of flexible employment contract phase
	Code *string `json:"Code,omitempty"`

	// Created: Creation date
	Created *Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description of flexible employment contract phase
	Description *string `json:"Description,omitempty"`

	// EndDate: End date of flexible employment contract
	EndDate *Date `json:"EndDate,omitempty"`

	// Modified: Last modified date
	Modified *Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// StartDate: Start date of flexible employment contract phase
	StartDate *Date `json:"StartDate,omitempty"`
}

func (s *PayrollTaxEmploymentEndFlexCodes) GetIdentifier() GUID {
	return *s.ID
}

// List the TaxEmploymentEndFlexCodes entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *PayrollTaxEmploymentEndFlexCodesService) List(ctx context.Context, division int, all bool) ([]*PayrollTaxEmploymentEndFlexCodes, error) {
	var entities []*PayrollTaxEmploymentEndFlexCodes
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/payroll/TaxEmploymentEndFlexCodes?$select=*", division)
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

/* // Get the TaxEmploymentEndFlexCodes enitity, by ID.
func (s *PayrollTaxEmploymentEndFlexCodesService) Get(ctx context.Context, division int, id GUID) (*PayrollTaxEmploymentEndFlexCodes, error) {
	var entities []*PayrollTaxEmploymentEndFlexCodes
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/payroll/TaxEmploymentEndFlexCodes?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d TaxEmploymentEndFlexCodes entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
