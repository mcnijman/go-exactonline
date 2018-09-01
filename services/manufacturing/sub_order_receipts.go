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

// SubOrderReceiptsEndpoint is responsible for communicating with
// the SubOrderReceipts endpoint of the Manufacturing service.
type SubOrderReceiptsEndpoint service

// SubOrderReceipts:
// Service: Manufacturing
// Entity: SubOrderReceipts
// URL: /api/v1/{division}/manufacturing/SubOrderReceipts
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ManufacturingSubOrderReceipts
type SubOrderReceipts struct {
	// ShopOrderReceiptStockTransactionId: ShopOrderReceipt.StockTransactionId related to this SubOrderReceipt
	ShopOrderReceiptStockTransactionId *types.GUID `json:"ShopOrderReceiptStockTransactionId,omitempty"`

	// CreatedBy: ID of creating user
	CreatedBy *types.GUID `json:"CreatedBy,omitempty"`

	// CreatedByFullName: Name of the creating user
	CreatedByFullName *string `json:"CreatedByFullName,omitempty"`

	// CreatedDate: Creation date of this SubOrderReceipt
	CreatedDate *types.Date `json:"CreatedDate,omitempty"`

	// DraftStockTransactionID: Serial or batch numbers are reserved prior to a POST to SubOrderReceipt. This DraftStockTransactionID represents the group of serial or batch numbers to be used in this transaction.
	DraftStockTransactionID *types.GUID `json:"DraftStockTransactionID,omitempty"`

	// HasReversibleQuantity: Indicates if this SubOrderReceipt has a quantity eligible to be reversed via SubOrderReversals
	HasReversibleQuantity *bool `json:"HasReversibleQuantity,omitempty"`

	// IsBatch: Does the SubOrderReceipt&#39;s item use batch numbers
	IsBatch *byte `json:"IsBatch,omitempty"`

	// IsFractionAllowedItem: Indicates if fractions (for example 0.35) are allowed for quantities of the SubOrderReceipt&#39;s item
	IsFractionAllowedItem *byte `json:"IsFractionAllowedItem,omitempty"`

	// IsSerial: Does the SubOrderReceipt&#39;s item use serial numbers
	IsSerial *byte `json:"IsSerial,omitempty"`

	// Item: Item of this SubOrderReceipt
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode: Code of this SubOrderReceipt&#39;s item
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription: Description of this SubOrderReceipt&#39;s item
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// ItemPictureUrl: Picture url of shop order item
	ItemPictureUrl *string `json:"ItemPictureUrl,omitempty"`

	// MaterialIssueStockTransactionId: MaterialIssue.StockTransactionId related to this SubOrderReceipt
	MaterialIssueStockTransactionId *types.GUID `json:"MaterialIssueStockTransactionId,omitempty"`

	// ParentShopOrder: Shop order issued to
	ParentShopOrder *types.GUID `json:"ParentShopOrder,omitempty"`

	// ParentShopOrderMaterialPlan: Shop order material plan issued to
	ParentShopOrderMaterialPlan *types.GUID `json:"ParentShopOrderMaterialPlan,omitempty"`

	// ParentShopOrderNumber: Number of shop order issued to
	ParentShopOrderNumber *int `json:"ParentShopOrderNumber,omitempty"`

	// Quantity: Quantity of this SubOrderReceipt
	Quantity *float64 `json:"Quantity,omitempty"`

	// SubShopOrder: Shop order issued from
	SubShopOrder *types.GUID `json:"SubShopOrder,omitempty"`

	// SubShopOrderNumber: Number of shop order issued from
	SubShopOrderNumber *int `json:"SubShopOrderNumber,omitempty"`

	// TransactionDate: Effective date of this SubOrderReceipt
	TransactionDate *types.Date `json:"TransactionDate,omitempty"`

	// Unit: Unit of measurement abbreviation of this SubOrderReceipt&#39;s item
	Unit *string `json:"Unit,omitempty"`

	// UnitDescription: Unit of measurement of this SubOrderReceipt&#39;s item
	UnitDescription *string `json:"UnitDescription,omitempty"`

	// Warehouse: ID of warehouse SubOrderReceipt
	Warehouse *types.GUID `json:"Warehouse,omitempty"`

	// WarehouseCode: Code of warehouse SubOrderReceipt
	WarehouseCode *string `json:"WarehouseCode,omitempty"`

	// WarehouseDescription: Description of warehouse SubOrderReceipt
	WarehouseDescription *string `json:"WarehouseDescription,omitempty"`
}

func (s *SubOrderReceipts) GetIdentifier() types.GUID {
	return *s.ShopOrderReceiptStockTransactionId
}

// List the SubOrderReceipts entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SubOrderReceiptsEndpoint) List(ctx context.Context, division int, all bool) ([]*SubOrderReceipts, error) {
	var entities []*SubOrderReceipts
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/SubOrderReceipts?$select=*", division)
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