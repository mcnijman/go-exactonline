// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// ManufacturingMaterialIssuesService is responsible for communicating with
// the MaterialIssues endpoint of the Manufacturing service.
type ManufacturingMaterialIssuesService service

// ManufacturingMaterialIssues:
// Service: Manufacturing
// Entity: MaterialIssues
// URL: /api/v1/{division}/manufacturing/MaterialIssues
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ManufacturingMaterialIssues
type ManufacturingMaterialIssues struct {
	// StockTransactionId: ID of stock transaction related to this material issue
	StockTransactionId *GUID `json:",omitempty"`

	// CreatedBy: ID of creating user
	CreatedBy *GUID `json:",omitempty"`

	// CreatedByFullName: Name of the creating user
	CreatedByFullName *string `json:",omitempty"`

	// CreatedDate: Date this material issue was created
	CreatedDate *Date `json:",omitempty"`

	// DraftStockTransactionID: Serial or batch numbers are reserved prior to a POST to MaterialIssues. This DraftStockTransactionID represents the group of serial or batch numbers to be used in this transaction.
	DraftStockTransactionID *GUID `json:",omitempty"`

	// HasReversibleQuantity: Indicates if this MaterialIssue has a quantity eligible to be reversed via MaterialReversals
	HasReversibleQuantity *bool `json:",omitempty"`

	// IsBackflush: Boolean indicating if this material issue was the result of shop order backflushing
	IsBackflush *byte `json:",omitempty"`

	// IsBatch: Does the material issue&#39;s item use batch numbers
	IsBatch *byte `json:",omitempty"`

	// IsFractionAllowedItem: Indicates if fractions (for example 0.35) are allowed for quantities of the material issue&#39;s item
	IsFractionAllowedItem *byte `json:",omitempty"`

	// IsIssueFromChild: Boolean indicating if this material issue was an issue to a parent shop order
	IsIssueFromChild *byte `json:",omitempty"`

	// IsSerial: Does the material issue&#39;s item use serial numbers
	IsSerial *byte `json:",omitempty"`

	// Item: Item issued
	Item *GUID `json:",omitempty"`

	// ItemCode: Code of item issued
	ItemCode *string `json:",omitempty"`

	// ItemDescription: Description of item issued
	ItemDescription *string `json:",omitempty"`

	// ItemPictureUrl: Picture url of item issued
	ItemPictureUrl *string `json:",omitempty"`

	// Note: Notes logged with this material issue
	Note *string `json:",omitempty"`

	// Quantity: Quantity of this material issue
	Quantity *float64 `json:",omitempty"`

	// RelatedStockTransaction: If this transaction was part of a SubOrderReceipt, this ID is the related ShopOrderReceipt.StockTransactionID.
	RelatedStockTransaction *GUID `json:",omitempty"`

	// ShopOrder: ID of shop order issued to
	ShopOrder *GUID `json:",omitempty"`

	// ShopOrderMaterialPlan: ID of shop order material plan
	ShopOrderMaterialPlan *GUID `json:",omitempty"`

	// ShopOrderNumber: Number of shop order issued to
	ShopOrderNumber *int `json:",omitempty"`

	// StorageLocation: ID of storage location issued from
	StorageLocation *GUID `json:",omitempty"`

	// StorageLocationCode: Code of storage location issued from
	StorageLocationCode *string `json:",omitempty"`

	// StorageLocationDescription: Description of storage location issued from
	StorageLocationDescription *string `json:",omitempty"`

	// TransactionDate: Effective date of this material issue
	TransactionDate *Date `json:",omitempty"`

	// Unit: Unit of measurement abbreviation of item issued
	Unit *string `json:",omitempty"`

	// UnitDescription: Unit of measurement of item issued
	UnitDescription *string `json:",omitempty"`

	// Warehouse: ID of warehouse issued from
	Warehouse *GUID `json:",omitempty"`

	// WarehouseCode: Code of warehouse issued from
	WarehouseCode *string `json:",omitempty"`

	// WarehouseDescription: Description of warehouse issued from
	WarehouseDescription *string `json:",omitempty"`
}

func (s *ManufacturingMaterialIssues) GetIdentifier() GUID {
	return *s.StockTransactionId
}

// List the MaterialIssues entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ManufacturingMaterialIssuesService) List(ctx context.Context, division int, all bool) ([]*ManufacturingMaterialIssues, error) {
	var entities []*ManufacturingMaterialIssues
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