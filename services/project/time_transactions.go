// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package project

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mcnijman/go-exactonline/api"
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
	MetaData *api.MetaData `json:"__metadata,omitempty"`
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

func (e *TimeTransactions) GetPrimary() *types.GUID {
	return e.ID
}

func (s *TimeTransactionsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "project/TimeTransactions", method)
}

// List the TimeTransactions entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *TimeTransactionsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*TimeTransactions, error) {
	var entities []*TimeTransactions
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/TimeTransactions", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the TimeTransactions entitiy in the provided division.
func (s *TimeTransactionsEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*TimeTransactions, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/TimeTransactions", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &TimeTransactions{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty TimeTransactions entity
func (s *TimeTransactionsEndpoint) New() *TimeTransactions {
	return &TimeTransactions{}
}

// Create the TimeTransactions entity in the provided division.
func (s *TimeTransactionsEndpoint) Create(ctx context.Context, division int, entity *TimeTransactions) (*TimeTransactions, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/TimeTransactions", division) // #nosec
	e := &TimeTransactions{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the TimeTransactions entity in the provided division.
func (s *TimeTransactionsEndpoint) Update(ctx context.Context, division int, entity *TimeTransactions) (*TimeTransactions, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/TimeTransactions", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &TimeTransactions{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the TimeTransactions entity in the provided division.
func (s *TimeTransactionsEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/TimeTransactions", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return err
	}

	_, r, requestError := s.client.NewRequestAndDo(ctx, "DELETE", u.String(), nil, nil)
	if requestError != nil {
		return requestError
	}

	if r.StatusCode != http.StatusNoContent {
		body, _ := ioutil.ReadAll(r.Body) // #nosec
		return fmt.Errorf("Failed with status %v and body %v", r.StatusCode, body)
	}

	return nil
}
