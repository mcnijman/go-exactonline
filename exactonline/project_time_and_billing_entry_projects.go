// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// ProjectTimeAndBillingEntryProjectsService is responsible for communicating with
// the TimeAndBillingEntryProjects endpoint of the Project service.
type ProjectTimeAndBillingEntryProjectsService service

// ProjectTimeAndBillingEntryProjects:
// Service: Project
// Entity: TimeAndBillingEntryProjects
// URL: /api/v1/{division}/read/project/TimeAndBillingEntryProjects
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadProjectTimeAndBillingEntryProjects
type ProjectTimeAndBillingEntryProjects struct {
	// ProjectId: Primary key
	ProjectId *GUID `json:"ProjectId,omitempty"`

	// ProjectCode: Code
	ProjectCode *string `json:"ProjectCode,omitempty"`

	// ProjectDescription: Description
	ProjectDescription *string `json:"ProjectDescription,omitempty"`
}

func (s *ProjectTimeAndBillingEntryProjects) GetIdentifier() GUID {
	return *s.ProjectId
}

// List the TimeAndBillingEntryProjects entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ProjectTimeAndBillingEntryProjectsService) List(ctx context.Context, division int, all bool) ([]*ProjectTimeAndBillingEntryProjects, error) {
	var entities []*ProjectTimeAndBillingEntryProjects
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/TimeAndBillingEntryProjects?$select=*", division)
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

/* // Get the TimeAndBillingEntryProjects enitity, by ProjectId.
func (s *ProjectTimeAndBillingEntryProjectsService) Get(ctx context.Context, division int, id GUID) (*ProjectTimeAndBillingEntryProjects, error) {
	var entities []*ProjectTimeAndBillingEntryProjects
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/TimeAndBillingEntryProjects?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d TimeAndBillingEntryProjects entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
