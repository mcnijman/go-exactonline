// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package bulk

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// DocumentsDocumentAttachmentsEndpoint is responsible for communicating with
// the DocumentsDocumentAttachments endpoint of the Bulk service.
type DocumentsDocumentAttachmentsEndpoint service

// DocumentsDocumentAttachments:
// Service: Bulk
// Entity: DocumentsDocumentAttachments
// URL: /api/v1/{division}/bulk/Documents/DocumentAttachments
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=BulkDocumentsDocumentAttachments
type DocumentsDocumentAttachments struct {
	// ID:
	ID *types.GUID `json:"ID,omitempty"`

	// Attachment:
	Attachment *[]byte `json:"Attachment,omitempty"`

	// Document:
	Document *types.GUID `json:"Document,omitempty"`

	// FileName:
	FileName *string `json:"FileName,omitempty"`

	// FileSize:
	FileSize *float64 `json:"FileSize,omitempty"`

	// Url:
	Url *string `json:"Url,omitempty"`
}

func (s *DocumentsDocumentAttachments) GetIdentifier() types.GUID {
	return *s.ID
}

// List the DocumentsDocumentAttachments entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *DocumentsDocumentAttachmentsEndpoint) List(ctx context.Context, division int, all bool) ([]*DocumentsDocumentAttachments, error) {
	var entities []*DocumentsDocumentAttachments
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/bulk/Documents/DocumentAttachments?$select=*", division)
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
