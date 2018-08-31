// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// HRMDivisionClassesService is responsible for communicating with
// the DivisionClasses endpoint of the HRM service.
type HRMDivisionClassesService service

// HRMDivisionClasses:
// Service: HRM
// Entity: DivisionClasses
// URL: /api/v1/{division}/hrm/DivisionClasses
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=HRMDivisionClasses
type HRMDivisionClasses struct {
	// ID: Primary key
	ID *GUID `json:"ID,omitempty"`

	// ClassNameCustomer: Classification customer ID
	ClassNameCustomer *GUID `json:"ClassNameCustomer,omitempty"`

	// ClassNameDescription: Related classification description
	ClassNameDescription *string `json:"ClassNameDescription,omitempty"`

	// ClassNameID: Related classification ID
	ClassNameID *GUID `json:"ClassNameID,omitempty"`

	// Code: Property code
	Code *string `json:"Code,omitempty"`

	// Created: Creation date
	Created *Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Property description
	Description *string `json:"Description,omitempty"`

	// DescriptionTermID: Property description term ID
	DescriptionTermID *int `json:"DescriptionTermID,omitempty"`

	// Modified: Last modified date
	Modified *Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// SequenceNr: Related classification sequence number
	SequenceNr *int `json:"SequenceNr,omitempty"`
}

func (s *HRMDivisionClasses) GetIdentifier() GUID {
	return *s.ID
}

// List the DivisionClasses entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *HRMDivisionClassesService) List(ctx context.Context, division int, all bool) ([]*HRMDivisionClasses, error) {
	var entities []*HRMDivisionClasses
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/DivisionClasses?$select=*", division)
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

/* // Get the DivisionClasses enitity, by ID.
func (s *HRMDivisionClassesService) Get(ctx context.Context, division int, id GUID) (*HRMDivisionClasses, error) {
	var entities []*HRMDivisionClasses
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/DivisionClasses?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d DivisionClasses entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
