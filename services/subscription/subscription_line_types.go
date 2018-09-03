// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package subscription

import (
	"context"

	"github.com/mcnijman/go-exactonline/api"
)

// SubscriptionLineTypesEndpoint is responsible for communicating with
// the SubscriptionLineTypes endpoint of the Subscription service.
type SubscriptionLineTypesEndpoint service

// SubscriptionLineTypes:
// Service: Subscription
// Entity: SubscriptionLineTypes
// URL: /api/v1/{division}/subscription/SubscriptionLineTypes
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SubscriptionSubscriptionLineTypes
type SubscriptionLineTypes struct {
	// ID: Primary key
	ID *int `json:"ID,omitempty"`

	// Description: Description
	Description *string `json:"Description,omitempty"`
}

// List the SubscriptionLineTypes entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SubscriptionLineTypesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*SubscriptionLineTypes, error) {
	var entities []*SubscriptionLineTypes
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/SubscriptionLineTypes", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
