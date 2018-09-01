// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package hrm

import "github.com/mcnijman/go-exactonline/api"

type service struct {
	client *api.Client
}

// HRMService is responsible for communication with the HRM
// endpoints of the Exact Online API.
type HRMService struct {
	client *api.Client

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Endpoints available under this service
	AbsenceRegistrations            *AbsenceRegistrationsEndpoint
	AbsenceRegistrationTransactions *AbsenceRegistrationTransactionsEndpoint
	Costcenters                     *CostcentersEndpoint
	Costunits                       *CostunitsEndpoint
	Departments                     *DepartmentsEndpoint
	DivisionClasses                 *DivisionClassesEndpoint
	DivisionClassNames              *DivisionClassNamesEndpoint
	DivisionClassValues             *DivisionClassValuesEndpoint
	Divisions                       *DivisionsEndpoint
	JobGroups                       *JobGroupsEndpoint
	JobTitles                       *JobTitlesEndpoint
	LeaveBuildUpRegistrations       *LeaveBuildUpRegistrationsEndpoint
	LeaveRegistrations              *LeaveRegistrationsEndpoint
	Schedules                       *SchedulesEndpoint
}

// NewHRMService creates a new initialized instance of the
// HRMService.
func NewHRMService(apiClient *api.Client) *HRMService {
	s := &HRMService{client: apiClient}

	s.common.client = apiClient

	s.AbsenceRegistrations = (*AbsenceRegistrationsEndpoint)(&s.common)
	s.AbsenceRegistrationTransactions = (*AbsenceRegistrationTransactionsEndpoint)(&s.common)
	s.Costcenters = (*CostcentersEndpoint)(&s.common)
	s.Costunits = (*CostunitsEndpoint)(&s.common)
	s.Departments = (*DepartmentsEndpoint)(&s.common)
	s.DivisionClasses = (*DivisionClassesEndpoint)(&s.common)
	s.DivisionClassNames = (*DivisionClassNamesEndpoint)(&s.common)
	s.DivisionClassValues = (*DivisionClassValuesEndpoint)(&s.common)
	s.Divisions = (*DivisionsEndpoint)(&s.common)
	s.JobGroups = (*JobGroupsEndpoint)(&s.common)
	s.JobTitles = (*JobTitlesEndpoint)(&s.common)
	s.LeaveBuildUpRegistrations = (*LeaveBuildUpRegistrationsEndpoint)(&s.common)
	s.LeaveRegistrations = (*LeaveRegistrationsEndpoint)(&s.common)
	s.Schedules = (*SchedulesEndpoint)(&s.common)

	return s
}
