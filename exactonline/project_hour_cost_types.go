// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// ProjectHourCostTypesService is responsible for communicating with
// the HourCostTypes endpoint of the Project service.
type ProjectHourCostTypesService service

// ProjectHourCostTypes:
// Service: Project
// Entity: HourCostTypes
// URL: /api/v1/{division}/read/project/HourCostTypes
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadProjectHourCostTypes
type ProjectHourCostTypes struct {
	// ItemId: Primary key
	ItemId *GUID `json:"ItemId,omitempty"`

	// ItemDescription: Description of Item
	ItemDescription *string `json:"ItemDescription,omitempty"`
}

func (s *ProjectHourCostTypes) GetIdentifier() GUID {
	return *s.ItemId
}

// List the HourCostTypes entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ProjectHourCostTypesService) List(ctx context.Context, division int, all bool) ([]*ProjectHourCostTypes, error) {
	var entities []*ProjectHourCostTypes
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/HourCostTypes?$select=*", division)
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

/* // Get the HourCostTypes enitity, by ItemId.
func (s *ProjectHourCostTypesService) Get(ctx context.Context, division int, id GUID) (*ProjectHourCostTypes, error) {
	var entities []*ProjectHourCostTypes
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/HourCostTypes?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d HourCostTypes entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
