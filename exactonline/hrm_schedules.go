// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// HRMSchedulesService is responsible for communicating with
// the Schedules endpoint of the HRM service.
type HRMSchedulesService service

// HRMSchedules:
// Service: HRM
// Entity: Schedules
// URL: /api/v1/{division}/hrm/Schedules
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=HRMSchedules
type HRMSchedules struct {
	// ID: Primary key
	ID *GUID `json:",omitempty"`

	// Active: Obsolete
	Active *byte `json:",omitempty"`

	// AverageHours: Average hours per week in a schedule
	AverageHours *float64 `json:",omitempty"`

	// Code: Schedule code
	Code *string `json:",omitempty"`

	// Created: Creation date
	Created *Date `json:",omitempty"`

	// Creator: ID of creator
	Creator *GUID `json:",omitempty"`

	// CreatorFullName: Name of the creator
	CreatorFullName *string `json:",omitempty"`

	// Days: Average days per week in the schedule
	Days *float64 `json:",omitempty"`

	// Description: Description of the schedule
	Description *string `json:",omitempty"`

	// Division: Division code
	Division *int `json:",omitempty"`

	// Employment: Employment ID for schedule
	Employment *GUID `json:",omitempty"`

	// EmploymentHID: Employment number
	EmploymentHID *int `json:",omitempty"`

	// EndDate: End date of the schedule
	EndDate *Date `json:",omitempty"`

	// Hours: Number of hours per week in a CLA for which the schedule is built
	Hours *float64 `json:",omitempty"`

	// LeaveHoursCompensation: Number of hours which are built up each week for later leave
	LeaveHoursCompensation *float64 `json:",omitempty"`

	// Main: Indication if the schedule is a main schedule for a CLA. 1 = Yes, 0 = No
	Main *byte `json:",omitempty"`

	// Modified: Last modified date
	Modified *Date `json:",omitempty"`

	// Modifier: ID of modifier
	Modifier *GUID `json:",omitempty"`

	// ModifierFullName: Name of the modifier
	ModifierFullName *string `json:",omitempty"`

	// PaymentParttimeFactor: Part-time factor for payroll calculation. Value between 0 and 1
	PaymentParttimeFactor *float64 `json:",omitempty"`

	// ScheduleType: Type of schedule. 1 = Hours and average days, 2 = Hours and specific days, 3 = Hours per day, 4 = Time frames per day
	ScheduleType *int `json:",omitempty"`

	// ScheduleTypeDescription: Description of the schedule type
	ScheduleTypeDescription *string `json:",omitempty"`

	// StartDate: Week in the schedule which is used to start with. By default the number will be 1.
	StartDate *Date `json:",omitempty"`

	// StartWeek: Week to start the schedule from for an employee
	StartWeek *int `json:",omitempty"`
}

func (s *HRMSchedules) GetIdentifier() GUID {
	return *s.ID
}

// List the Schedules entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *HRMSchedulesService) List(ctx context.Context, division int, all bool) ([]*HRMSchedules, error) {
	var entities []*HRMSchedules
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/Schedules?$select=*", division)
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