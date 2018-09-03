// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package users

import (
	"context"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

// UserRolesEndpoint is responsible for communicating with
// the UserRoles endpoint of the Users service.
type UserRolesEndpoint service

// UserRoles:
// Service: Users
// Entity: UserRoles
// URL: /api/v1/{division}/users/UserRoles
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=UsersUserRoles
type UserRoles struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of the creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of the creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description
	Description *string `json:"Description,omitempty"`

	// EndDate: Indicates the date and time when te role becomes inactive for the user
	EndDate *types.Date `json:"EndDate,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of the last modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of the last modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Role: The role that the user is linked to
	Role *int `json:"Role,omitempty"`

	// RoleLevel: Rolelevel sets the level on which a role for a user is active. This can be: 1 = Database, 2 = Customer, 3 = Division, 100 = Transferred to accountant
	RoleLevel *int `json:"RoleLevel,omitempty"`

	// StartDate: Indicates the date when the role becomes active for the user
	StartDate *types.Date `json:"StartDate,omitempty"`

	// UserID: The user that is linked to the role
	UserID *types.GUID `json:"UserID,omitempty"`
}

// List the UserRoles entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *UserRolesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*UserRoles, error) {
	var entities []*UserRoles
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserRoles", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
