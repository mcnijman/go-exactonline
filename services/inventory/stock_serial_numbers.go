// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package inventory

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// StockSerialNumbersEndpoint is responsible for communicating with
// the StockSerialNumbers endpoint of the Inventory service.
type StockSerialNumbersEndpoint service

// StockSerialNumbers:
// Service: Inventory
// Entity: StockSerialNumbers
// URL: /api/v1/{division}/inventory/StockSerialNumbers
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=InventoryStockSerialNumbers
type StockSerialNumbers struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// DraftStockTransactionID: ID representing a group of serial numbers being reserved for use in a subsequent stock transaction
	DraftStockTransactionID *types.GUID `json:"DraftStockTransactionID,omitempty"`

	// EndDate: End date of effective period for serial number
	EndDate *types.Date `json:"EndDate,omitempty"`

	// IsBlocked: Boolean value indicating whether or not the serial number is blocked
	IsBlocked *byte `json:"IsBlocked,omitempty"`

	// IsDraft: Boolean value indicating if this serial number is being reserved
	IsDraft *byte `json:"IsDraft,omitempty"`

	// Item: Item
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode: Item code
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription: Description of item
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Remarks: Remarks
	Remarks *string `json:"Remarks,omitempty"`

	// SerialNumber: Human readable serial number
	SerialNumber *string `json:"SerialNumber,omitempty"`

	// SerialNumberID: Serial number ID
	SerialNumberID *types.GUID `json:"SerialNumberID,omitempty"`

	// StartDate: Start date of effective period for serial number
	StartDate *types.Date `json:"StartDate,omitempty"`

	// StockCountLine: ID of stock count entry
	StockCountLine *types.GUID `json:"StockCountLine,omitempty"`

	// StockTransactionID: ID of the stock transaction in which this serial number was used
	StockTransactionID *types.GUID `json:"StockTransactionID,omitempty"`

	// StockTransactionType: Type of stock transaction associated with this serial number.Available values:10 = Opening balance120 = Goods delivery121 = Sales return122 = Stock out (Drop shipment)123 = Stock in (Drop shipment return)124 = Warehouse transfer delivery125 = Location Transfer Delivery130 = Goods receipt131 = Purchase return132 = Stock in (Drop shipment)133 = Stock out (Drop shipment return)134 = Warehouse transfer receipt135 = Location Transfer Receipt140 = Shop order stock receipt141 = Shop order stock reversal147 = Shop order by-product receipt148 = Shop order by-product reversal150 = Requirement issue151 = Requirement reversal155 = Subcontract issue156 = Subcontract return160 = Receipt (Assembly)161 = Return receipt (Disassembly)165 = Issue (Assembly)166 = Return issue (Disassembly)180 = Stock revaluation181 = Financial revaluation195 = Stock count196 = Adjust stock - out197 = Adjust stock - in
	StockTransactionType *int `json:"StockTransactionType,omitempty"`

	// StorageLocation: Storage location which this serial number is entering or leaving
	StorageLocation *types.GUID `json:"StorageLocation,omitempty"`

	// StorageLocationCode: Code of the storage location which this serial number is entering or leaving
	StorageLocationCode *string `json:"StorageLocationCode,omitempty"`

	// StorageLocationDescription: Description of the storage location which this serial number is entering or leaving
	StorageLocationDescription *string `json:"StorageLocationDescription,omitempty"`

	// Warehouse: Warehouse which this serial number is entering or leaving
	Warehouse *types.GUID `json:"Warehouse,omitempty"`

	// WarehouseCode: Code of the warehouse which this serial number is entering or leaving
	WarehouseCode *string `json:"WarehouseCode,omitempty"`

	// WarehouseDescription: Description of the warehouse which this serial number is entering or leaving
	WarehouseDescription *string `json:"WarehouseDescription,omitempty"`
}

func (s *StockSerialNumbers) GetIdentifier() types.GUID {
	return *s.ID
}

// List the StockSerialNumbers entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *StockSerialNumbersEndpoint) List(ctx context.Context, division int, all bool) ([]*StockSerialNumbers, error) {
	var entities []*StockSerialNumbers
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/inventory/StockSerialNumbers?$select=*", division)
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