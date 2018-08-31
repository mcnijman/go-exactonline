// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// InventoryStockBatchNumbersService is responsible for communicating with
// the StockBatchNumbers endpoint of the Inventory service.
type InventoryStockBatchNumbersService service

// InventoryStockBatchNumbers:
// Service: Inventory
// Entity: StockBatchNumbers
// URL: /api/v1/{division}/inventory/StockBatchNumbers
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=InventoryStockBatchNumbers
type InventoryStockBatchNumbers struct {
	// ID: Primary key
	ID *GUID `json:",omitempty"`

	// BatchNumber: Human readable batch number
	BatchNumber *string `json:",omitempty"`

	// BatchNumberID: Batch number ID
	BatchNumberID *GUID `json:",omitempty"`

	// Created: Creation date
	Created *Date `json:",omitempty"`

	// Creator: User ID of creator
	Creator *GUID `json:",omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:",omitempty"`

	// Division: Division code
	Division *int `json:",omitempty"`

	// DraftStockTransactionID: ID representing a group of batch numbers being reserved for use in a subsequent stock transaction
	DraftStockTransactionID *GUID `json:",omitempty"`

	// EndDate: End date of effective period for batch number
	EndDate *Date `json:",omitempty"`

	// IsBlocked: Boolean value indicating whether or not the batch number is blocked
	IsBlocked *byte `json:",omitempty"`

	// IsDraft: Boolean value indicating if this batch number is being reserved
	IsDraft *byte `json:",omitempty"`

	// Item: Item
	Item *GUID `json:",omitempty"`

	// ItemCode: Item code
	ItemCode *string `json:",omitempty"`

	// ItemDescription: Description of item
	ItemDescription *string `json:",omitempty"`

	// Modified: Last modified date
	Modified *Date `json:",omitempty"`

	// Modifier: User ID of modifier
	Modifier *GUID `json:",omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:",omitempty"`

	// Quantity: Quantity of this batch number entering or leaving inventory
	Quantity *float64 `json:",omitempty"`

	// Remarks: Remarks
	Remarks *string `json:",omitempty"`

	// StockCountLine: ID of stock count entry
	StockCountLine *GUID `json:",omitempty"`

	// StockTransactionID: ID of the stock transaction in which this batch number was used
	StockTransactionID *GUID `json:",omitempty"`

	// StockTransactionType: Type of stock transaction associated with this batch number.Available values:10 = Opening balance120 = Goods delivery121 = Sales return122 = Stock out (Drop shipment)123 = Stock in (Drop shipment return)124 = Warehouse transfer delivery125 = Location Transfer Delivery130 = Goods receipt131 = Purchase return132 = Stock in (Drop shipment)133 = Stock out (Drop shipment return)134 = Warehouse transfer receipt135 = Location Transfer Receipt140 = Shop order stock receipt141 = Shop order stock reversal147 = Shop order by-product receipt148 = Shop order by-product reversal150 = Requirement issue151 = Requirement reversal155 = Subcontract issue156 = Subcontract return160 = Receipt (Assembly)161 = Return receipt (Disassembly)165 = Issue (Assembly)166 = Return issue (Disassembly)180 = Stock revaluation181 = Financial revaluation195 = Stock count196 = Adjust stock - out197 = Adjust stock - in
	StockTransactionType *int `json:",omitempty"`

	// StorageLocation: Storage location which this batch number is entering or leaving
	StorageLocation *GUID `json:",omitempty"`

	// StorageLocationCode: Code of the storage location which this batch number is entering or leaving
	StorageLocationCode *string `json:",omitempty"`

	// StorageLocationDescription: Description of the storage location which this batch number is entering or leaving
	StorageLocationDescription *string `json:",omitempty"`

	// Warehouse: Warehouse which this batch number is entering or leaving
	Warehouse *GUID `json:",omitempty"`

	// WarehouseCode: Code of the warehouse which this batch number is entering or leaving
	WarehouseCode *string `json:",omitempty"`

	// WarehouseDescription: Description of the warehouse which this batch number is entering or leaving
	WarehouseDescription *string `json:",omitempty"`
}

func (s *InventoryStockBatchNumbers) GetIdentifier() GUID {
	return *s.ID
}

// List the StockBatchNumbers entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *InventoryStockBatchNumbersService) List(ctx context.Context, division int, all bool) ([]*InventoryStockBatchNumbers, error) {
	var entities []*InventoryStockBatchNumbers
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/inventory/StockBatchNumbers?$select=*", division)
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