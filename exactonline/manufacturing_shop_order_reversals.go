// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// ManufacturingShopOrderReversalsService is responsible for communicating with
// the ShopOrderReversals endpoint of the Manufacturing service.
type ManufacturingShopOrderReversalsService service

// ManufacturingShopOrderReversals:
// Service: Manufacturing
// Entity: ShopOrderReversals
// URL: /api/v1/{division}/manufacturing/ShopOrderReversals
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ManufacturingShopOrderReversals
type ManufacturingShopOrderReversals struct {
	// ReversalStockTransactionId: ID of stock transaction related to this ShopOrderReversal
	ReversalStockTransactionId *GUID `json:",omitempty"`

	// CreatedBy: ID of creating user
	CreatedBy *GUID `json:",omitempty"`

	// CreatedByFullName: Name of the creating user
	CreatedByFullName *string `json:",omitempty"`

	// CreatedDate: Date of this reversal
	CreatedDate *Date `json:",omitempty"`

	// IsBatch: Does the ShopOrderReversal&#39;s item use batch numbers
	IsBatch *byte `json:",omitempty"`

	// IsFractionAllowedItem: Indicates if fractions (for example 0.35) are allowed for quantities of the ShopOrderReversal&#39;s item
	IsFractionAllowedItem *byte `json:",omitempty"`

	// IsSerial: Does the ShopOrderReversal&#39;s item use serial numbers
	IsSerial *byte `json:",omitempty"`

	// Item: Item reversed
	Item *GUID `json:",omitempty"`

	// ItemCode: Code of item reversed
	ItemCode *string `json:",omitempty"`

	// ItemDescription: Description of item reversed
	ItemDescription *string `json:",omitempty"`

	// ItemPictureUrl: Picture url of shop order item
	ItemPictureUrl *string `json:",omitempty"`

	// Note: Notes associated with this reversal
	Note *string `json:",omitempty"`

	// OriginalStockTransactionId: ID of the original stock transaction, which was reversed
	OriginalStockTransactionId *GUID `json:",omitempty"`

	// Quantity: Quantity reversed
	Quantity *float64 `json:",omitempty"`

	// ShopOrder: Shop order being reversed
	ShopOrder *GUID `json:",omitempty"`

	// ShopOrderNumber: Number of shop order being reversed
	ShopOrderNumber *int `json:",omitempty"`

	// StorageLocation: ID of storage location reversed from
	StorageLocation *GUID `json:",omitempty"`

	// StorageLocationCode: Code of storage location reversed from
	StorageLocationCode *string `json:",omitempty"`

	// StorageLocationDescription: Description of storage location reversed from
	StorageLocationDescription *string `json:",omitempty"`

	// TransactionDate: Effective date of this ShopOrderReversal
	TransactionDate *Date `json:",omitempty"`

	// Unit: Unit of measurement abbreviation of item reversed
	Unit *string `json:",omitempty"`

	// UnitDescription: Unit of measurement of item reversed
	UnitDescription *string `json:",omitempty"`

	// Warehouse: ID of warehouse reversed from
	Warehouse *GUID `json:",omitempty"`

	// WarehouseCode: Code of warehouse reversed from
	WarehouseCode *string `json:",omitempty"`

	// WarehouseDescription: Description of warehouse reversed from
	WarehouseDescription *string `json:",omitempty"`
}

func (s *ManufacturingShopOrderReversals) GetIdentifier() GUID {
	return *s.ReversalStockTransactionId
}

// List the ShopOrderReversals entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ManufacturingShopOrderReversalsService) List(ctx context.Context, division int, all bool) ([]*ManufacturingShopOrderReversals, error) {
	var entities []*ManufacturingShopOrderReversals
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrderReversals?$select=*", division)
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