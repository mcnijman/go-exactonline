// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// FinancialReceivablesListService is responsible for communicating with
// the ReceivablesList endpoint of the Financial service.
type FinancialReceivablesListService service

// FinancialReceivablesList:
// Service: Financial
// Entity: ReceivablesList
// URL: /api/v1/{division}/read/financial/ReceivablesList
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadFinancialReceivablesList
type FinancialReceivablesList struct {
	// HID: Primary key, human readable ID
	HID *int64 `json:",omitempty"`

	// AccountCode: Code of Account
	AccountCode *string `json:",omitempty"`

	// AccountId: Reference to the account
	AccountId *GUID `json:",omitempty"`

	// AccountName: Name of Account
	AccountName *string `json:",omitempty"`

	// Amount: Amount
	Amount *float64 `json:",omitempty"`

	// AmountInTransit: Amount in transit
	AmountInTransit *float64 `json:",omitempty"`

	// CurrencyCode: Code of Currency
	CurrencyCode *string `json:",omitempty"`

	// Description: Description
	Description *string `json:",omitempty"`

	// DueDate: Date the invoice is due (This due date is not the discount due date)
	DueDate *Date `json:",omitempty"`

	// EntryNumber: Entry number
	EntryNumber *int `json:",omitempty"`

	// Id: Obsolete
	Id *GUID `json:",omitempty"`

	// InvoiceDate: Invoice date
	InvoiceDate *Date `json:",omitempty"`

	// InvoiceNumber: Invoice number. The value is 0 when the invoice number of the linked transaction is empty.
	InvoiceNumber *int `json:",omitempty"`

	// JournalCode: Code of Journal
	JournalCode *string `json:",omitempty"`

	// JournalDescription: Description of Journal
	JournalDescription *string `json:",omitempty"`

	// YourRef: Your reference
	YourRef *string `json:",omitempty"`
}

func (s *FinancialReceivablesList) GetIdentifier() int64 {
	return *s.HID
}

// List the ReceivablesList entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *FinancialReceivablesListService) List(ctx context.Context, division int, all bool) ([]*FinancialReceivablesList, error) {
	var entities []*FinancialReceivablesList
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/ReceivablesList?$select=*", division)
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