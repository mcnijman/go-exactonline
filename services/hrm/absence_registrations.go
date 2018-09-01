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

// AbsenceRegistrationsEndpoint is responsible for communicating with
// the AbsenceRegistrations endpoint of the HRM service.
type AbsenceRegistrationsEndpoint service

// AbsenceRegistrations:
// Service: HRM
// Entity: AbsenceRegistrations
// URL: /api/v1/{division}/hrm/AbsenceRegistrations
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=HRMAbsenceRegistrations
type AbsenceRegistrations struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// AbsenceRegistrationTransactions: Collection of absence registration transactions
	AbsenceRegistrationTransactions *[]byte `json:"AbsenceRegistrationTransactions,omitempty"`

	// Cause: Absence cause, only supported for the Netherland legislation
	Cause *int `json:"Cause,omitempty"`

	// CauseCode: Code for the absence cause, only supported for the Netherland legislation
	CauseCode *string `json:"CauseCode,omitempty"`

	// CauseDescription: Description for the absence cause, only supported for the Netherland legislation
	CauseDescription *string `json:"CauseDescription,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Employee: Employee linked to the absence
	Employee *types.GUID `json:"Employee,omitempty"`

	// EmployeeFullName: Employee full name
	EmployeeFullName *string `json:"EmployeeFullName,omitempty"`

	// EmployeeHID: Numeric ID of the employee
	EmployeeHID *int `json:"EmployeeHID,omitempty"`

	// Kind: Absence kind, only supported for the Netherland legislation
	Kind *int `json:"Kind,omitempty"`

	// KindCode: Code for the absence kind, only supported for the Netherland legislation
	KindCode *string `json:"KindCode,omitempty"`

	// KindDescription: Description for the absence kind, only supported for the Netherland legislation
	KindDescription *string `json:"KindDescription,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Extra information for absence
	Notes *string `json:"Notes,omitempty"`
}

func (s *AbsenceRegistrations) GetIdentifier() types.GUID {
	return *s.ID
}

// List the AbsenceRegistrations entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *AbsenceRegistrationsEndpoint) List(ctx context.Context, division int, all bool) ([]*AbsenceRegistrations, error) {
	var entities []*AbsenceRegistrations
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/AbsenceRegistrations?$select=*", division)
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
