// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// HRMDepartmentsService is responsible for communicating with
// the Departments endpoint of the HRM service.
type HRMDepartmentsService service

// HRMDepartments:
// Service: HRM
// Entity: Departments
// URL: /api/v1/{division}/hrm/Departments
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=HRMDepartments
type HRMDepartments struct {
	// ID: Primary key
	ID *GUID `json:"ID,omitempty"`

	// Code: Department Code
	Code *string `json:"Code,omitempty"`

	// Costcenter: Cost center Code
	Costcenter *string `json:"Costcenter,omitempty"`

	// CostcenterDescription: Cost center description
	CostcenterDescription *string `json:"CostcenterDescription,omitempty"`

	// Created: Creation date
	Created *Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Department description
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Modified: Last modified date
	Modified *Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Explanation or extra information can be stored in the notes
	Notes *string `json:"Notes,omitempty"`
}

func (s *HRMDepartments) GetIdentifier() GUID {
	return *s.ID
}

// List the Departments entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *HRMDepartmentsService) List(ctx context.Context, division int, all bool) ([]*HRMDepartments, error) {
	var entities []*HRMDepartments
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/Departments?$select=*", division)
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

/* // Get the Departments enitity, by ID.
func (s *HRMDepartmentsService) Get(ctx context.Context, division int, id GUID) (*HRMDepartments, error) {
	var entities []*HRMDepartments
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/Departments?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d Departments entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
