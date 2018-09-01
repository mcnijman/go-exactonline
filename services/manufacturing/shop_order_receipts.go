// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package manufacturing

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// ShopOrderReceiptsEndpoint is responsible for communicating with
// the ShopOrderReceipts endpoint of the Manufacturing service.
type ShopOrderReceiptsEndpoint service

// ShopOrderReceipts:
// Service: Manufacturing
// Entity: ShopOrderReceipts
// URL: /api/v1/{division}/manufacturing/ShopOrderReceipts
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ManufacturingShopOrderReceipts
type ShopOrderReceipts struct {
	// StockTransactionId: ID of stock transaction of this ShopOrderReceipt
	StockTransactionId *types.GUID `json:"StockTransactionId,omitempty"`

	// CreatedBy: ID of creating user
	CreatedBy *types.GUID `json:"CreatedBy,omitempty"`

	// CreatedByFullName: Name of the creating user
	CreatedByFullName *string `json:"CreatedByFullName,omitempty"`

	// CreatedDate: Date of this ShopOrderReceipt
	CreatedDate *types.Date `json:"CreatedDate,omitempty"`

	// DraftStockTransactionID: Serial or batch numbers are reserved prior to a POST to ShopOrderReceipts. This DraftStockTransactionID represents the group of serial or batch numbers to be used in this transaction.
	DraftStockTransactionID *types.GUID `json:"DraftStockTransactionID,omitempty"`

	// HasReversibleQuantity: Indicates if this ShopOrderReceipt has a quantity eligible to be reversed via ShopOrderReversals
	HasReversibleQuantity *bool `json:"HasReversibleQuantity,omitempty"`

	// IsBatch: Does the shop order receipt&#39;s item use batch numbers
	IsBatch *byte `json:"IsBatch,omitempty"`

	// IsFractionAllowedItem: Indicates if fractions (for example 0.35) are allowed for quantities of the shop order receipt&#39;s item
	IsFractionAllowedItem *byte `json:"IsFractionAllowedItem,omitempty"`

	// IsIssueToParent: Boolean indicating if this ShopOrderReceipt was part of an SubOrderReceipt
	IsIssueToParent *bool `json:"IsIssueToParent,omitempty"`

	// IsSerial: Does the shop order receipt&#39;s item use serial numbers
	IsSerial *byte `json:"IsSerial,omitempty"`

	// Item: Item finished
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode: Code of item finished
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription: Description of item finished
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// ItemPictureUrl: Picture url of shop order item
	ItemPictureUrl *string `json:"ItemPictureUrl,omitempty"`

	// ParentShopOrder: Parent shop order if this ShopOrderReceipt is part of a SubOrderReceipt
	ParentShopOrder *types.GUID `json:"ParentShopOrder,omitempty"`

	// ParentShopOrderNumber: Parent shop order number if this ShopOrderReceipt is part of a SubOrderReceipt
	ParentShopOrderNumber *int `json:"ParentShopOrderNumber,omitempty"`

	// Quantity: Quantity of this ShopOrderReceipt
	Quantity *float64 `json:"Quantity,omitempty"`

	// RelatedStockTransaction: If this transaction was part of a SubOrderReceipt, this ID is the related MaterialIssue.StockTransactionID.
	RelatedStockTransaction *types.GUID `json:"RelatedStockTransaction,omitempty"`

	// ShopOrder: Shop order finished
	ShopOrder *types.GUID `json:"ShopOrder,omitempty"`

	// ShopOrderNumber: Number of shop order finished
	ShopOrderNumber *int `json:"ShopOrderNumber,omitempty"`

	// StorageLocation: ID of storage location finished to
	StorageLocation *types.GUID `json:"StorageLocation,omitempty"`

	// StorageLocationCode: Code of storage location finished to
	StorageLocationCode *string `json:"StorageLocationCode,omitempty"`

	// StorageLocationDescription: Description of storage location finished to
	StorageLocationDescription *string `json:"StorageLocationDescription,omitempty"`

	// TransactionDate: Effective date of this ShopOrderReceipt
	TransactionDate *types.Date `json:"TransactionDate,omitempty"`

	// Unit: Unit of measurement abbreviation of item finished
	Unit *string `json:"Unit,omitempty"`

	// UnitDescription: Unit of measurement of item finished
	UnitDescription *string `json:"UnitDescription,omitempty"`

	// Warehouse: ID of warehouse finished to
	Warehouse *types.GUID `json:"Warehouse,omitempty"`

	// WarehouseCode: Code of warehouse finished to
	WarehouseCode *string `json:"WarehouseCode,omitempty"`

	// WarehouseDescription: Description of warehouse finished to
	WarehouseDescription *string `json:"WarehouseDescription,omitempty"`
}

func (s *ShopOrderReceipts) GetIdentifier() types.GUID {
	return *s.StockTransactionId
}

// List the ShopOrderReceipts entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ShopOrderReceiptsEndpoint) List(ctx context.Context, division int, all bool) ([]*ShopOrderReceipts, error) {
	var entities []*ShopOrderReceipts
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrderReceipts?$select=*", division)
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
