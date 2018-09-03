// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package crm

import (
	"context"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

// DocumentsAttachmentsEndpoint is responsible for communicating with
// the DocumentsAttachments endpoint of the CRM service.
type DocumentsAttachmentsEndpoint service

// DocumentsAttachments:
// Service: CRM
// Entity: DocumentsAttachments
// URL: /api/v1/{division}/read/crm/DocumentsAttachments
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadCRMDocumentsAttachments
type DocumentsAttachments struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// AttachmentFileName: Filename of the attachment
	AttachmentFileName *string `json:"AttachmentFileName,omitempty"`

	// AttachmentFileSize: File size of the attachment
	AttachmentFileSize *float64 `json:"AttachmentFileSize,omitempty"`

	// AttachmentUrl: Url for downloading the attachment. To get the file in its original format (xml, jpg, pdf, etc.) append &amp;Download=1 to the url.
	AttachmentUrl *string `json:"AttachmentUrl,omitempty"`

	// CanShowInWebView:
	CanShowInWebView *bool `json:"CanShowInWebView,omitempty"`
}

// List the DocumentsAttachments entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *DocumentsAttachmentsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*DocumentsAttachments, error) {
	var entities []*DocumentsAttachments
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/crm/DocumentsAttachments", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
