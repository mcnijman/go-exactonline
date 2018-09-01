// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package project

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// TimeAndBillingEntryRecentProjectsEndpoint is responsible for communicating with
// the TimeAndBillingEntryRecentProjects endpoint of the Project service.
type TimeAndBillingEntryRecentProjectsEndpoint service

// TimeAndBillingEntryRecentProjects:
// Service: Project
// Entity: TimeAndBillingEntryRecentProjects
// URL: /api/v1/{division}/read/project/TimeAndBillingEntryRecentProjects
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadProjectTimeAndBillingEntryRecentProjects
type TimeAndBillingEntryRecentProjects struct {
	// ProjectId: Primary key
	ProjectId *types.GUID `json:"ProjectId,omitempty"`

	// DateLastUsed: Date last used
	DateLastUsed *types.Date `json:"DateLastUsed,omitempty"`

	// ProjectCode: Code of project
	ProjectCode *string `json:"ProjectCode,omitempty"`

	// ProjectDescription: Description of project
	ProjectDescription *string `json:"ProjectDescription,omitempty"`
}

func (s *TimeAndBillingEntryRecentProjects) GetIdentifier() types.GUID {
	return *s.ProjectId
}

// List the TimeAndBillingEntryRecentProjects entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *TimeAndBillingEntryRecentProjectsEndpoint) List(ctx context.Context, division int, all bool) ([]*TimeAndBillingEntryRecentProjects, error) {
	var entities []*TimeAndBillingEntryRecentProjects
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/TimeAndBillingEntryRecentProjects?$select=*", division)
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
