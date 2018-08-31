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
	CurrencyCode *string `json:",omitempty"`

	// OutstandingPayableInvoiceAmount: Total invoice amount to be paid
	OutstandingPayableInvoiceAmount *float64 `json:",omitempty"`

	// OutstandingPayableInvoiceCount: Number of invoices to be paid
	OutstandingPayableInvoiceCount *float64 `json:",omitempty"`

	// OutstandingReceivableInvoiceAmount: Total invoice amount to be received
	OutstandingReceivableInvoiceAmount *float64 `json:",omitempty"`

	// OutstandingReceivableInvoiceCount: Number of invoices to be received
	OutstandingReceivableInvoiceCount *float64 `json:",omitempty"`

	// OverduePayableInvoiceAmount: Total payable invoice amount that is overdue
	OverduePayableInvoiceAmount *float64 `json:",omitempty"`

	// OverduePayableInvoiceCount: Number of payable invoices that are overdue
	OverduePayableInvoiceCount *float64 `json:",omitempty"`

	// OverdueReceivableInvoiceAmount: Total receivable invoice amount that is overdue
	OverdueReceivableInvoiceAmount *float64 `json:",omitempty"`

	// OverdueReceivableInvoiceCount: Number of receivable invoices that are overdue
	OverdueReceivableInvoiceCount *float64 `json:",omitempty"`
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