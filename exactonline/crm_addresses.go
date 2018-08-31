// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// CRMAddressesService is responsible for communicating with
// the Addresses endpoint of the CRM service.
type CRMAddressesService service

// CRMAddresses:
// Service: CRM
// Entity: Addresses
// URL: /api/v1/{division}/crm/Addresses
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=CRMAddresses
type CRMAddresses struct {
	// ID: Primary key
	ID *GUID `json:",omitempty"`

	// Account: Account linked to the address
	Account *GUID `json:",omitempty"`

	// AccountIsSupplier: Indicates if the account is a supplier
	AccountIsSupplier *bool `json:",omitempty"`

	// AccountName: Name of the account
	AccountName *string `json:",omitempty"`

	// AddressLine1: First address line
	AddressLine1 *string `json:",omitempty"`

	// AddressLine2: Second address line
	AddressLine2 *string `json:",omitempty"`

	// AddressLine3: Third address line
	AddressLine3 *string `json:",omitempty"`

	// City: City
	City *string `json:",omitempty"`

	// Contact: Contact linked to Address
	Contact *GUID `json:",omitempty"`

	// ContactName: Contact name
	ContactName *string `json:",omitempty"`

	// Country: Country code
	Country *string `json:",omitempty"`

	// CountryName: Country name
	CountryName *string `json:",omitempty"`

	// Created: Creation date
	Created *Date `json:",omitempty"`

	// Creator: User ID of creator
	Creator *GUID `json:",omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:",omitempty"`

	// Division: Division code
	Division *int `json:",omitempty"`

	// Fax: Fax number
	Fax *string `json:",omitempty"`

	// FreeBoolField_01: Free boolean field 1
	FreeBoolField_01 *bool `json:",omitempty"`

	// FreeBoolField_02: Free boolean field 2
	FreeBoolField_02 *bool `json:",omitempty"`

	// FreeBoolField_03: Free boolean field 3
	FreeBoolField_03 *bool `json:",omitempty"`

	// FreeBoolField_04: Free boolean field 4
	FreeBoolField_04 *bool `json:",omitempty"`

	// FreeBoolField_05: Free boolean field 5
	FreeBoolField_05 *bool `json:",omitempty"`

	// FreeDateField_01: Free date field 1
	FreeDateField_01 *Date `json:",omitempty"`

	// FreeDateField_02: Free date field 2
	FreeDateField_02 *Date `json:",omitempty"`

	// FreeDateField_03: Free date field 3
	FreeDateField_03 *Date `json:",omitempty"`

	// FreeDateField_04: Free date field 4
	FreeDateField_04 *Date `json:",omitempty"`

	// FreeDateField_05: Free date field 5
	FreeDateField_05 *Date `json:",omitempty"`

	// FreeNumberField_01: Free number field 1
	FreeNumberField_01 *float64 `json:",omitempty"`

	// FreeNumberField_02: Free number field 2
	FreeNumberField_02 *float64 `json:",omitempty"`

	// FreeNumberField_03: Free number field 3
	FreeNumberField_03 *float64 `json:",omitempty"`

	// FreeNumberField_04: Free number field 4
	FreeNumberField_04 *float64 `json:",omitempty"`

	// FreeNumberField_05: Free number field 5
	FreeNumberField_05 *float64 `json:",omitempty"`

	// FreeTextField_01: Free text field 1
	FreeTextField_01 *string `json:",omitempty"`

	// FreeTextField_02: Free text field 2
	FreeTextField_02 *string `json:",omitempty"`

	// FreeTextField_03: Free text field 3
	FreeTextField_03 *string `json:",omitempty"`

	// FreeTextField_04: Free text field 4
	FreeTextField_04 *string `json:",omitempty"`

	// FreeTextField_05: Free text field 5
	FreeTextField_05 *string `json:",omitempty"`

	// Mailbox: Mailbox
	Mailbox *string `json:",omitempty"`

	// Main: Indicates if the address is the main address for this type
	Main *bool `json:",omitempty"`

	// Modified: Last modified date
	Modified *Date `json:",omitempty"`

	// Modifier: User ID of modifier
	Modifier *GUID `json:",omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:",omitempty"`

	// NicNumber: Last 5 digits of SIRET number which is an intern sequential number of 4 digits representing the identification of the localization of the office
	NicNumber *string `json:",omitempty"`

	// Notes: Notes for an address
	Notes *string `json:",omitempty"`

	// Phone: Phone number
	Phone *string `json:",omitempty"`

	// PhoneExtension: Phone extension
	PhoneExtension *string `json:",omitempty"`

	// Postcode: Postcode
	Postcode *string `json:",omitempty"`

	// State: State
	State *string `json:",omitempty"`

	// StateDescription: Name of the State
	StateDescription *string `json:",omitempty"`

	// Type: The type of address. Visit=1, Postal=2, Invoice=3, Delivery=4
	Type *int `json:",omitempty"`

	// Warehouse: The warehouse linked to the address, if a warehouse is linked the account will be empty. Can only be filled for type=Delivery
	Warehouse *GUID `json:",omitempty"`

	// WarehouseCode: Code of the warehoude
	WarehouseCode *string `json:",omitempty"`

	// WarehouseDescription: Description of the warehouse
	WarehouseDescription *string `json:",omitempty"`
}

func (s *CRMAddresses) GetIdentifier() GUID {
	return *s.ID
}

// List the Addresses entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *CRMAddressesService) List(ctx context.Context, division int, all bool) ([]*CRMAddresses, error) {
	var entities []*CRMAddresses
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/Addresses?$select=*", division)
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