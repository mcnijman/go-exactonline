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

// TimeAndBillingEntryRecentHourCostTypesEndpoint is responsible for communicating with
// the TimeAndBillingEntryRecentHourCostTypes endpoint of the Project service.
type TimeAndBillingEntryRecentHourCostTypesEndpoint service

// TimeAndBillingEntryRecentHourCostTypes:
// Service: Project
// Entity: TimeAndBillingEntryRecentHourCostTypes
// URL: /api/v1/{division}/read/project/TimeAndBillingEntryRecentHourCostTypes
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadProjectTimeAndBillingEntryRecentHourCostTypes
type TimeAndBillingEntryRecentHourCostTypes struct {
	// ItemId: Primary key
	ItemId *types.GUID `json:"ItemId,omitempty"`

	// DateLastUsed: Date last used
	DateLastUsed *types.Date `json:"DateLastUsed,omitempty"`

	// ItemDescription: Description of item
	ItemDescription *string `json:"ItemDescription,omitempty"`
}

func (s *TimeAndBillingEntryRecentHourCostTypes) GetIdentifier() types.GUID {
	return *s.ItemId
}

// List the TimeAndBillingEntryRecentHourCostTypes entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *TimeAndBillingEntryRecentHourCostTypesEndpoint) List(ctx context.Context, division int, all bool) ([]*TimeAndBillingEntryRecentHourCostTypes, error) {
	var entities []*TimeAndBillingEntryRecentHourCostTypes
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/TimeAndBillingEntryRecentHourCostTypes?$select=*", division)
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
