// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// FinancialGLSchemesService is responsible for communicating with
// the GLSchemes endpoint of the Financial service.
type FinancialGLSchemesService service

// FinancialGLSchemes:
// Service: Financial
// Entity: GLSchemes
// URL: /api/v1/{division}/financial/GLSchemes
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=FinancialGLSchemes
type FinancialGLSchemes struct {
	// ID: Primary key
	ID *GUID `json:",omitempty"`

	// Code:
	Code *string `json:",omitempty"`

	// Created: Creation date
	Created *Date `json:",omitempty"`

	// Creator: User ID of creator
	Creator *GUID `json:",omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:",omitempty"`

	// Description: Description text
	Description *string `json:",omitempty"`

	// Division: Division is optional for this table. For taxonomies of Taxonomies.Type = 0 (general taxonomies), the Division is empty. For division specific taxonomies it is mandatory
	Division *int `json:",omitempty"`

	// Main: Only used for reporting schemes = division specific taxonomynamespaces. In this case, main = 1 denotes the main or default reporting scheme
	Main *byte `json:",omitempty"`

	// Modified: Last modified date
	Modified *Date `json:",omitempty"`

	// Modifier: User ID of modifier
	Modifier *GUID `json:",omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:",omitempty"`

	// TargetNamespace: URI, which is the unique identifier of the namespace
	TargetNamespace *string `json:",omitempty"`
}

func (s *FinancialGLSchemes) GetIdentifier() GUID {
	return *s.ID
}

// List the GLSchemes entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *FinancialGLSchemesService) List(ctx context.Context, division int, all bool) ([]*FinancialGLSchemes, error) {
	var entities []*FinancialGLSchemes
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/financial/GLSchemes?$select=*", division)
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