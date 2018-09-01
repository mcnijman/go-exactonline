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

// TimeAndBillingActivitiesAndExpensesEndpoint is responsible for communicating with
// the TimeAndBillingActivitiesAndExpenses endpoint of the Project service.
type TimeAndBillingActivitiesAndExpensesEndpoint service

// TimeAndBillingActivitiesAndExpenses:
// Service: Project
// Entity: TimeAndBillingActivitiesAndExpenses
// URL: /api/v1/{division}/read/project/TimeAndBillingActivitiesAndExpenses
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadProjectTimeAndBillingActivitiesAndExpenses
type TimeAndBillingActivitiesAndExpenses struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Description: Description
	Description *string `json:"Description,omitempty"`

	// ParentDescription: Description of Parent
	ParentDescription *string `json:"ParentDescription,omitempty"`
}

func (s *TimeAndBillingActivitiesAndExpenses) GetIdentifier() types.GUID {
	return *s.ID
}

// List the TimeAndBillingActivitiesAndExpenses entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *TimeAndBillingActivitiesAndExpensesEndpoint) List(ctx context.Context, division int, all bool) ([]*TimeAndBillingActivitiesAndExpenses, error) {
	var entities []*TimeAndBillingActivitiesAndExpenses
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/TimeAndBillingActivitiesAndExpenses?$select=*", division)
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