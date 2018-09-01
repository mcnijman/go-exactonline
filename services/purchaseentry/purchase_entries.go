// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package purchaseentry

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// PurchaseEntriesEndpoint is responsible for communicating with
// the PurchaseEntries endpoint of the PurchaseEntry service.
type PurchaseEntriesEndpoint service

// PurchaseEntries:
// Service: PurchaseEntry
// Entity: PurchaseEntries
// URL: /api/v1/{division}/purchaseentry/PurchaseEntries
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=PurchaseEntryPurchaseEntries
type PurchaseEntries struct {
	// EntryID:
	EntryID *types.GUID `json:"EntryID,omitempty"`

	// AmountDC:
	AmountDC *float64 `json:"AmountDC,omitempty"`

	// AmountFC:
	AmountFC *float64 `json:"AmountFC,omitempty"`

	// BatchNumber:
	BatchNumber *int `json:"BatchNumber,omitempty"`

	// Created:
	Created *types.Date `json:"Created,omitempty"`

	// Creator:
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Currency:
	Currency *string `json:"Currency,omitempty"`

	// Description:
	Description *string `json:"Description,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// Document:
	Document *types.GUID `json:"Document,omitempty"`

	// DocumentNumber:
	DocumentNumber *int `json:"DocumentNumber,omitempty"`

	// DocumentSubject:
	DocumentSubject *string `json:"DocumentSubject,omitempty"`

	// DueDate:
	DueDate *types.Date `json:"DueDate,omitempty"`

	// EntryDate:
	EntryDate *types.Date `json:"EntryDate,omitempty"`

	// EntryNumber:
	EntryNumber *int `json:"EntryNumber,omitempty"`

	// ExternalLinkDescription:
	ExternalLinkDescription *string `json:"ExternalLinkDescription,omitempty"`

	// ExternalLinkReference:
	ExternalLinkReference *string `json:"ExternalLinkReference,omitempty"`

	// GAccountAmountFC:
	GAccountAmountFC *float64 `json:"GAccountAmountFC,omitempty"`

	// InvoiceNumber:
	InvoiceNumber *int `json:"InvoiceNumber,omitempty"`

	// Journal:
	Journal *string `json:"Journal,omitempty"`

	// JournalDescription:
	JournalDescription *string `json:"JournalDescription,omitempty"`

	// Modified:
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// OrderNumber:
	OrderNumber *int `json:"OrderNumber,omitempty"`

	// PaymentCondition:
	PaymentCondition *string `json:"PaymentCondition,omitempty"`

	// PaymentConditionDescription:
	PaymentConditionDescription *string `json:"PaymentConditionDescription,omitempty"`

	// PaymentReference:
	PaymentReference *string `json:"PaymentReference,omitempty"`

	// ProcessNumber:
	ProcessNumber *int `json:"ProcessNumber,omitempty"`

	// PurchaseEntryLines:
	PurchaseEntryLines *[]byte `json:"PurchaseEntryLines,omitempty"`

	// Rate:
	Rate *float64 `json:"Rate,omitempty"`

	// ReportingPeriod:
	ReportingPeriod *int `json:"ReportingPeriod,omitempty"`

	// ReportingYear:
	ReportingYear *int `json:"ReportingYear,omitempty"`

	// Reversal:
	Reversal *bool `json:"Reversal,omitempty"`

	// Status:
	Status *int `json:"Status,omitempty"`

	// StatusDescription:
	StatusDescription *string `json:"StatusDescription,omitempty"`

	// Supplier:
	Supplier *types.GUID `json:"Supplier,omitempty"`

	// SupplierName:
	SupplierName *string `json:"SupplierName,omitempty"`

	// Type:
	Type *int `json:"Type,omitempty"`

	// TypeDescription:
	TypeDescription *string `json:"TypeDescription,omitempty"`

	// VATAmountDC:
	VATAmountDC *float64 `json:"VATAmountDC,omitempty"`

	// VATAmountFC:
	VATAmountFC *float64 `json:"VATAmountFC,omitempty"`

	// YourRef:
	YourRef *string `json:"YourRef,omitempty"`
}

func (s *PurchaseEntries) GetIdentifier() types.GUID {
	return *s.EntryID
}

// List the PurchaseEntries entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *PurchaseEntriesEndpoint) List(ctx context.Context, division int, all bool) ([]*PurchaseEntries, error) {
	var entities []*PurchaseEntries
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/purchaseentry/PurchaseEntries?$select=*", division)
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
