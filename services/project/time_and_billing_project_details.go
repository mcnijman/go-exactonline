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

// TimeAndBillingProjectDetailsEndpoint is responsible for communicating with
// the TimeAndBillingProjectDetails endpoint of the Project service.
type TimeAndBillingProjectDetailsEndpoint service

// TimeAndBillingProjectDetails:
// Service: Project
// Entity: TimeAndBillingProjectDetails
// URL: /api/v1/{division}/read/project/TimeAndBillingProjectDetails
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadProjectTimeAndBillingProjectDetails
type TimeAndBillingProjectDetails struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Account: The account for this project
	Account *types.GUID `json:"Account,omitempty"`

	// AccountName:
	AccountName *string `json:"AccountName,omitempty"`

	// Code: Code of project
	Code *string `json:"Code,omitempty"`

	// Description: Description of the project
	Description *string `json:"Description,omitempty"`

	// Type: Reference to ProjectTypes
	Type *int `json:"Type,omitempty"`
}

func (s *TimeAndBillingProjectDetails) GetIdentifier() types.GUID {
	return *s.ID
}

// List the TimeAndBillingProjectDetails entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *TimeAndBillingProjectDetailsEndpoint) List(ctx context.Context, division int, all bool) ([]*TimeAndBillingProjectDetails, error) {
	var entities []*TimeAndBillingProjectDetails
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