// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package hrm

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// LeaveBuildUpRegistrationsEndpoint is responsible for communicating with
// the LeaveBuildUpRegistrations endpoint of the HRM service.
type LeaveBuildUpRegistrationsEndpoint service

// LeaveBuildUpRegistrations:
// Service: HRM
// Entity: LeaveBuildUpRegistrations
// URL: /api/v1/{division}/hrm/LeaveBuildUpRegistrations
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=HRMLeaveBuildUpRegistrations
type LeaveBuildUpRegistrations struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Date: Date of leave build up
	Date *types.Date `json:"Date,omitempty"`

	// Description: Description of leave build up
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Employee: Employee linked to the leave build up
	Employee *types.GUID `json:"Employee,omitempty"`

	// EmployeeFullName: Employee full name
	EmployeeFullName *string `json:"EmployeeFullName,omitempty"`

	// EmployeeHID: Numeric ID of the employee
	EmployeeHID *int `json:"EmployeeHID,omitempty"`

	// Hours: Total number of leave build up hours
	Hours *float64 `json:"Hours,omitempty"`

	// LeaveType: Type of leave
	LeaveType *types.GUID `json:"LeaveType,omitempty"`

	// LeaveTypeCode: Code for type of leave
	LeaveTypeCode *string `json:"LeaveTypeCode,omitempty"`

	// LeaveTypeDescription: Description for type of leave
	LeaveTypeDescription *string `json:"LeaveTypeDescription,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Extra information for leave build up
	Notes *string `json:"Notes,omitempty"`

	// Status: Status of leave build up, 1 = Submitted, 2 = Approved, 3 = Rejected
	Status *int `json:"Status,omitempty"`
}

func (s *LeaveBuildUpRegistrations) GetIdentifier() types.GUID {
	return *s.ID
}

// List the LeaveBuildUpRegistrations entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *LeaveBuildUpRegistrationsEndpoint) List(ctx context.Context, division int, all bool) ([]*LeaveBuildUpRegistrations, error) {
	var entities []*LeaveBuildUpRegistrations
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/LeaveBuildUpRegistrations?$select=*", division)
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
