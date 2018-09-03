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

// ProjectHourBudgetsEndpoint is responsible for communicating with
// the ProjectHourBudgets endpoint of the Project service.
type ProjectHourBudgetsEndpoint service

// ProjectHourBudgets:
// Service: Project
// Entity: ProjectHourBudgets
// URL: /api/v1/{division}/project/ProjectHourBudgets
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ProjectProjectHourBudgets
type ProjectHourBudgets struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Budget: Number of hours
	Budget *float64 `json:"Budget,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division number
	Division *int `json:"Division,omitempty"`

	// Item: Hour type of budget
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode: Code of hour type
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription: Description of hour type
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Project: Reference to project
	Project *types.GUID `json:"Project,omitempty"`

	// ProjectCode: Code of project
	ProjectCode *string `json:"ProjectCode,omitempty"`

	// ProjectDescription: Description of project
	ProjectDescription *string `json:"ProjectDescription,omitempty"`
}

// List the ProjectHourBudgets entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ProjectHourBudgetsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*ProjectHourBudgets, error) {
	var entities []*ProjectHourBudgets
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectHourBudgets", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
