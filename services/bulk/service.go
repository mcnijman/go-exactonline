// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package bulk

import "github.com/mcnijman/go-exactonline/api"

type service struct {
	client *api.Client
}

// BulkService is responsible for communication with the Bulk
// endpoints of the Exact Online API.
type BulkService struct {
	client *api.Client

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Endpoints available under this service
	DocumentsDocumentAttachments *DocumentsDocumentAttachmentsEndpoint
	DocumentsDocuments           *DocumentsDocumentsEndpoint
	FinancialTransactionLines    *FinancialTransactionLinesEndpoint
	SalesOrderSalesOrderLines    *SalesOrderSalesOrderLinesEndpoint
}

// NewBulkService creates a new initialized instance of the
// BulkService.
func NewBulkService(apiClient *api.Client) *BulkService {
	s := &BulkService{client: apiClient}

	s.common.client = apiClient

	s.DocumentsDocumentAttachments = (*DocumentsDocumentAttachmentsEndpoint)(&s.common)
	s.DocumentsDocuments = (*DocumentsDocumentsEndpoint)(&s.common)
	s.FinancialTransactionLines = (*FinancialTransactionLinesEndpoint)(&s.common)
	s.SalesOrderSalesOrderLines = (*SalesOrderSalesOrderLinesEndpoint)(&s.common)

	return s
}
