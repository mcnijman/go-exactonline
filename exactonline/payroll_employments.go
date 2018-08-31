// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// PayrollEmploymentsService is responsible for communicating with
// the Employments endpoint of the Payroll service.
type PayrollEmploymentsService service

// PayrollEmployments:
// Service: Payroll
// Entity: Employments
// URL: /api/v1/{division}/payroll/Employments
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=PayrollEmployments
type PayrollEmployments struct {
	// ID: Primary key
	ID *GUID `json:"ID,omitempty"`

	// Created: Creation date
	Created *Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Employee: Employee ID
	Employee *GUID `json:"Employee,omitempty"`

	// EmployeeFullName: Name of employee
	EmployeeFullName *string `json:"EmployeeFullName,omitempty"`

	// EmployeeHID: Numeric number of Employee
	EmployeeHID *int `json:"EmployeeHID,omitempty"`

	// EndDate: End date of employment
	EndDate *Date `json:"EndDate,omitempty"`

	// HID: Numeric ID of the employment
	HID *int `json:"HID,omitempty"`

	// Modified: Last modified date
	Modified *Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// ReasonEnd: ID of employment ended
	ReasonEnd *int `json:"ReasonEnd,omitempty"`

	// ReasonEndDescription: Reason of end of employment
	ReasonEndDescription *string `json:"ReasonEndDescription,omitempty"`

	// ReasonEndFlex: Reason of ended flexible employment
	ReasonEndFlex *int `json:"ReasonEndFlex,omitempty"`

	// ReasonEndFlexDescription: Other reason for end of employment
	ReasonEndFlexDescription *string `json:"ReasonEndFlexDescription,omitempty"`

	// StartDate: Start date of employment
	StartDate *Date `json:"StartDate,omitempty"`

	// StartDateOrganization: Start date of the employee in the organization. This field is used to count the years in service.
	StartDateOrganization *Date `json:"StartDateOrganization,omitempty"`
}

func (s *PayrollEmployments) GetIdentifier() GUID {
	return *s.ID
}

// List the Employments entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *PayrollEmploymentsService) List(ctx context.Context, division int, all bool) ([]*PayrollEmployments, error) {
	var entities []*PayrollEmployments
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/payroll/Employments?$select=*", division)
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

/* // Get the Employments enitity, by ID.
func (s *PayrollEmploymentsService) Get(ctx context.Context, division int, id GUID) (*PayrollEmployments, error) {
	var entities []*PayrollEmployments
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/payroll/Employments?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d Employments entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
