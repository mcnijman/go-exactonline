// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// FinancialOutstandingInvoicesOverviewService is responsible for communicating with
// the OutstandingInvoicesOverview endpoint of the Financial service.
type FinancialOutstandingInvoicesOverviewService service

// FinancialOutstandingInvoicesOverview:
// Service: Financial
// Entity: OutstandingInvoicesOverview
// URL: /api/v1/{division}/read/financial/OutstandingInvoicesOverview
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadFinancialOutstandingInvoicesOverview
type FinancialOutstandingInvoicesOverview struct {
	// CurrencyCode: Primary key
	CurrencyCode *string `json:"CurrencyCode,omitempty"`

	// OutstandingPayableInvoiceAmount: Total invoice amount to be paid
	OutstandingPayableInvoiceAmount *float64 `json:"OutstandingPayableInvoiceAmount,omitempty"`

	// OutstandingPayableInvoiceCount: Number of invoices to be paid
	OutstandingPayableInvoiceCount *float64 `json:"OutstandingPayableInvoiceCount,omitempty"`

	// OutstandingReceivableInvoiceAmount: Total invoice amount to be received
	OutstandingReceivableInvoiceAmount *float64 `json:"OutstandingReceivableInvoiceAmount,omitempty"`

	// OutstandingReceivableInvoiceCount: Number of invoices to be received
	OutstandingReceivableInvoiceCount *float64 `json:"OutstandingReceivableInvoiceCount,omitempty"`

	// OverduePayableInvoiceAmount: Total payable invoice amount that is overdue
	OverduePayableInvoiceAmount *float64 `json:"OverduePayableInvoiceAmount,omitempty"`

	// OverduePayableInvoiceCount: Number of payable invoices that are overdue
	OverduePayableInvoiceCount *float64 `json:"OverduePayableInvoiceCount,omitempty"`

	// OverdueReceivableInvoiceAmount: Total receivable invoice amount that is overdue
	OverdueReceivableInvoiceAmount *float64 `json:"OverdueReceivableInvoiceAmount,omitempty"`

	// OverdueReceivableInvoiceCount: Number of receivable invoices that are overdue
	OverdueReceivableInvoiceCount *float64 `json:"OverdueReceivableInvoiceCount,omitempty"`
}

func (s *FinancialOutstandingInvoicesOverview) GetIdentifier() string {
	return *s.CurrencyCode
}

// List the OutstandingInvoicesOverview entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *FinancialOutstandingInvoicesOverviewService) List(ctx context.Context, division int, all bool) ([]*FinancialOutstandingInvoicesOverview, error) {
	var entities []*FinancialOutstandingInvoicesOverview
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/OutstandingInvoicesOverview?$select=*", division)
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

/* // Get the OutstandingInvoicesOverview enitity, by CurrencyCode.
func (s *FinancialOutstandingInvoicesOverviewService) Get(ctx context.Context, division int, id string) (*FinancialOutstandingInvoicesOverview, error) {
	var entities []*FinancialOutstandingInvoicesOverview
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/OutstandingInvoicesOverview?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d OutstandingInvoicesOverview entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
