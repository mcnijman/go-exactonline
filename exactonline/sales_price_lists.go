// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// SalesPriceListsService is responsible for communicating with
// the PriceLists endpoint of the Sales service.
type SalesPriceListsService service

// SalesPriceLists:
// Service: Sales
// Entity: PriceLists
// URL: /api/v1/{division}/sales/PriceLists
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SalesPriceLists
type SalesPriceLists struct {
	// ID: Primary key
	ID *GUID `json:",omitempty"`

	// Code: Code to indicate the price list
	Code *string `json:",omitempty"`

	// Created: Creation date
	Created *Date `json:",omitempty"`

	// Creator: User ID of creator
	Creator *GUID `json:",omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:",omitempty"`

	// Currency: All prices in the price list are stored in this currency
	Currency *string `json:",omitempty"`

	// Description: Description
	Description *string `json:",omitempty"`

	// Division: Division code
	Division *int `json:",omitempty"`

	// DivisionDescription: Description of Division
	DivisionDescription *string `json:",omitempty"`

	// Entity: Indicates the entity (Item, Item group, ..) on which this price list is based
	Entity *int `json:",omitempty"`

	// Modified: Last modified date
	Modified *Date `json:",omitempty"`

	// Modifier: User ID of modifier
	Modifier *GUID `json:",omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:",omitempty"`

	// Notes: Explanation or extra information can be stored in the notes
	Notes *string `json:",omitempty"`
}

func (s *SalesPriceLists) GetIdentifier() GUID {
	return *s.ID
}

// List the PriceLists entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SalesPriceListsService) List(ctx context.Context, division int, all bool) ([]*SalesPriceLists, error) {
	var entities []*SalesPriceLists
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/sales/PriceLists?$select=*", division)
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