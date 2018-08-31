// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// DocumentsDocumentCategoriesService is responsible for communicating with
// the DocumentCategories endpoint of the Documents service.
type DocumentsDocumentCategoriesService service

// DocumentsDocumentCategories:
// Service: Documents
// Entity: DocumentCategories
// URL: /api/v1/{division}/documents/DocumentCategories
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=DocumentsDocumentCategories
type DocumentsDocumentCategories struct {
	// ID: Primary key
	ID *GUID `json:",omitempty"`

	// Created: Creation date
	Created *Date `json:",omitempty"`

	// Description: Document category description
	Description *string `json:",omitempty"`

	// Modified: Last modified date
	Modified *Date `json:",omitempty"`
}

func (s *DocumentsDocumentCategories) GetIdentifier() GUID {
	return *s.ID
}

// List the DocumentCategories entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *DocumentsDocumentCategoriesService) List(ctx context.Context, division int, all bool) ([]*DocumentsDocumentCategories, error) {
	var entities []*DocumentsDocumentCategories
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/documents/DocumentCategories?$select=*", division)
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