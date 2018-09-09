// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package financial

import (
	"context"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

// ReceivablesListEndpoint is responsible for communicating with
// the ReceivablesList endpoint of the Financial service.
type ReceivablesListEndpoint service

// ReceivablesList:
// Service: Financial
// Entity: ReceivablesList
// URL: /api/v1/{division}/read/financial/ReceivablesList
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadFinancialReceivablesList
type ReceivablesList struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// HID: Primary key, human readable ID
	HID *int64 `json:"HID,omitempty"`

	// AccountCode: Code of Account
	AccountCode *string `json:"AccountCode,omitempty"`

	// AccountId: Reference to the account
	AccountId *types.GUID `json:"AccountId,omitempty"`

	// AccountName: Name of Account
	AccountName *string `json:"AccountName,omitempty"`

	// Amount: Amount
	Amount *float64 `json:"Amount,omitempty"`

	// AmountInTransit: Amount in transit
	AmountInTransit *float64 `json:"AmountInTransit,omitempty"`

	// CurrencyCode: Code of Currency
	CurrencyCode *string `json:"CurrencyCode,omitempty"`

	// Description: Description
	Description *string `json:"Description,omitempty"`

	// DueDate: Date the invoice is due (This due date is not the discount due date)
	DueDate *types.Date `json:"DueDate,omitempty"`

	// EntryNumber: Entry number
	EntryNumber *int `json:"EntryNumber,omitempty"`

	// Id: Obsolete
	Id *types.GUID `json:"Id,omitempty"`

	// InvoiceDate: Invoice date
	InvoiceDate *types.Date `json:"InvoiceDate,omitempty"`

	// InvoiceNumber: Invoice number. The value is 0 when the invoice number of the linked transaction is empty.
	InvoiceNumber *int `json:"InvoiceNumber,omitempty"`

	// JournalCode: Code of Journal
	JournalCode *string `json:"JournalCode,omitempty"`

	// JournalDescription: Description of Journal
	JournalDescription *string `json:"JournalDescription,omitempty"`

	// YourRef: Your reference
	YourRef *string `json:"YourRef,omitempty"`
}

func (e *ReceivablesList) GetPrimary() *int64 {
	return e.HID
}

func (s *ReceivablesListEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "financial/ReceivablesList", method)
}

// List the ReceivablesList entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ReceivablesListEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*ReceivablesList, error) {
	var entities []*ReceivablesList
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/ReceivablesList", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the ReceivablesList entitiy in the provided division.
func (s *ReceivablesListEndpoint) Get(ctx context.Context, division int, id *int64) (*ReceivablesList, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/ReceivablesList", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &ReceivablesList{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
