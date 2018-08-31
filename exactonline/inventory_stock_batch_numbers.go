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
	ID *GUID `json:"ID,omitempty"`

	// BatchNumber: Human readable batch number
	BatchNumber *string `json:"BatchNumber,omitempty"`

	// BatchNumberID: Batch number ID
	BatchNumberID *GUID `json:"BatchNumberID,omitempty"`

	// Created: Creation date
	Created *Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// DraftStockTransactionID: ID representing a group of batch numbers being reserved for use in a subsequent stock transaction
	DraftStockTransactionID *GUID `json:"DraftStockTransactionID,omitempty"`

	// EndDate: End date of effective period for batch number
	EndDate *Date `json:"EndDate,omitempty"`

	// IsBlocked: Boolean value indicating whether or not the batch number is blocked
	IsBlocked *byte `json:"IsBlocked,omitempty"`

	// IsDraft: Boolean value indicating if this batch number is being reserved
	IsDraft *byte `json:"IsDraft,omitempty"`

	// Item: Item
	Item *GUID `json:"Item,omitempty"`

	// ItemCode: Item code
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription: Description of item
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// Modified: Last modified date
	Modified *Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Quantity: Quantity of this batch number entering or leaving inventory
	Quantity *float64 `json:"Quantity,omitempty"`

	// Remarks: Remarks
	Remarks *string `json:"Remarks,omitempty"`

	// StockCountLine: ID of stock count entry
	StockCountLine *GUID `json:"StockCountLine,omitempty"`

	// StockTransactionID: ID of the stock transaction in which this batch number was used
	StockTransactionID *GUID `json:"StockTransactionID,omitempty"`

	// StockTransactionType: Type of stock transaction associated with this batch number.Available values:10 = Opening balance120 = Goods delivery121 = Sales return122 = Stock out (Drop shipment)123 = Stock in (Drop shipment return)124 = Warehouse transfer delivery125 = Location Transfer Delivery130 = Goods receipt131 = Purchase return132 = Stock in (Drop shipment)133 = Stock out (Drop shipment return)134 = Warehouse transfer receipt135 = Location Transfer Receipt140 = Shop order stock receipt141 = Shop order stock reversal147 = Shop order by-product receipt148 = Shop order by-product reversal150 = Requirement issue151 = Requirement reversal155 = Subcontract issue156 = Subcontract return160 = Receipt (Assembly)161 = Return receipt (Disassembly)165 = Issue (Assembly)166 = Return issue (Disassembly)180 = Stock revaluation181 = Financial revaluation195 = Stock count196 = Adjust stock - out197 = Adjust stock - in
	StockTransactionType *int `json:"StockTransactionType,omitempty"`

	// StorageLocation: Storage location which this batch number is entering or leaving
	StorageLocation *GUID `json:"StorageLocation,omitempty"`

	// StorageLocationCode: Code of the storage location which this batch number is entering or leaving
	StorageLocationCode *string `json:"StorageLocationCode,omitempty"`

	// StorageLocationDescription: Description of the storage location which this batch number is entering or leaving
	StorageLocationDescription *string `json:"StorageLocationDescription,omitempty"`

	// Warehouse: Warehouse which this batch number is entering or leaving
	Warehouse *GUID `json:"Warehouse,omitempty"`

	// WarehouseCode: Code of the warehouse which this batch number is entering or leaving
	WarehouseCode *string `json:"WarehouseCode,omitempty"`

	// WarehouseDescription: Description of the warehouse which this batch number is entering or leaving
	WarehouseDescription *string `json:"WarehouseDescription,omitempty"`
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

/* // Get the StockBatchNumbers enitity, by ID.
func (s *InventoryStockBatchNumbersService) Get(ctx context.Context, division int, id GUID) (*InventoryStockBatchNumbers, error) {
	var entities []*InventoryStockBatchNumbers
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/inventory/StockBatchNumbers?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d StockBatchNumbers entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
