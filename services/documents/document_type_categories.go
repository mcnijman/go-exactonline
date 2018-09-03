// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package documents

import (
	"context"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

// DocumentTypeCategoriesEndpoint is responsible for communicating with
// the DocumentTypeCategories endpoint of the Documents service.
type DocumentTypeCategoriesEndpoint service

// DocumentTypeCategories:
// Service: Documents
// Entity: DocumentTypeCategories
// URL: /api/v1/{division}/documents/DocumentTypeCategories
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=DocumentsDocumentTypeCategories
type DocumentTypeCategories struct {
	// ID: Primary key
	ID *int `json:"ID,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Description: Document category type description
	Description *string `json:"Description,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`
}

// List the DocumentTypeCategories entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *DocumentTypeCategoriesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*DocumentTypeCategories, error) {
	var entities []*DocumentTypeCategories
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/documents/DocumentTypeCategories", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
