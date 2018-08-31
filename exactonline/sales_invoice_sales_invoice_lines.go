// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// SalesInvoiceSalesInvoiceLinesService is responsible for communicating with
// the SalesInvoiceLines endpoint of the SalesInvoice service.
type SalesInvoiceSalesInvoiceLinesService service

// SalesInvoiceSalesInvoiceLines:
// Service: SalesInvoice
// Entity: SalesInvoiceLines
// URL: /api/v1/{division}/salesinvoice/SalesInvoiceLines
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SalesInvoiceSalesInvoiceLines
type SalesInvoiceSalesInvoiceLines struct {
	// ID:
	ID *GUID `json:"ID,omitempty"`

	// AmountDC:
	AmountDC *float64 `json:"AmountDC,omitempty"`

	// AmountFC:
	AmountFC *float64 `json:"AmountFC,omitempty"`

	// CostCenter:
	CostCenter *string `json:"CostCenter,omitempty"`

	// CostCenterDescription:
	CostCenterDescription *string `json:"CostCenterDescription,omitempty"`

	// CostUnit:
	CostUnit *string `json:"CostUnit,omitempty"`

	// CostUnitDescription:
	CostUnitDescription *string `json:"CostUnitDescription,omitempty"`

	// DeliveryDate:
	DeliveryDate *Date `json:"DeliveryDate,omitempty"`

	// Description:
	Description *string `json:"Description,omitempty"`

	// Discount:
	Discount *float64 `json:"Discount,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// Employee:
	Employee *GUID `json:"Employee,omitempty"`

	// EmployeeFullName:
	EmployeeFullName *string `json:"EmployeeFullName,omitempty"`

	// EndTime:
	EndTime *Date `json:"EndTime,omitempty"`

	// ExtraDutyAmountFC:
	ExtraDutyAmountFC *float64 `json:"ExtraDutyAmountFC,omitempty"`

	// ExtraDutyPercentage:
	ExtraDutyPercentage *float64 `json:"ExtraDutyPercentage,omitempty"`

	// GLAccount:
	GLAccount *GUID `json:"GLAccount,omitempty"`

	// GLAccountDescription:
	GLAccountDescription *string `json:"GLAccountDescription,omitempty"`

	// InvoiceID:
	InvoiceID *GUID `json:"InvoiceID,omitempty"`

	// Item:
	Item *GUID `json:"Item,omitempty"`

	// ItemCode:
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription:
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// LineNumber:
	LineNumber *int `json:"LineNumber,omitempty"`

	// NetPrice:
	NetPrice *float64 `json:"NetPrice,omitempty"`

	// Notes:
	Notes *string `json:"Notes,omitempty"`

	// Pricelist:
	Pricelist *GUID `json:"Pricelist,omitempty"`

	// PricelistDescription:
	PricelistDescription *string `json:"PricelistDescription,omitempty"`

	// Project:
	Project *GUID `json:"Project,omitempty"`

	// ProjectDescription:
	ProjectDescription *string `json:"ProjectDescription,omitempty"`

	// ProjectWBS:
	ProjectWBS *GUID `json:"ProjectWBS,omitempty"`

	// ProjectWBSDescription:
	ProjectWBSDescription *string `json:"ProjectWBSDescription,omitempty"`

	// Quantity:
	Quantity *float64 `json:"Quantity,omitempty"`

	// SalesOrder:
	SalesOrder *GUID `json:"SalesOrder,omitempty"`

	// SalesOrderLine:
	SalesOrderLine *GUID `json:"SalesOrderLine,omitempty"`

	// SalesOrderLineNumber:
	SalesOrderLineNumber *int `json:"SalesOrderLineNumber,omitempty"`

	// SalesOrderNumber:
	SalesOrderNumber *int `json:"SalesOrderNumber,omitempty"`

	// StartTime:
	StartTime *Date `json:"StartTime,omitempty"`

	// Subscription:
	Subscription *GUID `json:"Subscription,omitempty"`

	// SubscriptionDescription:
	SubscriptionDescription *string `json:"SubscriptionDescription,omitempty"`

	// TaxSchedule:
	TaxSchedule *GUID `json:"TaxSchedule,omitempty"`

	// TaxScheduleCode:
	TaxScheduleCode *string `json:"TaxScheduleCode,omitempty"`

	// TaxScheduleDescription:
	TaxScheduleDescription *string `json:"TaxScheduleDescription,omitempty"`

	// UnitCode:
	UnitCode *string `json:"UnitCode,omitempty"`

	// UnitDescription:
	UnitDescription *string `json:"UnitDescription,omitempty"`

	// UnitPrice:
	UnitPrice *float64 `json:"UnitPrice,omitempty"`

	// VATAmountDC:
	VATAmountDC *float64 `json:"VATAmountDC,omitempty"`

	// VATAmountFC:
	VATAmountFC *float64 `json:"VATAmountFC,omitempty"`

	// VATCode:
	VATCode *string `json:"VATCode,omitempty"`

	// VATCodeDescription:
	VATCodeDescription *string `json:"VATCodeDescription,omitempty"`

	// VATPercentage:
	VATPercentage *float64 `json:"VATPercentage,omitempty"`
}

func (s *SalesInvoiceSalesInvoiceLines) GetIdentifier() GUID {
	return *s.ID
}

// List the SalesInvoiceLines entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SalesInvoiceSalesInvoiceLinesService) List(ctx context.Context, division int, all bool) ([]*SalesInvoiceSalesInvoiceLines, error) {
	var entities []*SalesInvoiceSalesInvoiceLines
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/salesinvoice/SalesInvoiceLines?$select=*", division)
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

/* // Get the SalesInvoiceLines enitity, by ID.
func (s *SalesInvoiceSalesInvoiceLinesService) Get(ctx context.Context, division int, id GUID) (*SalesInvoiceSalesInvoiceLines, error) {
	var entities []*SalesInvoiceSalesInvoiceLines
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/salesinvoice/SalesInvoiceLines?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d SalesInvoiceLines entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
