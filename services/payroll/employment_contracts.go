// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package payroll

import (
	"context"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

// EmploymentContractsEndpoint is responsible for communicating with
// the EmploymentContracts endpoint of the Payroll service.
type EmploymentContractsEndpoint service

// EmploymentContracts:
// Service: Payroll
// Entity: EmploymentContracts
// URL: /api/v1/{division}/payroll/EmploymentContracts
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=PayrollEmploymentContracts
type EmploymentContracts struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// ContractFlexPhase: Flexible employment contract phase
	ContractFlexPhase *int `json:"ContractFlexPhase,omitempty"`

	// ContractFlexPhaseDescription: Flexible employment contract phase description.
	ContractFlexPhaseDescription *string `json:"ContractFlexPhaseDescription,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Document: Document ID of the employment contract
	Document *types.GUID `json:"Document,omitempty"`

	// Employee: ID of employee
	Employee *types.GUID `json:"Employee,omitempty"`

	// EmployeeFullName: Name of employee
	EmployeeFullName *string `json:"EmployeeFullName,omitempty"`

	// EmployeeHID: Numeric ID of the employee
	EmployeeHID *int `json:"EmployeeHID,omitempty"`

	// EmployeeType: Type of employee. 1 - Employee, 2 - Contractor, 3 - Temporary, 4 - Student, 5 - Flexworker
	EmployeeType *int `json:"EmployeeType,omitempty"`

	// EmployeeTypeDescription: Employee type description
	EmployeeTypeDescription *string `json:"EmployeeTypeDescription,omitempty"`

	// Employment: Employment ID
	Employment *types.GUID `json:"Employment,omitempty"`

	// EmploymentHID: Numeric ID of the employment
	EmploymentHID *int `json:"EmploymentHID,omitempty"`

	// EndDate: End date of employment contract
	EndDate *types.Date `json:"EndDate,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Notes of employment contract
	Notes *string `json:"Notes,omitempty"`

	// ProbationEndDate: Employment probation end date
	ProbationEndDate *types.Date `json:"ProbationEndDate,omitempty"`

	// ProbationPeriod: Employment probation period
	ProbationPeriod *int `json:"ProbationPeriod,omitempty"`

	// ReasonContract: Employment contract reason code. 1 - New employment, 2 - Employment change, 3 - New legal employer, 4 - Acquisition 5 - Previous contract expired, 6 - Other
	ReasonContract *int `json:"ReasonContract,omitempty"`

	// ReasonContractDescription: Employment contract reason description
	ReasonContractDescription *string `json:"ReasonContractDescription,omitempty"`

	// Sequence: Sequence number
	Sequence *int `json:"Sequence,omitempty"`

	// StartDate: Start date of employment contract
	StartDate *types.Date `json:"StartDate,omitempty"`

	// Type: Type of employment contract. 1 - Definite, 2 - Indefinite, 3 - External
	Type *int `json:"Type,omitempty"`

	// TypeDescription: Description of employment contract type
	TypeDescription *string `json:"TypeDescription,omitempty"`
}

// List the EmploymentContracts entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *EmploymentContractsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*EmploymentContracts, error) {
	var entities []*EmploymentContracts
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/payroll/EmploymentContracts", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
