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

// TimeTransactionsEndpoint is responsible for communicating with
// the TimeTransactions endpoint of the Manufacturing service.
type TimeTransactionsEndpoint service

// TimeTransactions:
// Service: Manufacturing
// Entity: TimeTransactions
// URL: /api/v1/{division}/manufacturing/TimeTransactions
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ManufacturingTimeTransactions
type TimeTransactions struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Activity: Manufacturing time type: Setup = 10, Run = 20
	Activity *int `json:"Activity,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Date: Date
	Date *types.Date `json:"Date,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Employee: Employee linked to the transaction
	Employee *types.GUID `json:"Employee,omitempty"`

	// Hours: Machine hours
	Hours *float64 `json:"Hours,omitempty"`

	// IsOperationFinished: Is the operation finished?
	IsOperationFinished *byte `json:"IsOperationFinished,omitempty"`

	// LaborHours: Labor Hours on the operation
	LaborHours *float64 `json:"LaborHours,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Notes linked to the time transaction
	Notes *string `json:"Notes,omitempty"`

	// PercentComplete: Percentage of the operation that is complete
	PercentComplete *float64 `json:"PercentComplete,omitempty"`

	// Quantity: Quantity
	Quantity *float64 `json:"Quantity,omitempty"`

	// RoutingStepPlan: Routing step linked to the transaction
	RoutingStepPlan *types.GUID `json:"RoutingStepPlan,omitempty"`

	// ShopOrder: Shop order linked to the transaction
	ShopOrder *types.GUID `json:"ShopOrder,omitempty"`

	// Status: Status of the transaction: Draft = 1, Rejected = 2, Submitted = 10, Final = 20
	Status *int `json:"Status,omitempty"`

	// TimedTimeTransaction: Timed time transaction linked to the transaction
	TimedTimeTransaction *types.GUID `json:"TimedTimeTransaction,omitempty"`

	// WorkCenter: Workcenter linked to the transaction
	WorkCenter *types.GUID `json:"WorkCenter,omitempty"`
}

// List the TimeTransactions entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *TimeTransactionsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*TimeTransactions, error) {
	var entities []*TimeTransactions
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/TimeTransactions", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
