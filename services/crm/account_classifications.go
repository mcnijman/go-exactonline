// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package crm

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// AccountClassificationsEndpoint is responsible for communicating with
// the AccountClassifications endpoint of the CRM service.
type AccountClassificationsEndpoint service

// AccountClassifications:
// Service: CRM
// Entity: AccountClassifications
// URL: /api/v1/{division}/crm/AccountClassifications
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=CRMAccountClassifications
type AccountClassifications struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// AccountClassificationName: Reference to Account classification name
	AccountClassificationName *types.GUID `json:"AccountClassificationName,omitempty"`

	// AccountClassificationNameDescription: Description of AccountClassificationName
	AccountClassificationNameDescription *string `json:"AccountClassificationNameDescription,omitempty"`

	// Code: Account classification code
	Code *string `json:"Code,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`
}

func (s *AccountClassifications) GetIdentifier() types.GUID {
	return *s.ID
}

// List the AccountClassifications entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *AccountClassificationsEndpoint) List(ctx context.Context, division int, all bool) ([]*AccountClassifications, error) {
	var entities []*AccountClassifications
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/AccountClassifications?$select=*", division)
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