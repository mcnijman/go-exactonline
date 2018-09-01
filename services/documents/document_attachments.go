// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package documents

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// DocumentAttachmentsEndpoint is responsible for communicating with
// the DocumentAttachments endpoint of the Documents service.
type DocumentAttachmentsEndpoint service

// DocumentAttachments:
// Service: Documents
// Entity: DocumentAttachments
// URL: /api/v1/{division}/documents/DocumentAttachments
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=DocumentsDocumentAttachments
type DocumentAttachments struct {
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

func (s *DocumentAttachments) GetIdentifier() types.GUID {
	return *s.ID
}

// List the DocumentAttachments entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *DocumentAttachmentsEndpoint) List(ctx context.Context, division int, all bool) ([]*DocumentAttachments, error) {
	var entities []*DocumentAttachments
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/documents/DocumentAttachments?$select=*", division)
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