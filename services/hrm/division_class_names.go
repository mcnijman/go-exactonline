// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package hrm

import (
	"context"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

// DivisionClassNamesEndpoint is responsible for communicating with
// the DivisionClassNames endpoint of the HRM service.
type DivisionClassNamesEndpoint service

// DivisionClassNames:
// Service: HRM
// Entity: DivisionClassNames
// URL: /api/v1/{division}/hrm/DivisionClassNames
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=HRMDivisionClassNames
type DivisionClassNames struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Customer: ID of customer
	Customer *types.GUID `json:"Customer,omitempty"`

	// Description: Description of classification
	Description *string `json:"Description,omitempty"`

	// DescriptionTermID: Term ID of the classification
	DescriptionTermID *int `json:"DescriptionTermID,omitempty"`

	// DivisionClasses: Collection of classification properties
	DivisionClasses *[]byte `json:"DivisionClasses,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// SequenceNr: Sequence number
	SequenceNr *int `json:"SequenceNr,omitempty"`
}

// List the DivisionClassNames entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *DivisionClassNamesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*DivisionClassNames, error) {
	var entities []*DivisionClassNames
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/DivisionClassNames", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
