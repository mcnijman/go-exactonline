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

// TimeAndBillingEntryAccountsEndpoint is responsible for communicating with
// the TimeAndBillingEntryAccounts endpoint of the Project service.
type TimeAndBillingEntryAccountsEndpoint service

// TimeAndBillingEntryAccounts:
// Service: Project
// Entity: TimeAndBillingEntryAccounts
// URL: /api/v1/{division}/read/project/TimeAndBillingEntryAccounts
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadProjectTimeAndBillingEntryAccounts
type TimeAndBillingEntryAccounts struct {
	// AccountId: Primary key
	AccountId *types.GUID `json:"AccountId,omitempty"`

	// AccountName: Name of account
	AccountName *string `json:"AccountName,omitempty"`
}

func (s *TimeAndBillingEntryAccounts) GetIdentifier() types.GUID {
	return *s.AccountId
}

// List the TimeAndBillingEntryAccounts entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *TimeAndBillingEntryAccountsEndpoint) List(ctx context.Context, division int, all bool) ([]*TimeAndBillingEntryAccounts, error) {
	var entities []*TimeAndBillingEntryAccounts
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/TimeAndBillingEntryAccounts?$select=*", division)
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
