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

// DocumentTypesEndpoint is responsible for communicating with
// the DocumentTypes endpoint of the Documents service.
type DocumentTypesEndpoint service

// DocumentTypes:
// Service: Documents
// Entity: DocumentTypes
// URL: /api/v1/{division}/documents/DocumentTypes
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=DocumentsDocumentTypes
type DocumentTypes struct {
	// ID: Primary key
	ID *int `json:"ID,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Description: Document type description
	Description *string `json:"Description,omitempty"`

	// DocumentIsCreatable: Indicates if documents of this type can be created
	DocumentIsCreatable *bool `json:"DocumentIsCreatable,omitempty"`

	// DocumentIsDeletable: Indicates if documents of this type can be deleted
	DocumentIsDeletable *bool `json:"DocumentIsDeletable,omitempty"`

	// DocumentIsUpdatable: Indicates if documents of this type can be updated
	DocumentIsUpdatable *bool `json:"DocumentIsUpdatable,omitempty"`

	// DocumentIsViewable: Indicates if documents of this type can be retrieved
	DocumentIsViewable *bool `json:"DocumentIsViewable,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// TypeCategory: ID of the document type category
	TypeCategory *int `json:"TypeCategory,omitempty"`
}

func (s *DocumentTypes) GetIdentifier() int {
	return *s.ID
}

// List the DocumentTypes entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *DocumentTypesEndpoint) List(ctx context.Context, division int, all bool) ([]*DocumentTypes, error) {
	var entities []*DocumentTypes
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/documents/DocumentTypes?$select=*", division)
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