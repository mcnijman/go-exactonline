// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package payroll

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// ActiveEmploymentsEndpoint is responsible for communicating with
// the ActiveEmployments endpoint of the Payroll service.
type ActiveEmploymentsEndpoint service

// ActiveEmployments:
// Service: Payroll
// Entity: ActiveEmployments
// URL: /api/v1/{division}/payroll/ActiveEmployments
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=PayrollActiveEmployments
type ActiveEmployments struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// AverageDaysPerWeek: The average number of contract days that an employee works per week
	AverageDaysPerWeek *float64 `json:"AverageDaysPerWeek,omitempty"`

	// AverageHoursPerWeek: The average number of contract hours that an employee works per week
	AverageHoursPerWeek *float64 `json:"AverageHoursPerWeek,omitempty"`

	// Contract: Employment contract ID
	Contract *types.GUID `json:"Contract,omitempty"`

	// ContractDocument: Document ID of the employment contract
	ContractDocument *types.GUID `json:"ContractDocument,omitempty"`

	// ContractEndDate: End date of employment contract
	ContractEndDate *types.Date `json:"ContractEndDate,omitempty"`

	// ContractProbationEndDate: Employment probation end date
	ContractProbationEndDate *types.Date `json:"ContractProbationEndDate,omitempty"`

	// ContractProbationPeriod: Employment probation period
	ContractProbationPeriod *int `json:"ContractProbationPeriod,omitempty"`

	// ContractStartDate: Start date of employment contract
	ContractStartDate *types.Date `json:"ContractStartDate,omitempty"`

	// ContractType: Type of employment contract. 1 - Definite, 2 - Indefinite, 3 - External
	ContractType *int `json:"ContractType,omitempty"`

	// ContractTypeDescription: Description of employment contract type
	ContractTypeDescription *string `json:"ContractTypeDescription,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Department: Department of employee
	Department *types.GUID `json:"Department,omitempty"`

	// DepartmentCode: Department code of employee
	DepartmentCode *string `json:"DepartmentCode,omitempty"`

	// DepartmentDescription: Description of department
	DepartmentDescription *string `json:"DepartmentDescription,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Employee: Employee ID
	Employee *types.GUID `json:"Employee,omitempty"`

	// EmployeeFullName: Name of employee
	EmployeeFullName *string `json:"EmployeeFullName,omitempty"`

	// EmployeeHID: Numeric number of Employee
	EmployeeHID *int `json:"EmployeeHID,omitempty"`

	// EmploymentOrganization: Organization of employment
	EmploymentOrganization *types.GUID `json:"EmploymentOrganization,omitempty"`

	// EndDate: End date of employment
	EndDate *types.Date `json:"EndDate,omitempty"`

	// HID: Numeric ID of the employment
	HID *int `json:"HID,omitempty"`

	// HourlyWage: Hourly wage
	HourlyWage *float64 `json:"HourlyWage,omitempty"`

	// InternalRate: Internal rate for time &amp; billing or professional service user
	InternalRate *float64 `json:"InternalRate,omitempty"`

	// Jobtitle: Job title of employee
	Jobtitle *types.GUID `json:"Jobtitle,omitempty"`

	// JobtitleDescription: Description of job title
	JobtitleDescription *string `json:"JobtitleDescription,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

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

	// Salary: Employment salary
	Salary *types.GUID `json:"Salary,omitempty"`

	// Schedule: Work schedule
	Schedule *types.GUID `json:"Schedule,omitempty"`

	// ScheduleAverageHours: Average hours per week in a schedule.
	ScheduleAverageHours *float64 `json:"ScheduleAverageHours,omitempty"`

	// ScheduleCode: Work schedule code
	ScheduleCode *string `json:"ScheduleCode,omitempty"`

	// ScheduleDays: Number of days of work per week
	ScheduleDays *float64 `json:"ScheduleDays,omitempty"`

	// ScheduleDescription: Description of work schedule
	ScheduleDescription *string `json:"ScheduleDescription,omitempty"`

	// ScheduleHours: Number of work hours per week.
	ScheduleHours *float64 `json:"ScheduleHours,omitempty"`

	// StartDate: Start date of employment
	StartDate *types.Date `json:"StartDate,omitempty"`

	// StartDateOrganization: Start date of the employee in the organization. This field is used to count the years in service.
	StartDateOrganization *types.Date `json:"StartDateOrganization,omitempty"`
}

func (s *ActiveEmployments) GetIdentifier() types.GUID {
	return *s.ID
}

// List the ActiveEmployments entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ActiveEmploymentsEndpoint) List(ctx context.Context, division int, all bool) ([]*ActiveEmployments, error) {
	var entities []*ActiveEmployments
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/payroll/ActiveEmployments?$select=*", division)
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
