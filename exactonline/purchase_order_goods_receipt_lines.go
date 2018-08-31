// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// PurchaseOrderGoodsReceiptLinesService is responsible for communicating with
// the GoodsReceiptLines endpoint of the PurchaseOrder service.
type PurchaseOrderGoodsReceiptLinesService service

// PurchaseOrderGoodsReceiptLines:
// Service: PurchaseOrder
// Entity: GoodsReceiptLines
// URL: /api/v1/{division}/purchaseorder/GoodsReceiptLines
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=PurchaseOrderGoodsReceiptLines
type PurchaseOrderGoodsReceiptLines struct {
	// ID: The unique identifier of a stock transaction for a goods receipt line. A goods receipt line can be split into multiple storage locations. In this case, multiple storage locations will have the same stock transaction ID.
	ID *GUID `json:",omitempty"`

	// BatchNumbers: Collection of batch numbers
	BatchNumbers *[]byte `json:",omitempty"`

	// Created: Creation date
	Created *Date `json:",omitempty"`

	// Creator: User ID of the creator
	Creator *GUID `json:",omitempty"`

	// CreatorFullName: Name of the creator
	CreatorFullName *string `json:",omitempty"`

	// Description: Goods receipt line description
	Description *string `json:",omitempty"`

	// Division: Division code
	Division *int `json:",omitempty"`

	// GoodsReceiptID: All the lines of a goods receipt have the same GoodsReceiptID
	GoodsReceiptID *GUID `json:",omitempty"`

	// Item: ID of the received item
	Item *GUID `json:",omitempty"`

	// ItemCode: Code of the received item
	ItemCode *string `json:",omitempty"`

	// ItemDescription: Item description
	ItemDescription *string `json:",omitempty"`

	// ItemUnitCode: Unit code of the purchase
	ItemUnitCode *string `json:",omitempty"`

	// LineNumber: Line number
	LineNumber *int `json:",omitempty"`

	// Location: ID of the storage location in the warehouse where the item is received
	Location *GUID `json:",omitempty"`

	// LocationCode: Code of the storage location in the warehouse where the item is received
	LocationCode *string `json:",omitempty"`

	// LocationDescription: Description of the storage location in the warehouse where the item is received
	LocationDescription *string `json:",omitempty"`

	// Modified: Last modified date
	Modified *Date `json:",omitempty"`

	// Modifier: User ID of the last modifier
	Modifier *GUID `json:",omitempty"`

	// ModifierFullName: Name of the last modifier
	ModifierFullName *string `json:",omitempty"`

	// Notes: Notes
	Notes *string `json:",omitempty"`

	// Project: Reference to project
	Project *GUID `json:",omitempty"`

	// ProjectCode: Project code
	ProjectCode *string `json:",omitempty"`

	// ProjectDescription: Project description
	ProjectDescription *string `json:",omitempty"`

	// PurchaseOrderID: Reference to purchase order
	PurchaseOrderID *GUID `json:",omitempty"`

	// PurchaseOrderLineID: ID of the purchase order line that is received
	PurchaseOrderLineID *GUID `json:",omitempty"`

	// PurchaseOrderNumber: Order number of the purchase order that is received
	PurchaseOrderNumber *int `json:",omitempty"`

	// QuantityOrdered: Quantity ordered
	QuantityOrdered *float64 `json:",omitempty"`

	// QuantityReceived: Quantity received
	QuantityReceived *float64 `json:",omitempty"`

	// SerialNumbers: Collection of serial numbers
	SerialNumbers *[]byte `json:",omitempty"`

	// SupplierItemCode: Supplier item code
	SupplierItemCode *string `json:",omitempty"`
}

func (s *PurchaseOrderGoodsReceiptLines) GetIdentifier() GUID {
	return *s.ID
}

// List the GoodsReceiptLines entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *PurchaseOrderGoodsReceiptLinesService) List(ctx context.Context, division int, all bool) ([]*PurchaseOrderGoodsReceiptLines, error) {
	var entities []*PurchaseOrderGoodsReceiptLines
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/purchaseorder/GoodsReceiptLines?$select=*", division)
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