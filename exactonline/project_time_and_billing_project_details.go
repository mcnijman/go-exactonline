// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// ProjectTimeAndBillingProjectDetailsService is responsible for communicating with
// the TimeAndBillingProjectDetails endpoint of the Project service.
type ProjectTimeAndBillingProjectDetailsService service

// ProjectTimeAndBillingProjectDetails:
// Service: Project
// Entity: TimeAndBillingProjectDetails
// URL: /api/v1/{division}/read/project/TimeAndBillingProjectDetails
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadProjectTimeAndBillingProjectDetails
type ProjectTimeAndBillingProjectDetails struct {
	// ID: Primary key
	ID *GUID `json:"ID,omitempty"`

	// Account: The account for this project
	Account *GUID `json:"Account,omitempty"`

	// AccountName:
	AccountName *string `json:"AccountName,omitempty"`

	// Code: Code of project
	Code *string `json:"Code,omitempty"`

	// Description: Description of the project
	Description *string `json:"Description,omitempty"`

	// Type: Reference to ProjectTypes
	Type *int `json:"Type,omitempty"`
}

func (s *ProjectTimeAndBillingProjectDetails) GetIdentifier() GUID {
	return *s.ID
}

// List the TimeAndBillingProjectDetails entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ProjectTimeAndBillingProjectDetailsService) List(ctx context.Context, division int, all bool) ([]*ProjectTimeAndBillingProjectDetails, error) {
	var entities []*ProjectTimeAndBillingProjectDetails
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/TimeAndBillingProjectDetails?$select=*", division)
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

/* // Get the TimeAndBillingProjectDetails enitity, by ID.
func (s *ProjectTimeAndBillingProjectDetailsService) Get(ctx context.Context, division int, id GUID) (*ProjectTimeAndBillingProjectDetails, error) {
	var entities []*ProjectTimeAndBillingProjectDetails
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/TimeAndBillingProjectDetails?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d TimeAndBillingProjectDetails entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
