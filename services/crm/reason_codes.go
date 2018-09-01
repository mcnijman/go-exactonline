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

// ReasonCodesEndpoint is responsible for communicating with
// the ReasonCodes endpoint of the CRM service.
type ReasonCodesEndpoint service

// ReasonCodes:
// Service: CRM
// Entity: ReasonCodes
// URL: /api/v1/{division}/crm/ReasonCodes
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=CRMReasonCodes
type ReasonCodes struct {
	// ID: Primary key.
	ID *types.GUID `json:"ID,omitempty"`

	// Active: Indicates if the reason code is active.
	Active *byte `json:"Active,omitempty"`

	// Code: Code of the reason.
	Code *string `json:"Code,omitempty"`

	// Created: Creation date.
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator.
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator.
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description of the reason code.
	Description *string `json:"Description,omitempty"`

	// Division: Division code.
	Division *int `json:"Division,omitempty"`

	// Modified: Last modified date.
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier.
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier.
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Extra notes.
	Notes *string `json:"Notes,omitempty"`

	// Type: Type of the reason code.
	Type *int `json:"Type,omitempty"`

	// TypeDescription: Description of the type of the reason code.
	TypeDescription *string `json:"TypeDescription,omitempty"`
}

func (s *ReasonCodes) GetIdentifier() types.GUID {
	return *s.ID
}

// List the ReasonCodes entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ReasonCodesEndpoint) List(ctx context.Context, division int, all bool) ([]*ReasonCodes, error) {
	var entities []*ReasonCodes
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/ReasonCodes?$select=*", division)
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
