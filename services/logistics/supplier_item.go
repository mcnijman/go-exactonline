// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package logistics

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// SupplierItemEndpoint is responsible for communicating with
// the SupplierItem endpoint of the Logistics service.
type SupplierItemEndpoint service

// SupplierItem:
// Service: Logistics
// Entity: SupplierItem
// URL: /api/v1/{division}/logistics/SupplierItem
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=LogisticsSupplierItem
type SupplierItem struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// CopyRemarks: Copy purchase remarks to purchase lines
	CopyRemarks *byte `json:"CopyRemarks,omitempty"`

	// CountryOfOrigin: Country of origin code
	CountryOfOrigin *string `json:"CountryOfOrigin,omitempty"`

	// CountryOfOriginDescription: Description of country of origin
	CountryOfOriginDescription *string `json:"CountryOfOriginDescription,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Currency: Currency of item price
	Currency *string `json:"Currency,omitempty"`

	// CurrencyDescription: Description of currency of item price
	CurrencyDescription *string `json:"CurrencyDescription,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// DropShipment: Indicates that the supplier will deliver the item directly to customer. Values: 0 = No, 1 = Yes, 2 = Optional
	DropShipment *byte `json:"DropShipment,omitempty"`

	// Item: Item ID
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode: Item code
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription: Description of Item
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// MainSupplier: Indicates this is a main supplier
	MainSupplier *bool `json:"MainSupplier,omitempty"`

	// MinimumQuantity: Minimum quantity of the item for purchase, only available for Wholesale &amp; Distribution (Premium only)
	MinimumQuantity *float64 `json:"MinimumQuantity,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Notes
	Notes *string `json:"Notes,omitempty"`

	// PurchaseLeadTime: The number of days between placing an order with a supplier and receiving items from the supplier
	PurchaseLeadTime *int `json:"PurchaseLeadTime,omitempty"`

	// PurchasePrice: Purchase price
	PurchasePrice *float64 `json:"PurchasePrice,omitempty"`

	// PurchaseUnit: Unit code
	PurchaseUnit *string `json:"PurchaseUnit,omitempty"`

	// PurchaseUnitDescription: Description of unit
	PurchaseUnitDescription *string `json:"PurchaseUnitDescription,omitempty"`

	// PurchaseUnitFactor: This is the multiplication factor when going from default item unit to the unit of this price
	PurchaseUnitFactor *float64 `json:"PurchaseUnitFactor,omitempty"`

	// PurchaseVATCode: VAT code
	PurchaseVATCode *string `json:"PurchaseVATCode,omitempty"`

	// PurchaseVATCodeDescription: Description of VAT
	PurchaseVATCodeDescription *string `json:"PurchaseVATCodeDescription,omitempty"`

	// Supplier: Supplier ID
	Supplier *types.GUID `json:"Supplier,omitempty"`

	// SupplierCode: Supplier code
	SupplierCode *string `json:"SupplierCode,omitempty"`

	// SupplierDescription: Description of supplier
	SupplierDescription *string `json:"SupplierDescription,omitempty"`

	// SupplierItemCode: Supplier’s item code
	SupplierItemCode *string `json:"SupplierItemCode,omitempty"`
}

func (s *SupplierItem) GetIdentifier() types.GUID {
	return *s.ID
}

// List the SupplierItem entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SupplierItemEndpoint) List(ctx context.Context, division int, all bool) ([]*SupplierItem, error) {
	var entities []*SupplierItem
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/SupplierItem?$select=*", division)
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