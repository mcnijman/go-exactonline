// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// CRMDocumentsService is responsible for communicating with
// the Documents endpoint of the CRM service.
type CRMDocumentsService service

// CRMDocuments:
// Service: CRM
// Entity: Documents
// URL: /api/v1/{division}/read/crm/Documents
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadCRMDocuments
type CRMDocuments struct {
	// ID: Primary key
	ID *GUID `json:",omitempty"`

	// Account: ID of the related account of this document
	Account *GUID `json:",omitempty"`

	// Attachments: Attachments linked to the document. Binaries are not sent in the response.
	Attachments *[]byte `json:",omitempty"`

	// Created: Creation date
	Created *Date `json:",omitempty"`

	// Creator: User ID of creator
	Creator *GUID `json:",omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:",omitempty"`

	// Division: Division code
	Division *int `json:",omitempty"`

	// DocumentDate: Entry date of the incoming document
	DocumentDate *Date `json:",omitempty"`

	// DocumentFolder: Id of document folder
	DocumentFolder *GUID `json:",omitempty"`

	// DocumentViewUrl: Url to view the document
	DocumentViewUrl *string `json:",omitempty"`

	// HasEmptyBody: Indicates that the document body is empty
	HasEmptyBody *bool `json:",omitempty"`

	// HID: Human-readable ID, formatted as xx.xxx.xxx. Unique. May not be equal to zero
	HID *int `json:",omitempty"`

	// Modified: Last modified date
	Modified *Date `json:",omitempty"`

	// Modifier: User ID of modifier
	Modifier *GUID `json:",omitempty"`

	// Opportunity: The opportunity linked to the document
	Opportunity *GUID `json:",omitempty"`

	// PurchaseInvoiceNumber: Purchase invoice number.
	PurchaseInvoiceNumber *int `json:",omitempty"`

	// PurchaseOrderNumber: Purchase order number.
	PurchaseOrderNumber *int `json:",omitempty"`

	// SalesInvoiceNumber: &#39;Our reference&#39; of the transaction that belongs to this document
	SalesInvoiceNumber *int `json:",omitempty"`

	// SalesOrderNumber: Number of the sales order
	SalesOrderNumber *int `json:",omitempty"`

	// SendMethod: Send Method
	SendMethod *int `json:",omitempty"`

	// Subject: Subject of this document
	Subject *string `json:",omitempty"`

	// Type: The document type
	Type *int `json:",omitempty"`

	// TypeDescription: Translated description of the document type. $filter and $orderby are not supported for this property.
	TypeDescription *string `json:",omitempty"`
}

func (s *CRMDocuments) GetIdentifier() GUID {
	return *s.ID
}

// List the Documents entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *CRMDocumentsService) List(ctx context.Context, division int, all bool) ([]*CRMDocuments, error) {
	var entities []*CRMDocuments
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/read/crm/Documents?$select=*", division)
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