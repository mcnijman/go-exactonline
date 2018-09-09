// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package crm

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

// AddressesEndpoint is responsible for communicating with
// the Addresses endpoint of the CRM service.
type AddressesEndpoint service

// Addresses:
// Service: CRM
// Entity: Addresses
// URL: /api/v1/{division}/crm/Addresses
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=CRMAddresses
type Addresses struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Account: Account linked to the address
	Account *types.GUID `json:"Account,omitempty"`

	// AccountIsSupplier: Indicates if the account is a supplier
	AccountIsSupplier *bool `json:"AccountIsSupplier,omitempty"`

	// AccountName: Name of the account
	AccountName *string `json:"AccountName,omitempty"`

	// AddressLine1: First address line
	AddressLine1 *string `json:"AddressLine1,omitempty"`

	// AddressLine2: Second address line
	AddressLine2 *string `json:"AddressLine2,omitempty"`

	// AddressLine3: Third address line
	AddressLine3 *string `json:"AddressLine3,omitempty"`

	// City: City
	City *string `json:"City,omitempty"`

	// Contact: Contact linked to Address
	Contact *types.GUID `json:"Contact,omitempty"`

	// ContactName: Contact name
	ContactName *string `json:"ContactName,omitempty"`

	// Country: Country code
	Country *string `json:"Country,omitempty"`

	// CountryName: Country name
	CountryName *string `json:"CountryName,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Fax: Fax number
	Fax *string `json:"Fax,omitempty"`

	// FreeBoolField_01: Free boolean field 1
	FreeBoolField_01 *bool `json:"FreeBoolField_01,omitempty"`

	// FreeBoolField_02: Free boolean field 2
	FreeBoolField_02 *bool `json:"FreeBoolField_02,omitempty"`

	// FreeBoolField_03: Free boolean field 3
	FreeBoolField_03 *bool `json:"FreeBoolField_03,omitempty"`

	// FreeBoolField_04: Free boolean field 4
	FreeBoolField_04 *bool `json:"FreeBoolField_04,omitempty"`

	// FreeBoolField_05: Free boolean field 5
	FreeBoolField_05 *bool `json:"FreeBoolField_05,omitempty"`

	// FreeDateField_01: Free date field 1
	FreeDateField_01 *types.Date `json:"FreeDateField_01,omitempty"`

	// FreeDateField_02: Free date field 2
	FreeDateField_02 *types.Date `json:"FreeDateField_02,omitempty"`

	// FreeDateField_03: Free date field 3
	FreeDateField_03 *types.Date `json:"FreeDateField_03,omitempty"`

	// FreeDateField_04: Free date field 4
	FreeDateField_04 *types.Date `json:"FreeDateField_04,omitempty"`

	// FreeDateField_05: Free date field 5
	FreeDateField_05 *types.Date `json:"FreeDateField_05,omitempty"`

	// FreeNumberField_01: Free number field 1
	FreeNumberField_01 *float64 `json:"FreeNumberField_01,omitempty"`

	// FreeNumberField_02: Free number field 2
	FreeNumberField_02 *float64 `json:"FreeNumberField_02,omitempty"`

	// FreeNumberField_03: Free number field 3
	FreeNumberField_03 *float64 `json:"FreeNumberField_03,omitempty"`

	// FreeNumberField_04: Free number field 4
	FreeNumberField_04 *float64 `json:"FreeNumberField_04,omitempty"`

	// FreeNumberField_05: Free number field 5
	FreeNumberField_05 *float64 `json:"FreeNumberField_05,omitempty"`

	// FreeTextField_01: Free text field 1
	FreeTextField_01 *string `json:"FreeTextField_01,omitempty"`

	// FreeTextField_02: Free text field 2
	FreeTextField_02 *string `json:"FreeTextField_02,omitempty"`

	// FreeTextField_03: Free text field 3
	FreeTextField_03 *string `json:"FreeTextField_03,omitempty"`

	// FreeTextField_04: Free text field 4
	FreeTextField_04 *string `json:"FreeTextField_04,omitempty"`

	// FreeTextField_05: Free text field 5
	FreeTextField_05 *string `json:"FreeTextField_05,omitempty"`

	// Mailbox: Mailbox
	Mailbox *string `json:"Mailbox,omitempty"`

	// Main: Indicates if the address is the main address for this type
	Main *bool `json:"Main,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// NicNumber: Last 5 digits of SIRET number which is an intern sequential number of 4 digits representing the identification of the localization of the office
	NicNumber *string `json:"NicNumber,omitempty"`

	// Notes: Notes for an address
	Notes *string `json:"Notes,omitempty"`

	// Phone: Phone number
	Phone *string `json:"Phone,omitempty"`

	// PhoneExtension: Phone extension
	PhoneExtension *string `json:"PhoneExtension,omitempty"`

	// Postcode: Postcode
	Postcode *string `json:"Postcode,omitempty"`

	// State: State
	State *string `json:"State,omitempty"`

	// StateDescription: Name of the State
	StateDescription *string `json:"StateDescription,omitempty"`

	// Type: The type of address. Visit=1, Postal=2, Invoice=3, Delivery=4
	Type *int `json:"Type,omitempty"`

	// Warehouse: The warehouse linked to the address, if a warehouse is linked the account will be empty. Can only be filled for type=Delivery
	Warehouse *types.GUID `json:"Warehouse,omitempty"`

	// WarehouseCode: Code of the warehoude
	WarehouseCode *string `json:"WarehouseCode,omitempty"`

	// WarehouseDescription: Description of the warehouse
	WarehouseDescription *string `json:"WarehouseDescription,omitempty"`
}

func (e *Addresses) GetPrimary() *types.GUID {
	return e.ID
}

func (s *AddressesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "crm/Addresses", method)
}

// List the Addresses entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *AddressesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*Addresses, error) {
	var entities []*Addresses
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/Addresses", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the Addresses entitiy in the provided division.
func (s *AddressesEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*Addresses, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/Addresses", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &Addresses{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty Addresses entity
func (s *AddressesEndpoint) New() *Addresses {
	return &Addresses{}
}

// Create the Addresses entity in the provided division.
func (s *AddressesEndpoint) Create(ctx context.Context, division int, entity *Addresses) (*Addresses, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/Addresses", division) // #nosec
	e := &Addresses{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the Addresses entity in the provided division.
func (s *AddressesEndpoint) Update(ctx context.Context, division int, entity *Addresses) (*Addresses, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/Addresses", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &Addresses{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the Addresses entity in the provided division.
func (s *AddressesEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/Addresses", division) // #nosec
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
