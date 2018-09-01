// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package project

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// TimeTransactionsEndpoint is responsible for communicating with
// the TimeTransactions endpoint of the Project service.
type TimeTransactionsEndpoint service

// TimeTransactions:
// Service: Project
// Entity: TimeTransactions
// URL: /api/v1/{division}/project/TimeTransactions
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ProjectTimeTransactions
type TimeTransactions struct {
	// ID:
	ID *types.GUID `json:"ID,omitempty"`

	// Account:
	Account *types.GUID `json:"Account,omitempty"`

	// AccountName:
	AccountName *string `json:"AccountName,omitempty"`

	// Activity:
	Activity *types.GUID `json:"Activity,omitempty"`

	// ActivityDescription:
	ActivityDescription *string `json:"ActivityDescription,omitempty"`

	// Amount:
	Amount *float64 `json:"Amount,omitempty"`

	// AmountFC:
	AmountFC *float64 `json:"AmountFC,omitempty"`

	// Attachment:
	Attachment *types.GUID `json:"Attachment,omitempty"`

	// Created:
	Created *types.Date `json:"Created,omitempty"`

	// Creator:
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Currency:
	Currency *string `json:"Currency,omitempty"`

	// Date:
	Date *types.Date `json:"Date,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// DivisionDescription:
	DivisionDescription *string `json:"DivisionDescription,omitempty"`

	// Employee:
	Employee *types.GUID `json:"Employee,omitempty"`

	// EndTime:
	EndTime *types.Date `json:"EndTime,omitempty"`

	// EntryNumber:
	EntryNumber *int `json:"EntryNumber,omitempty"`

	// ErrorText:
	ErrorText *string `json:"ErrorText,omitempty"`

	// HourStatus:
	HourStatus *int `json:"HourStatus,omitempty"`

	// Item:
	Item *types.GUID `json:"Item,omitempty"`

	// ItemDescription:
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// ItemDivisable:
	ItemDivisable *bool `json:"ItemDivisable,omitempty"`

	// Modified:
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes:
	Notes *string `json:"Notes,omitempty"`

	// Price:
	Price *float64 `json:"Price,omitempty"`

	// PriceFC:
	PriceFC *float64 `json:"PriceFC,omitempty"`

	// Project:
	Project *types.GUID `json:"Project,omitempty"`

	// ProjectAccount:
	ProjectAccount *types.GUID `json:"ProjectAccount,omitempty"`

	// ProjectAccountCode:
	ProjectAccountCode *string `json:"ProjectAccountCode,omitempty"`

	// ProjectAccountName:
	ProjectAccountName *string `json:"ProjectAccountName,omitempty"`

	// ProjectCode:
	ProjectCode *string `json:"ProjectCode,omitempty"`

	// ProjectDescription:
	ProjectDescription *string `json:"ProjectDescription,omitempty"`

	// Quantity:
	Quantity *float64 `json:"Quantity,omitempty"`

	// SkipValidation:
	SkipValidation *bool `json:"SkipValidation,omitempty"`

	// StartTime:
	StartTime *types.Date `json:"StartTime,omitempty"`

	// Subscription:
	Subscription *types.GUID `json:"Subscription,omitempty"`

	// SubscriptionAccount:
	SubscriptionAccount *types.GUID `json:"SubscriptionAccount,omitempty"`

	// SubscriptionAccountCode:
	SubscriptionAccountCode *string `json:"SubscriptionAccountCode,omitempty"`

	// SubscriptionAccountName:
	SubscriptionAccountName *string `json:"SubscriptionAccountName,omitempty"`

	// SubscriptionDescription:
	SubscriptionDescription *string `json:"SubscriptionDescription,omitempty"`

	// SubscriptionNumber:
	SubscriptionNumber *int `json:"SubscriptionNumber,omitempty"`

	// Type:
	Type *int `json:"Type,omitempty"`
}

func (s *TimeTransactions) GetIdentifier() types.GUID {
	return *s.ID
}

// List the TimeTransactions entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *TimeTransactionsEndpoint) List(ctx context.Context, division int, all bool) ([]*TimeTransactions, error) {
	var entities []*TimeTransactions
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/project/TimeTransactions?$select=*", division)
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