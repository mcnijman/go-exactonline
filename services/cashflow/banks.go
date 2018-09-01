// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package cashflow

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// BanksEndpoint is responsible for communicating with
// the Banks endpoint of the Cashflow service.
type BanksEndpoint service

// Banks:
// Service: Cashflow
// Entity: Banks
// URL: /api/v1/{division}/cashflow/Banks
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=CashflowBanks
type Banks struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// BankName: The name of the bank
	BankName *string `json:"BankName,omitempty"`

	// BICCode: The bank identification code of the bank
	BICCode *string `json:"BICCode,omitempty"`

	// Country: The country in which the bank is based
	Country *string `json:"Country,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Description: The extended description of the bank
	Description *string `json:"Description,omitempty"`

	// Format: The account format used by the bank
	Format *string `json:"Format,omitempty"`

	// HomePageAddress: The website of the bank
	HomePageAddress *string `json:"HomePageAddress,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Status: The status of the bank. A = Active, P = Passive
	Status *string `json:"Status,omitempty"`
}

func (s *Banks) GetIdentifier() types.GUID {
	return *s.ID
}

// List the Banks entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *BanksEndpoint) List(ctx context.Context, division int, all bool) ([]*Banks, error) {
	var entities []*Banks
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/cashflow/Banks?$select=*", division)
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
