// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// SystemAccountantInfoService is responsible for communicating with
// the AccountantInfo endpoint of the System service.
type SystemAccountantInfoService service

// SystemAccountantInfo:
// Service: System
// Entity: AccountantInfo
// URL: /api/v1/{division}/system/AccountantInfo
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SystemSystemAccountantInfo
type SystemAccountantInfo struct {
	// ID: The account ID of the accountant.
	ID *GUID `json:"ID,omitempty"`

	// AddressLine1: First address line.
	AddressLine1 *string `json:"AddressLine1,omitempty"`

	// AddressLine2: Second address line.
	AddressLine2 *string `json:"AddressLine2,omitempty"`

	// AddressLine3: Third address line.
	AddressLine3 *string `json:"AddressLine3,omitempty"`

	// City: City.
	City *string `json:"City,omitempty"`

	// Email: Email.
	Email *string `json:"Email,omitempty"`

	// IsAccountant: Indicates if the customer is an accountant himself.
	IsAccountant *bool `json:"IsAccountant,omitempty"`

	// Logo: Logo.
	Logo *[]byte `json:"Logo,omitempty"`

	// MenuLogoUrl: Url to retrieve the logo of the accountant.
	MenuLogoUrl *string `json:"MenuLogoUrl,omitempty"`

	// Name: The name of the accountant.
	Name *string `json:"Name,omitempty"`

	// Phone: Phone.
	Phone *string `json:"Phone,omitempty"`

	// Postcode: Postcode.
	Postcode *string `json:"Postcode,omitempty"`

	// Website: Website.
	Website *string `json:"Website,omitempty"`
}

func (s *SystemAccountantInfo) GetIdentifier() GUID {
	return *s.ID
}

// List the AccountantInfo entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SystemAccountantInfoService) List(ctx context.Context, division int, all bool) ([]*SystemAccountantInfo, error) {
	var entities []*SystemAccountantInfo
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/system/AccountantInfo?$select=*", division)
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

/* // Get the AccountantInfo enitity, by ID.
func (s *SystemAccountantInfoService) Get(ctx context.Context, division int, id GUID) (*SystemAccountantInfo, error) {
	var entities []*SystemAccountantInfo
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/system/AccountantInfo?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d AccountantInfo entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
