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

// ProjectPlanningEndpoint is responsible for communicating with
// the ProjectPlanning endpoint of the Project service.
type ProjectPlanningEndpoint service

// ProjectPlanning:
// Service: Project
// Entity: ProjectPlanning
// URL: /api/v1/{division}/project/ProjectPlanning
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ProjectProjectPlanning
type ProjectPlanning struct {
	// ID:
	ID *types.GUID `json:"ID,omitempty"`

	// Account:
	Account *types.GUID `json:"Account,omitempty"`

	// AccountCode:
	AccountCode *string `json:"AccountCode,omitempty"`

	// AccountName:
	AccountName *string `json:"AccountName,omitempty"`

	// BGTStatus:
	BGTStatus *int `json:"BGTStatus,omitempty"`

	// CommunicationErrorStatus:
	CommunicationErrorStatus *int `json:"CommunicationErrorStatus,omitempty"`

	// Created:
	Created *types.Date `json:"Created,omitempty"`

	// Creator:
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description:
	Description *string `json:"Description,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// Employee:
	Employee *types.GUID `json:"Employee,omitempty"`

	// EmployeeCode:
	EmployeeCode *string `json:"EmployeeCode,omitempty"`

	// EmployeeHID:
	EmployeeHID *int `json:"EmployeeHID,omitempty"`

	// EndDate:
	EndDate *types.Date `json:"EndDate,omitempty"`

	// Hours:
	Hours *float64 `json:"Hours,omitempty"`

	// HourType:
	HourType *types.GUID `json:"HourType,omitempty"`

	// HourTypeCode:
	HourTypeCode *string `json:"HourTypeCode,omitempty"`

	// HourTypeDescription:
	HourTypeDescription *string `json:"HourTypeDescription,omitempty"`

	// IsBrokenRecurrence:
	IsBrokenRecurrence *bool `json:"IsBrokenRecurrence,omitempty"`

	// Modified:
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes:
	Notes *string `json:"Notes,omitempty"`

	// OverAllocate:
	OverAllocate *bool `json:"OverAllocate,omitempty"`

	// Project:
	Project *types.GUID `json:"Project,omitempty"`

	// ProjectCode:
	ProjectCode *string `json:"ProjectCode,omitempty"`

	// ProjectDescription:
	ProjectDescription *string `json:"ProjectDescription,omitempty"`

	// ProjectPlanningRecurring:
	ProjectPlanningRecurring *types.GUID `json:"ProjectPlanningRecurring,omitempty"`

	// ProjectWBS:
	ProjectWBS *types.GUID `json:"ProjectWBS,omitempty"`

	// ProjectWBSDescription:
	ProjectWBSDescription *string `json:"ProjectWBSDescription,omitempty"`

	// StartDate:
	StartDate *types.Date `json:"StartDate,omitempty"`

	// Status:
	Status *int `json:"Status,omitempty"`

	// Type:
	Type *int `json:"Type,omitempty"`
}

// List the ProjectPlanning entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ProjectPlanningEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*ProjectPlanning, error) {
	var entities []*ProjectPlanning
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectPlanning", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
