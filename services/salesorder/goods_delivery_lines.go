// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package salesorder

import (
	"context"
	"encoding/json"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

// GoodsDeliveryLinesEndpoint is responsible for communicating with
// the GoodsDeliveryLines endpoint of the SalesOrder service.
type GoodsDeliveryLinesEndpoint service

// GoodsDeliveryLines:
// Service: SalesOrder
// Entity: GoodsDeliveryLines
// URL: /api/v1/{division}/salesorder/GoodsDeliveryLines
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST PUT
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SalesOrderGoodsDeliveryLines
type GoodsDeliveryLines struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID:
	ID *types.GUID `json:"ID,omitempty"`

	// BatchNumbers:
	BatchNumbers *json.RawMessage `json:"BatchNumbers,omitempty"`

	// Created:
	Created *types.Date `json:"Created,omitempty"`

	// Creator:
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// DeliveryDate:
	DeliveryDate *types.Date `json:"DeliveryDate,omitempty"`

	// Description:
	Description *string `json:"Description,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// EntryID:
	EntryID *types.GUID `json:"EntryID,omitempty"`

	// Item:
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode:
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription:
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// LineNumber:
	LineNumber *int `json:"LineNumber,omitempty"`

	// Modified:
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes:
	Notes *string `json:"Notes,omitempty"`

	// QuantityDelivered:
	QuantityDelivered *float64 `json:"QuantityDelivered,omitempty"`

	// QuantityOrdered:
	QuantityOrdered *float64 `json:"QuantityOrdered,omitempty"`

	// SalesOrderLineID:
	SalesOrderLineID *types.GUID `json:"SalesOrderLineID,omitempty"`

	// SalesOrderLineNumber:
	SalesOrderLineNumber *int `json:"SalesOrderLineNumber,omitempty"`

	// SalesOrderNumber:
	SalesOrderNumber *int `json:"SalesOrderNumber,omitempty"`

	// SerialNumbers:
	SerialNumbers *json.RawMessage `json:"SerialNumbers,omitempty"`

	// StorageLocation:
	StorageLocation *types.GUID `json:"StorageLocation,omitempty"`

	// StorageLocationCode:
	StorageLocationCode *string `json:"StorageLocationCode,omitempty"`

	// StorageLocationDescription:
	StorageLocationDescription *string `json:"StorageLocationDescription,omitempty"`

	// TrackingNumber:
	TrackingNumber *string `json:"TrackingNumber,omitempty"`

	// Unitcode:
	Unitcode *string `json:"Unitcode,omitempty"`
}

func (e *GoodsDeliveryLines) GetPrimary() *types.GUID {
	return e.ID
}

func (s *GoodsDeliveryLinesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "salesorder/GoodsDeliveryLines", method)
}

// List the GoodsDeliveryLines entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *GoodsDeliveryLinesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*GoodsDeliveryLines, error) {
	var entities []*GoodsDeliveryLines
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/GoodsDeliveryLines", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the GoodsDeliveryLines entitiy in the provided division.
func (s *GoodsDeliveryLinesEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*GoodsDeliveryLines, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/GoodsDeliveryLines", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &GoodsDeliveryLines{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty GoodsDeliveryLines entity
func (s *GoodsDeliveryLinesEndpoint) New() *GoodsDeliveryLines {
	return &GoodsDeliveryLines{}
}

// Create the GoodsDeliveryLines entity in the provided division.
func (s *GoodsDeliveryLinesEndpoint) Create(ctx context.Context, division int, entity *GoodsDeliveryLines) (*GoodsDeliveryLines, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/GoodsDeliveryLines", division) // #nosec
	e := &GoodsDeliveryLines{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the GoodsDeliveryLines entity in the provided division.
func (s *GoodsDeliveryLinesEndpoint) Update(ctx context.Context, division int, entity *GoodsDeliveryLines) (*GoodsDeliveryLines, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/GoodsDeliveryLines", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &GoodsDeliveryLines{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}
