// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// BudgetBudgetScenariosService is responsible for communicating with
// the BudgetScenarios endpoint of the Budget service.
type BudgetBudgetScenariosService service

// BudgetBudgetScenarios:
// Service: Budget
// Entity: BudgetScenarios
// URL: /api/v1/beta/{division}/budget/BudgetScenarios
// HasWebhook: false
// IsInBeta: true
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=BudgetBudgetScenarios
type BudgetBudgetScenarios struct {
	// ID: Primary key
	ID *GUID `json:",omitempty"`

	// Code: Budget scenario code
	Code *string `json:",omitempty"`

	// Created: Creation date
	Created *Date `json:",omitempty"`

	// Creator: User ID of creator
	Creator *GUID `json:",omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:",omitempty"`

	// Description: Budget scenario description
	Description *string `json:",omitempty"`

	// Division: Division code
	Division *int `json:",omitempty"`

	// FromYear: From year
	FromYear *int `json:",omitempty"`

	// Modified: Last modified date
	Modified *Date `json:",omitempty"`

	// Modifier: User ID of modifier
	Modifier *GUID `json:",omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:",omitempty"`

	// ToYear: To year
	ToYear *int `json:",omitempty"`
}

func (s *BudgetBudgetScenarios) GetIdentifier() GUID {
	return *s.ID
}

// List the BudgetScenarios entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *BudgetBudgetScenariosService) List(ctx context.Context, division int, all bool) ([]*BudgetBudgetScenarios, error) {
	var entities []*BudgetBudgetScenarios
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