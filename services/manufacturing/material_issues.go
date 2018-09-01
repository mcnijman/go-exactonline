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

// MaterialIssuesEndpoint is responsible for communicating with
// the MaterialIssues endpoint of the Manufacturing service.
type MaterialIssuesEndpoint service

// MaterialIssues:
// Service: Manufacturing
// Entity: MaterialIssues
// URL: /api/v1/{division}/manufacturing/MaterialIssues
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ManufacturingMaterialIssues
type MaterialIssues struct {
	// StockTransactionId: ID of stock transaction related to this material issue
	StockTransactionId *types.GUID `json:"StockTransactionId,omitempty"`

	// CreatedBy: ID of creating user
	CreatedBy *types.GUID `json:"CreatedBy,omitempty"`

	// CreatedByFullName: Name of the creating user
	CreatedByFullName *string `json:"CreatedByFullName,omitempty"`

	// CreatedDate: Date this material issue was created
	CreatedDate *types.Date `json:"CreatedDate,omitempty"`

	// DraftStockTransactionID: Serial or batch numbers are reserved prior to a POST to MaterialIssues. This DraftStockTransactionID represents the group of serial or batch numbers to be used in this transaction.
	DraftStockTransactionID *types.GUID `json:"DraftStockTransactionID,omitempty"`

	// HasReversibleQuantity: Indicates if this MaterialIssue has a quantity eligible to be reversed via MaterialReversals
	HasReversibleQuantity *bool `json:"HasReversibleQuantity,omitempty"`

	// IsBackflush: Boolean indicating if this material issue was the result of shop order backflushing
	IsBackflush *byte `json:"IsBackflush,omitempty"`

	// IsBatch: Does the material issue&#39;s item use batch numbers
	IsBatch *byte `json:"IsBatch,omitempty"`

	// IsFractionAllowedItem: Indicates if fractions (for example 0.35) are allowed for quantities of the material issue&#39;s item
	IsFractionAllowedItem *byte `json:"IsFractionAllowedItem,omitempty"`

	// IsIssueFromChild: Boolean indicating if this material issue was an issue to a parent shop order
	IsIssueFromChild *byte `json:"IsIssueFromChild,omitempty"`

	// IsSerial: Does the material issue&#39;s item use serial numbers
	IsSerial *byte `json:"IsSerial,omitempty"`

	// Item: Item issued
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode: Code of item issued
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription: Description of item issued
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// ItemPictureUrl: Picture url of item issued
	ItemPictureUrl *string `json:"ItemPictureUrl,omitempty"`

	// Note: Notes logged with this material issue
	Note *string `json:"Note,omitempty"`

	// Quantity: Quantity of this material issue
	Quantity *float64 `json:"Quantity,omitempty"`

	// RelatedStockTransaction: If this transaction was part of a SubOrderReceipt, this ID is the related ShopOrderReceipt.StockTransactionID.
	RelatedStockTransaction *types.GUID `json:"RelatedStockTransaction,omitempty"`

	// ShopOrder: ID of shop order issued to
	ShopOrder *types.GUID `json:"ShopOrder,omitempty"`

	// ShopOrderMaterialPlan: ID of shop order material plan
	ShopOrderMaterialPlan *types.GUID `json:"ShopOrderMaterialPlan,omitempty"`

	// ShopOrderNumber: Number of shop order issued to
	ShopOrderNumber *int `json:"ShopOrderNumber,omitempty"`

	// StorageLocation: ID of storage location issued from
	StorageLocation *types.GUID `json:"StorageLocation,omitempty"`

	// StorageLocationCode: Code of storage location issued from
	StorageLocationCode *string `json:"StorageLocationCode,omitempty"`

	// StorageLocationDescription: Description of storage location issued from
	StorageLocationDescription *string `json:"StorageLocationDescription,omitempty"`

	// TransactionDate: Effective date of this material issue
	TransactionDate *types.Date `json:"TransactionDate,omitempty"`

	// Unit: Unit of measurement abbreviation of item issued
	Unit *string `json:"Unit,omitempty"`

	// UnitDescription: Unit of measurement of item issued
	UnitDescription *string `json:"UnitDescription,omitempty"`

	// Warehouse: ID of warehouse issued from
	Warehouse *types.GUID `json:"Warehouse,omitempty"`

	// WarehouseCode: Code of warehouse issued from
	WarehouseCode *string `json:"WarehouseCode,omitempty"`

	// WarehouseDescription: Description of warehouse issued from
	WarehouseDescription *string `json:"WarehouseDescription,omitempty"`
}

func (s *MaterialIssues) GetIdentifier() types.GUID {
	return *s.StockTransactionId
}

// List the MaterialIssues entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *MaterialIssuesEndpoint) List(ctx context.Context, division int, all bool) ([]*MaterialIssues, error) {
	var entities []*MaterialIssues
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/MaterialIssues?$select=*", division)
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