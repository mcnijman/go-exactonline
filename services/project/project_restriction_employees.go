// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package project

import (
	"context"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

// ProjectRestrictionEmployeesEndpoint is responsible for communicating with
// the ProjectRestrictionEmployees endpoint of the Project service.
type ProjectRestrictionEmployeesEndpoint service

// ProjectRestrictionEmployees:
// Service: Project
// Entity: ProjectRestrictionEmployees
// URL: /api/v1/{division}/project/ProjectRestrictionEmployees
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ProjectProjectRestrictionEmployees
type ProjectRestrictionEmployees struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Created: Date created
	Created *types.Date `json:"Created,omitempty"`

	// Creator: Creator user ID
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Creator name
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Employee: Employee linked to the restriction
	Employee *types.GUID `json:"Employee,omitempty"`

	// EmployeeFullName: Name of employee
	EmployeeFullName *string `json:"EmployeeFullName,omitempty"`

	// EmployeeHID: Readable ID of the employee
	EmployeeHID *int `json:"EmployeeHID,omitempty"`

	// Modified: Date modified
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: Modifier user ID
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Modifier name
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Project: Project linked to the restriction
	Project *types.GUID `json:"Project,omitempty"`

	// ProjectCode: Project code
	ProjectCode *string `json:"ProjectCode,omitempty"`

	// ProjectDescription: Project description
	ProjectDescription *string `json:"ProjectDescription,omitempty"`
}

// List the ProjectRestrictionEmployees entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ProjectRestrictionEmployeesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*ProjectRestrictionEmployees, error) {
	var entities []*ProjectRestrictionEmployees
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectRestrictionEmployees", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
