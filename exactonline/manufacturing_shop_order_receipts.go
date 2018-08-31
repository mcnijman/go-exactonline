// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// ManufacturingShopOrderReceiptsService is responsible for communicating with
// the ShopOrderReceipts endpoint of the Manufacturing service.
type ManufacturingShopOrderReceiptsService service

// ManufacturingShopOrderReceipts:
// Service: Manufacturing
// Entity: ShopOrderReceipts
// URL: /api/v1/{division}/manufacturing/ShopOrderReceipts
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ManufacturingShopOrderReceipts
type ManufacturingShopOrderReceipts struct {
	// StockTransactionId: ID of stock transaction of this ShopOrderReceipt
	StockTransactionId *GUID `json:",omitempty"`

	// CreatedBy: ID of creating user
	CreatedBy *GUID `json:",omitempty"`

	// CreatedByFullName: Name of the creating user
	CreatedByFullName *string `json:",omitempty"`

	// CreatedDate: Date of this ShopOrderReceipt
	CreatedDate *Date `json:",omitempty"`

	// DraftStockTransactionID: Serial or batch numbers are reserved prior to a POST to ShopOrderReceipts. This DraftStockTransactionID represents the group of serial or batch numbers to be used in this transaction.
	DraftStockTransactionID *GUID `json:",omitempty"`

	// HasReversibleQuantity: Indicates if this ShopOrderReceipt has a quantity eligible to be reversed via ShopOrderReversals
	HasReversibleQuantity *bool `json:",omitempty"`

	// IsBatch: Does the shop order receipt&#39;s item use batch numbers
	IsBatch *byte `json:",omitempty"`

	// IsFractionAllowedItem: Indicates if fractions (for example 0.35) are allowed for quantities of the shop order receipt&#39;s item
	IsFractionAllowedItem *byte `json:",omitempty"`

	// IsIssueToParent: Boolean indicating if this ShopOrderReceipt was part of an SubOrderReceipt
	IsIssueToParent *bool `json:",omitempty"`

	// IsSerial: Does the shop order receipt&#39;s item use serial numbers
	IsSerial *byte `json:",omitempty"`

	// Item: Item finished
	Item *GUID `json:",omitempty"`

	// ItemCode: Code of item finished
	ItemCode *string `json:",omitempty"`

	// ItemDescription: Description of item finished
	ItemDescription *string `json:",omitempty"`

	// ItemPictureUrl: Picture url of shop order item
	ItemPictureUrl *string `json:",omitempty"`

	// ParentShopOrder: Parent shop order if this ShopOrderReceipt is part of a SubOrderReceipt
	ParentShopOrder *GUID `json:",omitempty"`

	// ParentShopOrderNumber: Parent shop order number if this ShopOrderReceipt is part of a SubOrderReceipt
	ParentShopOrderNumber *int `json:",omitempty"`

	// Quantity: Quantity of this ShopOrderReceipt
	Quantity *float64 `json:",omitempty"`

	// RelatedStockTransaction: If this transaction was part of a SubOrderReceipt, this ID is the related MaterialIssue.StockTransactionID.
	RelatedStockTransaction *GUID `json:",omitempty"`

	// ShopOrder: Shop order finished
	ShopOrder *GUID `json:",omitempty"`

	// ShopOrderNumber: Number of shop order finished
	ShopOrderNumber *int `json:",omitempty"`

	// StorageLocation: ID of storage location finished to
	StorageLocation *GUID `json:",omitempty"`

	// StorageLocationCode: Code of storage location finished to
	StorageLocationCode *string `json:",omitempty"`

	// StorageLocationDescription: Description of storage location finished to
	StorageLocationDescription *string `json:",omitempty"`

	// TransactionDate: Effective date of this ShopOrderReceipt
	TransactionDate *Date `json:",omitempty"`

	// Unit: Unit of measurement abbreviation of item finished
	Unit *string `json:",omitempty"`

	// UnitDescription: Unit of measurement of item finished
	UnitDescription *string `json:",omitempty"`

	// Warehouse: ID of warehouse finished to
	Warehouse *GUID `json:",omitempty"`

	// WarehouseCode: Code of warehouse finished to
	WarehouseCode *string `json:",omitempty"`

	// WarehouseDescription: Description of warehouse finished to
	WarehouseDescription *string `json:",omitempty"`
}

func (s *ManufacturingShopOrderReceipts) GetIdentifier() GUID {
	return *s.StockTransactionId
}

// List the ShopOrderReceipts entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ManufacturingShopOrderReceiptsService) List(ctx context.Context, division int, all bool) ([]*ManufacturingShopOrderReceipts, error) {
	var entities []*ManufacturingShopOrderReceipts
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