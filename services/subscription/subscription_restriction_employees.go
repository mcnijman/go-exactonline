// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package subscription

import (
	"context"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

// SubscriptionRestrictionEmployeesEndpoint is responsible for communicating with
// the SubscriptionRestrictionEmployees endpoint of the Subscription service.
type SubscriptionRestrictionEmployeesEndpoint service

// SubscriptionRestrictionEmployees:
// Service: Subscription
// Entity: SubscriptionRestrictionEmployees
// URL: /api/v1/{division}/subscription/SubscriptionRestrictionEmployees
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SubscriptionSubscriptionRestrictionEmployees
type SubscriptionRestrictionEmployees struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of the creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of the creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Employee: Employee linked to the restriction
	Employee *types.GUID `json:"Employee,omitempty"`

	// EmployeeFullName: Name of employee
	EmployeeFullName *string `json:"EmployeeFullName,omitempty"`

	// EmployeeHID: Readable ID of employee
	EmployeeHID *int `json:"EmployeeHID,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of the last modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of the last modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Subscription: Reference to subscription
	Subscription *types.GUID `json:"Subscription,omitempty"`

	// SubscriptionDescription: Description of subscription
	SubscriptionDescription *string `json:"SubscriptionDescription,omitempty"`

	// SubscriptionNumber: Number of subscription
	SubscriptionNumber *int `json:"SubscriptionNumber,omitempty"`
}

// List the SubscriptionRestrictionEmployees entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SubscriptionRestrictionEmployeesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*SubscriptionRestrictionEmployees, error) {
	var entities []*SubscriptionRestrictionEmployees
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/SubscriptionRestrictionEmployees", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
