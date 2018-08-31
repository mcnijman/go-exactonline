// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// InventoryBatchNumbersService is responsible for communicating with
// the BatchNumbers endpoint of the Inventory service.
type InventoryBatchNumbersService service

// InventoryBatchNumbers:
// Service: Inventory
// Entity: BatchNumbers
// URL: /api/v1/{division}/inventory/BatchNumbers
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=InventoryBatchNumbers
type InventoryBatchNumbers struct {
	// ID: Primary key
	ID *GUID `json:"ID,omitempty"`

	// AvailableQuantity: Available quantity of this batch number
	AvailableQuantity *float64 `json:"AvailableQuantity,omitempty"`

	// BatchNumber: Human readable batch number
	BatchNumber *string `json:"BatchNumber,omitempty"`

	// Created: Creation date
	Created *Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// ExpiryDate: Expiry date of effective period for batch number
	ExpiryDate *Date `json:"ExpiryDate,omitempty"`

	// IsBlocked: Boolean value indicating whether or not the batch number is blocked
	IsBlocked *byte `json:"IsBlocked,omitempty"`

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

	// Remarks: Remarks
	Remarks *string `json:"Remarks,omitempty"`

	// StorageLocations: Total quantity available per location
	StorageLocations *[]byte `json:"StorageLocations,omitempty"`

	// Warehouses: Total quantity available per warehouse
	Warehouses *[]byte `json:"Warehouses,omitempty"`
}

func (s *InventoryBatchNumbers) GetIdentifier() GUID {
	return *s.ID
}

// List the BatchNumbers entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *InventoryBatchNumbersService) List(ctx context.Context, division int, all bool) ([]*InventoryBatchNumbers, error) {
	var entities []*InventoryBatchNumbers
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/inventory/BatchNumbers?$select=*", division)
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

/* // Get the BatchNumbers enitity, by ID.
func (s *InventoryBatchNumbersService) Get(ctx context.Context, division int, id GUID) (*InventoryBatchNumbers, error) {
	var entities []*InventoryBatchNumbers
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/inventory/BatchNumbers?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d BatchNumbers entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
