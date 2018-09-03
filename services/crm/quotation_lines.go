// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package crm

import (
	"context"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

// QuotationLinesEndpoint is responsible for communicating with
// the QuotationLines endpoint of the CRM service.
type QuotationLinesEndpoint service

// QuotationLines:
// Service: CRM
// Entity: QuotationLines
// URL: /api/v1/{division}/crm/QuotationLines
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=CRMQuotationLines
type QuotationLines struct {
	// ID:
	ID *types.GUID `json:"ID,omitempty"`

	// AmountDC:
	AmountDC *float64 `json:"AmountDC,omitempty"`

	// AmountFC:
	AmountFC *float64 `json:"AmountFC,omitempty"`

	// Description:
	Description *string `json:"Description,omitempty"`

	// Discount:
	Discount *float64 `json:"Discount,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// Item:
	Item *types.GUID `json:"Item,omitempty"`

	// ItemDescription:
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// LineNumber:
	LineNumber *int `json:"LineNumber,omitempty"`

	// NetPrice:
	NetPrice *float64 `json:"NetPrice,omitempty"`

	// Notes:
	Notes *string `json:"Notes,omitempty"`

	// Quantity:
	Quantity *float64 `json:"Quantity,omitempty"`

	// QuotationID:
	QuotationID *types.GUID `json:"QuotationID,omitempty"`

	// QuotationNumber:
	QuotationNumber *int `json:"QuotationNumber,omitempty"`

	// UnitCode:
	UnitCode *string `json:"UnitCode,omitempty"`

	// UnitDescription:
	UnitDescription *string `json:"UnitDescription,omitempty"`

	// UnitPrice:
	UnitPrice *float64 `json:"UnitPrice,omitempty"`

	// VATAmountFC:
	VATAmountFC *float64 `json:"VATAmountFC,omitempty"`

	// VATCode:
	VATCode *string `json:"VATCode,omitempty"`

	// VATDescription:
	VATDescription *string `json:"VATDescription,omitempty"`

	// VATPercentage:
	VATPercentage *float64 `json:"VATPercentage,omitempty"`

	// VersionNumber:
	VersionNumber *int `json:"VersionNumber,omitempty"`
}

// List the QuotationLines entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *QuotationLinesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*QuotationLines, error) {
	var entities []*QuotationLines
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/QuotationLines", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
