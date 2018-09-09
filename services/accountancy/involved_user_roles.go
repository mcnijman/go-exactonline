// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package accountancy

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

// InvolvedUserRolesEndpoint is responsible for communicating with
// the InvolvedUserRoles endpoint of the Accountancy service.
type InvolvedUserRolesEndpoint service

// InvolvedUserRoles:
// Service: Accountancy
// Entity: InvolvedUserRoles
// URL: /api/v1/{division}/accountancy/InvolvedUserRoles
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyInvolvedUserRoles
type InvolvedUserRoles struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Code: Code of the involved user role
	Code *string `json:"Code,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description of role
	Description *string `json:"Description,omitempty"`

	// DescriptionTermID: Description term code of role
	DescriptionTermID *int `json:"DescriptionTermID,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`
}

func (e *InvolvedUserRoles) GetPrimary() *types.GUID {
	return e.ID
}

func (s *InvolvedUserRolesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "accountancy/InvolvedUserRoles", method)
}

// List the InvolvedUserRoles entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *InvolvedUserRolesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*InvolvedUserRoles, error) {
	var entities []*InvolvedUserRoles
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/accountancy/InvolvedUserRoles", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the InvolvedUserRoles entitiy in the provided division.
func (s *InvolvedUserRolesEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*InvolvedUserRoles, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/accountancy/InvolvedUserRoles", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &InvolvedUserRoles{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty InvolvedUserRoles entity
func (s *InvolvedUserRolesEndpoint) New() *InvolvedUserRoles {
	return &InvolvedUserRoles{}
}

// Create the InvolvedUserRoles entity in the provided division.
func (s *InvolvedUserRolesEndpoint) Create(ctx context.Context, division int, entity *InvolvedUserRoles) (*InvolvedUserRoles, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/accountancy/InvolvedUserRoles", division) // #nosec
	e := &InvolvedUserRoles{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the InvolvedUserRoles entity in the provided division.
func (s *InvolvedUserRolesEndpoint) Update(ctx context.Context, division int, entity *InvolvedUserRoles) (*InvolvedUserRoles, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/accountancy/InvolvedUserRoles", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &InvolvedUserRoles{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the InvolvedUserRoles entity in the provided division.
func (s *InvolvedUserRolesEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/accountancy/InvolvedUserRoles", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return err
	}

	_, r, requestError := s.client.NewRequestAndDo(ctx, "DELETE", u.String(), nil, nil)
	if requestError != nil {
		return requestError
	}

	if r.StatusCode != http.StatusNoContent {
		body, _ := ioutil.ReadAll(r.Body) // #nosec
		return fmt.Errorf("Failed with status %v and body %v", r.StatusCode, body)
	}

	return nil
}
