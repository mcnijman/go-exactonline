// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package budget

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// BudgetScenariosEndpoint is responsible for communicating with
// the BudgetScenarios endpoint of the Budget service.
type BudgetScenariosEndpoint service

// BudgetScenarios:
// Service: Budget
// Entity: BudgetScenarios
// URL: /api/v1/beta/{division}/budget/BudgetScenarios
// HasWebhook: false
// IsInBeta: true
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=BudgetBudgetScenarios
type BudgetScenarios struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Code: Budget scenario code
	Code *string `json:"Code,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Budget scenario description
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// FromYear: From year
	FromYear *int `json:"FromYear,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// ToYear: To year
	ToYear *int `json:"ToYear,omitempty"`
}

func (s *BudgetScenarios) GetIdentifier() types.GUID {
	return *s.ID
}

// List the BudgetScenarios entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *BudgetScenariosEndpoint) List(ctx context.Context, division int, all bool) ([]*BudgetScenarios, error) {
	var entities []*BudgetScenarios
	u, err := s.client.ResolvePathWithDivision("/api/v1/beta/{division}/budget/BudgetScenarios?$select=*", division)
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