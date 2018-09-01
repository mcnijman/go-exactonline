// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package purchase

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// PurchaseInvoicesEndpoint is responsible for communicating with
// the PurchaseInvoices endpoint of the Purchase service.
type PurchaseInvoicesEndpoint service

// PurchaseInvoices:
// Service: Purchase
// Entity: PurchaseInvoices
// URL: /api/v1/{division}/purchase/PurchaseInvoices
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=PurchasePurchaseInvoices
type PurchaseInvoices struct {
	// ID: A guid that is the unique identifier of the purchase invoice.
	ID *types.GUID `json:"ID,omitempty"`

	// Amount: The amount including VAT in the currency of the invoice.
	Amount *float64 `json:"Amount,omitempty"`

	// ContactPerson: Guid identifying the contact person of the supplier.
	ContactPerson *types.GUID `json:"ContactPerson,omitempty"`

	// Currency: The code of the currency of the invoiced amount.
	Currency *string `json:"Currency,omitempty"`

	// Description: The description of the invoice.
	Description *string `json:"Description,omitempty"`

	// Document: Guid identifying a document that is attached to the invoice.
	Document *types.GUID `json:"Document,omitempty"`

	// DueDate: The date before which the invoice has to be paid.
	DueDate *types.Date `json:"DueDate,omitempty"`

	// EntryNumber: The unique number of the purchase invoice. The entry number is based on a setting in the purchase journal and incremented for each new purchase invoice.
	EntryNumber *int `json:"EntryNumber,omitempty"`

	// ExchangeRate: The exchange rate between the invoice currency and the default currency of the division.
	ExchangeRate *float64 `json:"ExchangeRate,omitempty"`

	// FinancialPeriod: The financial period in which the invoice is entered.
	FinancialPeriod *int `json:"FinancialPeriod,omitempty"`

	// FinancialYear: The financial year in which the invoice is entered.
	FinancialYear *int `json:"FinancialYear,omitempty"`

	// InvoiceDate: The date on which the supplier entered the invoice.
	InvoiceDate *types.Date `json:"InvoiceDate,omitempty"`

	// Journal: The code of the purchase journal in which the invoice is entered.
	Journal *string `json:"Journal,omitempty"`

	// Modified: The date and time the invoice was last modified.
	Modified *types.Date `json:"Modified,omitempty"`

	// PaymentCondition: The code of the payment condition that is used to calculate the due date and discount.
	PaymentCondition *string `json:"PaymentCondition,omitempty"`

	// PaymentReference: Unique reference to match payments and invoices.
	PaymentReference *string `json:"PaymentReference,omitempty"`

	// PurchaseInvoiceLines: The collection of lines that belong to the purchase invoice.
	PurchaseInvoiceLines *[]byte `json:"PurchaseInvoiceLines,omitempty"`

	// Remarks: The user can enter remarks related to the invoice here.
	Remarks *string `json:"Remarks,omitempty"`

	// Source: Indicates the origin of the invoice. 1 Manual entry, 3 Purchase invoice, 4 Purchase order, 5 Web service.
	Source *int `json:"Source,omitempty"`

	// Status: The status of the invoice. 10 Draft, 20 Open, 50 Processed.
	Status *int `json:"Status,omitempty"`

	// Supplier: Guid that identifies the supplier.
	Supplier *types.GUID `json:"Supplier,omitempty"`

	// Type: Indicates the type of the purchase invoice. 8030 Direct purchase invoice, 8031 Direct purchase invoice (Credit), 8033 Purchase invoice, 8034 Purchase invoice (Credit)
	Type *int `json:"Type,omitempty"`

	// VATAmount: The total VAT amount of the purchase invoice.
	VATAmount *float64 `json:"VATAmount,omitempty"`

	// Warehouse: Guid that identifies the warehouse that will receive the purchased goods. This is mandatory for creating a direct purchase invoice.
	Warehouse *types.GUID `json:"Warehouse,omitempty"`

	// YourRef: The invoice number provided by the supplier.
	YourRef *string `json:"YourRef,omitempty"`
}

func (s *PurchaseInvoices) GetIdentifier() types.GUID {
	return *s.ID
}

// List the PurchaseInvoices entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *PurchaseInvoicesEndpoint) List(ctx context.Context, division int, all bool) ([]*PurchaseInvoices, error) {
	var entities []*PurchaseInvoices
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/purchase/PurchaseInvoices?$select=*", division)
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
