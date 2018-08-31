// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// HRMAbsenceRegistrationTransactionsService is responsible for communicating with
// the AbsenceRegistrationTransactions endpoint of the HRM service.
type HRMAbsenceRegistrationTransactionsService service

// HRMAbsenceRegistrationTransactions:
// Service: HRM
// Entity: AbsenceRegistrationTransactions
// URL: /api/v1/{division}/hrm/AbsenceRegistrationTransactions
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=HRMAbsenceRegistrationTransactions
type HRMAbsenceRegistrationTransactions struct {
	// ID: Primary key
	ID *GUID `json:",omitempty"`

	// AbsenceRegistration: Reference key to Absence Registration
	AbsenceRegistration *GUID `json:",omitempty"`

	// Created: Creation date
	Created *Date `json:",omitempty"`

	// Creator: User ID of creator
	Creator *GUID `json:",omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:",omitempty"`

	// Division: Division code
	Division *int `json:",omitempty"`

	// EndTime: End time on the last day of absence stored as DateTime, and the date should be ignored
	EndTime *Date `json:",omitempty"`

	// ExpectedEndDate: Expected end date of absence
	ExpectedEndDate *Date `json:",omitempty"`

	// Hours: Total number of absence hours
	Hours *float64 `json:",omitempty"`

	// HoursFirstDay: Hours of absence on the first day
	HoursFirstDay *float64 `json:",omitempty"`

	// HoursLastDay: Hours of absence on the last day
	HoursLastDay *float64 `json:",omitempty"`

	// Modified: Last modified date
	Modified *Date `json:",omitempty"`

	// Modifier: User ID of modifier
	Modifier *GUID `json:",omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:",omitempty"`

	// Notes: Extra information for absence
	Notes *string `json:",omitempty"`

	// NotificationMoment: Notification moment of absence
	NotificationMoment *Date `json:",omitempty"`

	// PercentageDisablement: Percentage disablement
	PercentageDisablement *float64 `json:",omitempty"`

	// StartDate: Start date of absence
	StartDate *Date `json:",omitempty"`

	// StartTime: Start time on the first day of absence stored as DateTime, and the date should be ignored
	StartTime *Date `json:",omitempty"`

	// Status: Status of absence, 0 = Open, 1 = Recovered
	Status *int `json:",omitempty"`
}

func (s *HRMAbsenceRegistrationTransactions) GetIdentifier() GUID {
	return *s.ID
}

// List the AbsenceRegistrationTransactions entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *HRMAbsenceRegistrationTransactionsService) List(ctx context.Context, division int, all bool) ([]*HRMAbsenceRegistrationTransactions, error) {
	var entities []*HRMAbsenceRegistrationTransactions
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/AbsenceRegistrationTransactions?$select=*", division)
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