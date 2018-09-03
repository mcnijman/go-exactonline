// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package manufacturing

import (
	"context"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

// BillOfMaterialMaterialsEndpoint is responsible for communicating with
// the BillOfMaterialMaterials endpoint of the Manufacturing service.
type BillOfMaterialMaterialsEndpoint service

// BillOfMaterialMaterials:
// Service: Manufacturing
// Entity: BillOfMaterialMaterials
// URL: /api/v1/{division}/manufacturing/BillOfMaterialMaterials
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ManufacturingBillOfMaterialMaterials
type BillOfMaterialMaterials struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// AverageCost: Item average cost available when average cost method is used
	AverageCost *float64 `json:"AverageCost,omitempty"`

	// Backflush: Indicates if this is a backflush item
	Backflush *byte `json:"Backflush,omitempty"`

	// CalculatorType: Calculator type
	CalculatorType *int `json:"CalculatorType,omitempty"`

	// CostBatch: Cost batch
	CostBatch *float64 `json:"CostBatch,omitempty"`

	// CostCenter: Cost center
	CostCenter *string `json:"CostCenter,omitempty"`

	// CostCenterDescription: Cost center description
	CostCenterDescription *string `json:"CostCenterDescription,omitempty"`

	// CostUnit: Cost unit
	CostUnit *string `json:"CostUnit,omitempty"`

	// CostUnitDescription: Cost unit description
	CostUnitDescription *string `json:"CostUnitDescription,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description of the material
	Description *string `json:"Description,omitempty"`

	// DetailDrawing: Detail drawing reference
	DetailDrawing *string `json:"DetailDrawing,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// ItemVersion: Key of item version
	ItemVersion *types.GUID `json:"ItemVersion,omitempty"`

	// LineNumber: Line number
	LineNumber *int `json:"LineNumber,omitempty"`

	// NetWeight: Net weight
	NetWeight *float64 `json:"NetWeight,omitempty"`

	// NetWeightUnit: Net weight unit of measure
	NetWeightUnit *string `json:"NetWeightUnit,omitempty"`

	// Notes: Notes
	Notes *string `json:"Notes,omitempty"`

	// PartItem: Key of part item
	PartItem *types.GUID `json:"PartItem,omitempty"`

	// PartItemCode: Part item code
	PartItemCode *string `json:"PartItemCode,omitempty"`

	// PartItemCostPriceStandard: Item standard cost available when standard cost method is used
	PartItemCostPriceStandard *float64 `json:"PartItemCostPriceStandard,omitempty"`

	// PartItemDescription: Part item description
	PartItemDescription *string `json:"PartItemDescription,omitempty"`

	// Quantity: Quantity
	Quantity *float64 `json:"Quantity,omitempty"`

	// QuantityBatch: Quantity batch
	QuantityBatch *float64 `json:"QuantityBatch,omitempty"`

	// Syscreated: Creation date
	Syscreated *types.Date `json:"syscreated,omitempty"`

	// Syscreator: User ID of creator
	Syscreator *types.GUID `json:"syscreator,omitempty"`

	// Sysmodified: Modified date
	Sysmodified *types.Date `json:"sysmodified,omitempty"`

	// Sysmodifier: User ID of modifier
	Sysmodifier *types.GUID `json:"sysmodifier,omitempty"`

	// Type: Material type 1 indicates material, 2 indicates byproduct
	Type *int `json:"Type,omitempty"`
}

// List the BillOfMaterialMaterials entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *BillOfMaterialMaterialsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*BillOfMaterialMaterials, error) {
	var entities []*BillOfMaterialMaterials
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/BillOfMaterialMaterials", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
