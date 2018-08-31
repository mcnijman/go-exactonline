// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// UsersUserRolesService is responsible for communicating with
// the UserRoles endpoint of the Users service.
type UsersUserRolesService service

// UsersUserRoles:
// Service: Users
// Entity: UserRoles
// URL: /api/v1/{division}/users/UserRoles
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=UsersUserRoles
type UsersUserRoles struct {
	// ID: Primary key
	ID *GUID `json:"ID,omitempty"`

	// Created: Creation date
	Created *Date `json:"Created,omitempty"`

	// Creator: User ID of the creator
	Creator *GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of the creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description
	Description *string `json:"Description,omitempty"`

	// EndDate: Indicates the date and time when te role becomes inactive for the user
	EndDate *Date `json:"EndDate,omitempty"`

	// Modified: Last modified date
	Modified *Date `json:"Modified,omitempty"`

	// Modifier: User ID of the last modifier
	Modifier *GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of the last modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Role: The role that the user is linked to
	Role *int `json:"Role,omitempty"`

	// RoleLevel: Rolelevel sets the level on which a role for a user is active. This can be: 1 = Database, 2 = Customer, 3 = Division, 100 = Transferred to accountant
	RoleLevel *int `json:"RoleLevel,omitempty"`

	// StartDate: Indicates the date when the role becomes active for the user
	StartDate *Date `json:"StartDate,omitempty"`

	// UserID: The user that is linked to the role
	UserID *GUID `json:"UserID,omitempty"`
}

func (s *UsersUserRoles) GetIdentifier() GUID {
	return *s.ID
}

// List the UserRoles entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *UsersUserRolesService) List(ctx context.Context, division int, all bool) ([]*UsersUserRoles, error) {
	var entities []*UsersUserRoles
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserRoles?$select=*", division)
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

/* // Get the UserRoles enitity, by ID.
func (s *UsersUserRolesService) Get(ctx context.Context, division int, id GUID) (*UsersUserRoles, error) {
	var entities []*UsersUserRoles
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserRoles?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d UserRoles entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
