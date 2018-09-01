// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package logistics

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// ItemGroupsEndpoint is responsible for communicating with
// the ItemGroups endpoint of the Logistics service.
type ItemGroupsEndpoint service

// ItemGroups:
// Service: Logistics
// Entity: ItemGroups
// URL: /api/v1/{division}/logistics/ItemGroups
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=LogisticsItemGroups
type ItemGroups struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Code: Code of the item group
	Code *string `json:"Code,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description of the item group
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// GLCosts: GL account on which the costs of items of this group will be booked
	GLCosts *types.GUID `json:"GLCosts,omitempty"`

	// GLCostsCode: Code of GLCosts
	GLCostsCode *string `json:"GLCostsCode,omitempty"`

	// GLCostsDescription: Description of GLCosts
	GLCostsDescription *string `json:"GLCostsDescription,omitempty"`

	// GLPurchaseAccount: GL Purchase account for purchase invoicing according to (non-) perpetual inventory method
	GLPurchaseAccount *types.GUID `json:"GLPurchaseAccount,omitempty"`

	// GLPurchaseAccountCode: Code of GLPurchase
	GLPurchaseAccountCode *string `json:"GLPurchaseAccountCode,omitempty"`

	// GLPurchaseAccountDescription: Description of GLPurchaseAccount
	GLPurchaseAccountDescription *string `json:"GLPurchaseAccountDescription,omitempty"`

	// GLPurchasePriceDifference: GL account that will be used for the &#39;Standard cost price&#39; valuation method to balance the difference between purchase price and cost price
	GLPurchasePriceDifference *types.GUID `json:"GLPurchasePriceDifference,omitempty"`

	// GLPurchasePriceDifferenceCode: Code of GLPurchasePriceDifference
	GLPurchasePriceDifferenceCode *string `json:"GLPurchasePriceDifferenceCode,omitempty"`

	// GLPurchasePriceDifferenceDescr: Description of GLPurchasePriceDifference
	GLPurchasePriceDifferenceDescr *string `json:"GLPurchasePriceDifferenceDescr,omitempty"`

	// GLRevenue: GL account on which the revenue for items of this group will be booked
	GLRevenue *types.GUID `json:"GLRevenue,omitempty"`

	// GLRevenueCode: Code of GLRevenue
	GLRevenueCode *string `json:"GLRevenueCode,omitempty"`

	// GLRevenueDescription: Description of GLRevenue
	GLRevenueDescription *string `json:"GLRevenueDescription,omitempty"`

	// GLStock: GL account on which stock entries will be booked for items of this group
	GLStock *types.GUID `json:"GLStock,omitempty"`

	// GLStockCode: Code of GLStock
	GLStockCode *string `json:"GLStockCode,omitempty"`

	// GLStockDescription: Description of GLStock
	GLStockDescription *string `json:"GLStockDescription,omitempty"`

	// GLStockVariance: GL stock variance account for perpetual inventory
	GLStockVariance *types.GUID `json:"GLStockVariance,omitempty"`

	// GLStockVarianceCode: Code of GLStockVariance
	GLStockVarianceCode *string `json:"GLStockVarianceCode,omitempty"`

	// GLStockVarianceDescription: Description of GLStockVariance
	GLStockVarianceDescription *string `json:"GLStockVarianceDescription,omitempty"`

	// IsDefault: Indicates if this is the default item group that will be assigned when a new item is created
	IsDefault *byte `json:"IsDefault,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Notes
	Notes *string `json:"Notes,omitempty"`
}

func (s *ItemGroups) GetIdentifier() types.GUID {
	return *s.ID
}

// List the ItemGroups entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ItemGroupsEndpoint) List(ctx context.Context, division int, all bool) ([]*ItemGroups, error) {
	var entities []*ItemGroups
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/ItemGroups?$select=*", division)
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
