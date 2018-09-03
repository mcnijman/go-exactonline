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

// ProjectRestrictionRebillingsEndpoint is responsible for communicating with
// the ProjectRestrictionRebillings endpoint of the Project service.
type ProjectRestrictionRebillingsEndpoint service

// ProjectRestrictionRebillings:
// Service: Project
// Entity: ProjectRestrictionRebillings
// URL: /api/v1/{division}/project/ProjectRestrictionRebillings
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ProjectProjectRestrictionRebillings
type ProjectRestrictionRebillings struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// CostTypeRebill: Cost type reference
	CostTypeRebill *types.GUID `json:"CostTypeRebill,omitempty"`

	// CostTypeRebillCode: Cost type code
	CostTypeRebillCode *string `json:"CostTypeRebillCode,omitempty"`

	// CostTypeRebillDescription: Cost type description
	CostTypeRebillDescription *string `json:"CostTypeRebillDescription,omitempty"`

	// Created: Date created
	Created *types.Date `json:"Created,omitempty"`

	// Creator: Creator user ID
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Creator name
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

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

// List the ProjectRestrictionRebillings entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ProjectRestrictionRebillingsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*ProjectRestrictionRebillings, error) {
	var entities []*ProjectRestrictionRebillings
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectRestrictionRebillings", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
