// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package salesinvoice

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

// SalesInvoicesEndpoint is responsible for communicating with
// the SalesInvoices endpoint of the SalesInvoice service.
type SalesInvoicesEndpoint service

// SalesInvoices:
// Service: SalesInvoice
// Entity: SalesInvoices
// URL: /api/v1/{division}/salesinvoice/SalesInvoices
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SalesInvoiceSalesInvoices
type SalesInvoices struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// InvoiceID:
	InvoiceID *types.GUID `json:"InvoiceID,omitempty"`

	// AmountDC:
	AmountDC *float64 `json:"AmountDC,omitempty"`

	// AmountDiscount:
	AmountDiscount *float64 `json:"AmountDiscount,omitempty"`

	// AmountDiscountExclVat:
	AmountDiscountExclVat *float64 `json:"AmountDiscountExclVat,omitempty"`

	// AmountFC:
	AmountFC *float64 `json:"AmountFC,omitempty"`

	// AmountFCExclVat:
	AmountFCExclVat *float64 `json:"AmountFCExclVat,omitempty"`

	// Created:
	Created *types.Date `json:"Created,omitempty"`

	// Creator:
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Currency:
	Currency *string `json:"Currency,omitempty"`

	// DeliverTo:
	DeliverTo *types.GUID `json:"DeliverTo,omitempty"`

	// DeliverToAddress:
	DeliverToAddress *types.GUID `json:"DeliverToAddress,omitempty"`

	// DeliverToContactPerson:
	DeliverToContactPerson *types.GUID `json:"DeliverToContactPerson,omitempty"`

	// DeliverToContactPersonFullName:
	DeliverToContactPersonFullName *string `json:"DeliverToContactPersonFullName,omitempty"`

	// DeliverToName:
	DeliverToName *string `json:"DeliverToName,omitempty"`

	// Description:
	Description *string `json:"Description,omitempty"`

	// Discount:
	Discount *float64 `json:"Discount,omitempty"`

	// DiscountType:
	DiscountType *int `json:"DiscountType,omitempty"`

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

	// ExtraDutyAmountFC:
	ExtraDutyAmountFC *float64 `json:"ExtraDutyAmountFC,omitempty"`

	// GAccountAmountFC:
	GAccountAmountFC *float64 `json:"GAccountAmountFC,omitempty"`

	// InvoiceDate:
	InvoiceDate *types.Date `json:"InvoiceDate,omitempty"`

	// InvoiceNumber:
	InvoiceNumber *int `json:"InvoiceNumber,omitempty"`

	// InvoiceTo:
	InvoiceTo *types.GUID `json:"InvoiceTo,omitempty"`

	// InvoiceToContactPerson:
	InvoiceToContactPerson *types.GUID `json:"InvoiceToContactPerson,omitempty"`

	// InvoiceToContactPersonFullName:
	InvoiceToContactPersonFullName *string `json:"InvoiceToContactPersonFullName,omitempty"`

	// InvoiceToName:
	InvoiceToName *string `json:"InvoiceToName,omitempty"`

	// IsExtraDuty:
	IsExtraDuty *bool `json:"IsExtraDuty,omitempty"`

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

	// OrderDate:
	OrderDate *types.Date `json:"OrderDate,omitempty"`

	// OrderedBy:
	OrderedBy *types.GUID `json:"OrderedBy,omitempty"`

	// OrderedByContactPerson:
	OrderedByContactPerson *types.GUID `json:"OrderedByContactPerson,omitempty"`

	// OrderedByContactPersonFullName:
	OrderedByContactPersonFullName *string `json:"OrderedByContactPersonFullName,omitempty"`

	// OrderedByName:
	OrderedByName *string `json:"OrderedByName,omitempty"`

	// OrderNumber:
	OrderNumber *int `json:"OrderNumber,omitempty"`

	// PaymentCondition:
	PaymentCondition *string `json:"PaymentCondition,omitempty"`

	// PaymentConditionDescription:
	PaymentConditionDescription *string `json:"PaymentConditionDescription,omitempty"`

	// PaymentReference:
	PaymentReference *string `json:"PaymentReference,omitempty"`

	// Remarks:
	Remarks *string `json:"Remarks,omitempty"`

	// SalesInvoiceLines:
	SalesInvoiceLines *json.RawMessage `json:"SalesInvoiceLines,omitempty"`

	// Salesperson:
	Salesperson *types.GUID `json:"Salesperson,omitempty"`

	// SalespersonFullName:
	SalespersonFullName *string `json:"SalespersonFullName,omitempty"`

	// StarterSalesInvoiceStatus:
	StarterSalesInvoiceStatus *int `json:"StarterSalesInvoiceStatus,omitempty"`

	// StarterSalesInvoiceStatusDescription:
	StarterSalesInvoiceStatusDescription *string `json:"StarterSalesInvoiceStatusDescription,omitempty"`

	// Status:
	Status *int `json:"Status,omitempty"`

	// StatusDescription:
	StatusDescription *string `json:"StatusDescription,omitempty"`

	// TaxSchedule:
	TaxSchedule *types.GUID `json:"TaxSchedule,omitempty"`

	// TaxScheduleCode:
	TaxScheduleCode *string `json:"TaxScheduleCode,omitempty"`

	// TaxScheduleDescription:
	TaxScheduleDescription *string `json:"TaxScheduleDescription,omitempty"`

	// Type:
	Type *int `json:"Type,omitempty"`

	// TypeDescription:
	TypeDescription *string `json:"TypeDescription,omitempty"`

	// VATAmountDC:
	VATAmountDC *float64 `json:"VATAmountDC,omitempty"`

	// VATAmountFC:
	VATAmountFC *float64 `json:"VATAmountFC,omitempty"`

	// WithholdingTaxAmountFC:
	WithholdingTaxAmountFC *float64 `json:"WithholdingTaxAmountFC,omitempty"`

	// WithholdingTaxBaseAmount:
	WithholdingTaxBaseAmount *float64 `json:"WithholdingTaxBaseAmount,omitempty"`

	// WithholdingTaxPercentage:
	WithholdingTaxPercentage *float64 `json:"WithholdingTaxPercentage,omitempty"`

	// YourRef:
	YourRef *string `json:"YourRef,omitempty"`
}

func (e *SalesInvoices) GetPrimary() *types.GUID {
	return e.InvoiceID
}

func (s *SalesInvoicesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "salesinvoice/SalesInvoices", method)
}

// List the SalesInvoices entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SalesInvoicesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*SalesInvoices, error) {
	var entities []*SalesInvoices
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesinvoice/SalesInvoices", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the SalesInvoices entitiy in the provided division.
func (s *SalesInvoicesEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*SalesInvoices, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesinvoice/SalesInvoices", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &SalesInvoices{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty SalesInvoices entity
func (s *SalesInvoicesEndpoint) New() *SalesInvoices {
	return &SalesInvoices{}
}

// Create the SalesInvoices entity in the provided division.
func (s *SalesInvoicesEndpoint) Create(ctx context.Context, division int, entity *SalesInvoices) (*SalesInvoices, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesinvoice/SalesInvoices", division) // #nosec
	e := &SalesInvoices{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the SalesInvoices entity in the provided division.
func (s *SalesInvoicesEndpoint) Update(ctx context.Context, division int, entity *SalesInvoices) (*SalesInvoices, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesinvoice/SalesInvoices", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &SalesInvoices{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the SalesInvoices entity in the provided division.
func (s *SalesInvoicesEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesinvoice/SalesInvoices", division) // #nosec
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
