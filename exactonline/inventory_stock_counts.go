// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// InventoryStockCountsService is responsible for communicating with
// the StockCounts endpoint of the Inventory service.
type InventoryStockCountsService service

// InventoryStockCounts:
// Service: Inventory
// Entity: StockCounts
// URL: /api/v1/{division}/inventory/StockCounts
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=InventoryStockCounts
type InventoryStockCounts struct {
	// StockCountID: Primary key
	StockCountID *GUID `json:"StockCountID,omitempty"`

	// CountedBy: Stock count user
	CountedBy *GUID `json:"CountedBy,omitempty"`

	// Created: Creation date
	Created *Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description of the stock count
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// EntryNumber: Entry number of the stock transaction
	EntryNumber *int `json:"EntryNumber,omitempty"`

	// Modified: Last modified date
	Modified *Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// OffsetGLInventory: Offset GL account of inventory
	OffsetGLInventory *GUID `json:"OffsetGLInventory,omitempty"`

	// OffsetGLInventoryCode: GLAccount code
	OffsetGLInventoryCode *string `json:"OffsetGLInventoryCode,omitempty"`

	// OffsetGLInventoryDescription: GLAccount description
	OffsetGLInventoryDescription *string `json:"OffsetGLInventoryDescription,omitempty"`

	// Source: Source of stock count entry: 1-Manual entry, 2-Import, 3-Stock count, 4-Web service
	Source *int `json:"Source,omitempty"`

	// Status: Stock count status: 12-Draft, 21-Processed
	Status *int `json:"Status,omitempty"`

	// StockCountDate: Stock count date
	StockCountDate *Date `json:"StockCountDate,omitempty"`

	// StockCountLines: Collection of stock count lines
	StockCountLines *[]byte `json:"StockCountLines,omitempty"`

	// StockCountNumber: Human readable id of the stock count
	StockCountNumber *int `json:"StockCountNumber,omitempty"`

	// Warehouse: Warehouse
	Warehouse *GUID `json:"Warehouse,omitempty"`

	// WarehouseCode: Code of Warehouse
	WarehouseCode *string `json:"WarehouseCode,omitempty"`

	// WarehouseDescription: Description of Warehouse
	WarehouseDescription *string `json:"WarehouseDescription,omitempty"`
}

func (s *InventoryStockCounts) GetIdentifier() GUID {
	return *s.StockCountID
}

// List the StockCounts entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *InventoryStockCountsService) List(ctx context.Context, division int, all bool) ([]*InventoryStockCounts, error) {
	var entities []*InventoryStockCounts
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/inventory/StockCounts?$select=*", division)
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

/* // Get the StockCounts enitity, by StockCountID.
func (s *InventoryStockCountsService) Get(ctx context.Context, division int, id GUID) (*InventoryStockCounts, error) {
	var entities []*InventoryStockCounts
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/inventory/StockCounts?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d StockCounts entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
