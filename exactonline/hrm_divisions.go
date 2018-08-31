// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// HRMDivisionsService is responsible for communicating with
// the Divisions endpoint of the HRM service.
type HRMDivisionsService service

// HRMDivisions:
// Service: HRM
// Entity: Divisions
// URL: /api/v1/{division}/hrm/Divisions
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=HRMDivisions
type HRMDivisions struct {
	// Code: Primary key
	Code *int `json:"Code,omitempty"`

	// BlockingStatus: Values: 0 = Not blocked 1 = Backup 2 = Conversion busy 3 = Conversion shadow 4 = Conversion waiting 5 = Copy data waiting 6 = Copy data buzy 100 = Wait for deletion 101 = Deleted 102 = Deletion failed
	BlockingStatus *int `json:"BlockingStatus,omitempty"`

	// Class_01: First division classification. User should have access rights to view division classifications.
	Class_01 *[]byte `json:"Class_01,omitempty"`

	// Class_02: Second division classification. User should have access rights to view division classifications.
	Class_02 *[]byte `json:"Class_02,omitempty"`

	// Class_03: Third division classification. User should have access rights to view division classifications.
	Class_03 *[]byte `json:"Class_03,omitempty"`

	// Class_04: Fourth division classification. User should have access rights to view division classifications.
	Class_04 *[]byte `json:"Class_04,omitempty"`

	// Class_05: Fifth division classification. User should have access rights to view division classifications.
	Class_05 *[]byte `json:"Class_05,omitempty"`

	// Country: Country of the division. Is used for determination of legislation
	Country *string `json:"Country,omitempty"`

	// CountryDescription: Description of Country
	CountryDescription *string `json:"CountryDescription,omitempty"`

	// Created: Creation date
	Created *Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of the creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Currency: Default currency of the division
	Currency *string `json:"Currency,omitempty"`

	// CurrencyDescription: Description of Currency
	CurrencyDescription *string `json:"CurrencyDescription,omitempty"`

	// Customer: Owner account of the division
	Customer *GUID `json:"Customer,omitempty"`

	// CustomerCode: Owner account code of the division
	CustomerCode *string `json:"CustomerCode,omitempty"`

	// CustomerName: Owner account name of the division
	CustomerName *string `json:"CustomerName,omitempty"`

	// Description: Description
	Description *string `json:"Description,omitempty"`

	// HID: Number that customers give to the division
	HID *int64 `json:"HID,omitempty"`

	// Main: True for the main (hosting) division
	Main *bool `json:"Main,omitempty"`

	// Modified: Last modified date
	Modified *Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of the last modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// SiretNumber: Siret Number of the division (France)
	SiretNumber *string `json:"SiretNumber,omitempty"`

	// StartDate: Date on which the division becomes active
	StartDate *Date `json:"StartDate,omitempty"`

	// Status: Regular administrations will have status 0.  Currently, the only other possibility is &#39;archived&#39; (1), which means the administration is not actively used, but still needs to be accessible for the customer/accountant to meet legal obligations
	Status *int `json:"Status,omitempty"`

	// TaxOfficeNumber: Number of your local tax authority (Germany)
	TaxOfficeNumber *string `json:"TaxOfficeNumber,omitempty"`

	// TaxReferenceNumber: Local tax reference number (Germany)
	TaxReferenceNumber *string `json:"TaxReferenceNumber,omitempty"`

	// VATNumber: VAT number
	VATNumber *string `json:"VATNumber,omitempty"`

	// Website: Customer value, hyperlink to external website
	Website *string `json:"Website,omitempty"`
}

func (s *HRMDivisions) GetIdentifier() int {
	return *s.Code
}

// List the Divisions entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *HRMDivisionsService) List(ctx context.Context, division int, all bool) ([]*HRMDivisions, error) {
	var entities []*HRMDivisions
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/Divisions?$select=*", division)
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

/* // Get the Divisions enitity, by Code.
func (s *HRMDivisionsService) Get(ctx context.Context, division int, id int) (*HRMDivisions, error) {
	var entities []*HRMDivisions
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/Divisions?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d Divisions entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
