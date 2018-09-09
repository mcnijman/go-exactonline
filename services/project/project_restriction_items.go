// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package project

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

// ProjectRestrictionItemsEndpoint is responsible for communicating with
// the ProjectRestrictionItems endpoint of the Project service.
type ProjectRestrictionItemsEndpoint service

// ProjectRestrictionItems:
// Service: Project
// Entity: ProjectRestrictionItems
// URL: /api/v1/{division}/project/ProjectRestrictionItems
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ProjectProjectRestrictionItems
type ProjectRestrictionItems struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Created: Date created
	Created *types.Date `json:"Created,omitempty"`

	// Creator: Creator user ID
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Creator name
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Item: Item linked to the restriction
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode: Item code
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription: Description of the item
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// ItemIsTime: Indicates if the item is a time unit item
	ItemIsTime *byte `json:"ItemIsTime,omitempty"`

	// Modified: Date modified
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: Modifier user ID
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Modifier name
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Project: Project linked to the restriction
	Project *types.GUID `json:"Project,omitempty"`

	// ProjectCode: Project code
	ProjectCode *string `json:"ProjectCode,omitempty"`

	// ProjectDescription: Project description
	ProjectDescription *string `json:"ProjectDescription,omitempty"`
}

func (e *ProjectRestrictionItems) GetPrimary() *types.GUID {
	return e.ID
}

func (s *ProjectRestrictionItemsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "project/ProjectRestrictionItems", method)
}

// List the ProjectRestrictionItems entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ProjectRestrictionItemsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*ProjectRestrictionItems, error) {
	var entities []*ProjectRestrictionItems
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectRestrictionItems", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the ProjectRestrictionItems entitiy in the provided division.
func (s *ProjectRestrictionItemsEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*ProjectRestrictionItems, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectRestrictionItems", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &ProjectRestrictionItems{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty ProjectRestrictionItems entity
func (s *ProjectRestrictionItemsEndpoint) New() *ProjectRestrictionItems {
	return &ProjectRestrictionItems{}
}

// Create the ProjectRestrictionItems entity in the provided division.
func (s *ProjectRestrictionItemsEndpoint) Create(ctx context.Context, division int, entity *ProjectRestrictionItems) (*ProjectRestrictionItems, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectRestrictionItems", division) // #nosec
	e := &ProjectRestrictionItems{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the ProjectRestrictionItems entity in the provided division.
func (s *ProjectRestrictionItemsEndpoint) Update(ctx context.Context, division int, entity *ProjectRestrictionItems) (*ProjectRestrictionItems, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectRestrictionItems", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &ProjectRestrictionItems{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the ProjectRestrictionItems entity in the provided division.
func (s *ProjectRestrictionItemsEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectRestrictionItems", division) // #nosec
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
