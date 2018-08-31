// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// CRMContactsService is responsible for communicating with
// the Contacts endpoint of the CRM service.
type CRMContactsService service

// CRMContacts:
// Service: CRM
// Entity: Contacts
// URL: /api/v1/{division}/crm/Contacts
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=CRMContacts
type CRMContacts struct {
	// ID:
	ID *GUID `json:"ID,omitempty"`

	// Account:
	Account *GUID `json:"Account,omitempty"`

	// AccountIsCustomer:
	AccountIsCustomer *bool `json:"AccountIsCustomer,omitempty"`

	// AccountIsSupplier:
	AccountIsSupplier *bool `json:"AccountIsSupplier,omitempty"`

	// AccountMainContact:
	AccountMainContact *GUID `json:"AccountMainContact,omitempty"`

	// AccountName:
	AccountName *string `json:"AccountName,omitempty"`

	// AddressLine2:
	AddressLine2 *string `json:"AddressLine2,omitempty"`

	// AddressStreet:
	AddressStreet *string `json:"AddressStreet,omitempty"`

	// AddressStreetNumber:
	AddressStreetNumber *string `json:"AddressStreetNumber,omitempty"`

	// AddressStreetNumberSuffix:
	AddressStreetNumberSuffix *string `json:"AddressStreetNumberSuffix,omitempty"`

	// AllowMailing:
	AllowMailing *int `json:"AllowMailing,omitempty"`

	// BirthDate:
	BirthDate *Date `json:"BirthDate,omitempty"`

	// BirthName:
	BirthName *string `json:"BirthName,omitempty"`

	// BirthNamePrefix:
	BirthNamePrefix *string `json:"BirthNamePrefix,omitempty"`

	// BirthPlace:
	BirthPlace *string `json:"BirthPlace,omitempty"`

	// BusinessEmail:
	BusinessEmail *string `json:"BusinessEmail,omitempty"`

	// BusinessFax:
	BusinessFax *string `json:"BusinessFax,omitempty"`

	// BusinessMobile:
	BusinessMobile *string `json:"BusinessMobile,omitempty"`

	// BusinessPhone:
	BusinessPhone *string `json:"BusinessPhone,omitempty"`

	// BusinessPhoneExtension:
	BusinessPhoneExtension *string `json:"BusinessPhoneExtension,omitempty"`

	// City:
	City *string `json:"City,omitempty"`

	// Code:
	Code *string `json:"Code,omitempty"`

	// Country:
	Country *string `json:"Country,omitempty"`

	// Created:
	Created *Date `json:"Created,omitempty"`

	// Creator:
	Creator *GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// Email:
	Email *string `json:"Email,omitempty"`

	// EndDate:
	EndDate *Date `json:"EndDate,omitempty"`

	// FirstName:
	FirstName *string `json:"FirstName,omitempty"`

	// FullName:
	FullName *string `json:"FullName,omitempty"`

	// Gender:
	Gender *string `json:"Gender,omitempty"`

	// HID:
	HID *int `json:"HID,omitempty"`

	// IdentificationDate:
	IdentificationDate *Date `json:"IdentificationDate,omitempty"`

	// IdentificationDocument:
	IdentificationDocument *GUID `json:"IdentificationDocument,omitempty"`

	// IdentificationUser:
	IdentificationUser *GUID `json:"IdentificationUser,omitempty"`

	// Initials:
	Initials *string `json:"Initials,omitempty"`

	// IsAnonymised:
	IsAnonymised *byte `json:"IsAnonymised,omitempty"`

	// IsMailingExcluded:
	IsMailingExcluded *bool `json:"IsMailingExcluded,omitempty"`

	// IsMainContact:
	IsMainContact *bool `json:"IsMainContact,omitempty"`

	// JobTitleDescription:
	JobTitleDescription *string `json:"JobTitleDescription,omitempty"`

	// Language:
	Language *string `json:"Language,omitempty"`

	// LastName:
	LastName *string `json:"LastName,omitempty"`

	// LeadPurpose:
	LeadPurpose *GUID `json:"LeadPurpose,omitempty"`

	// LeadSource:
	LeadSource *GUID `json:"LeadSource,omitempty"`

	// MarketingNotes:
	MarketingNotes *string `json:"MarketingNotes,omitempty"`

	// MiddleName:
	MiddleName *string `json:"MiddleName,omitempty"`

	// Mobile:
	Mobile *string `json:"Mobile,omitempty"`

	// Modified:
	Modified *Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Nationality:
	Nationality *string `json:"Nationality,omitempty"`

	// Notes:
	Notes *string `json:"Notes,omitempty"`

	// PartnerName:
	PartnerName *string `json:"PartnerName,omitempty"`

	// PartnerNamePrefix:
	PartnerNamePrefix *string `json:"PartnerNamePrefix,omitempty"`

	// Person:
	Person *GUID `json:"Person,omitempty"`

	// Phone:
	Phone *string `json:"Phone,omitempty"`

	// PhoneExtension:
	PhoneExtension *string `json:"PhoneExtension,omitempty"`

	// Picture:
	Picture *[]byte `json:"Picture,omitempty"`

	// PictureName:
	PictureName *string `json:"PictureName,omitempty"`

	// PictureThumbnailUrl:
	PictureThumbnailUrl *string `json:"PictureThumbnailUrl,omitempty"`

	// PictureUrl:
	PictureUrl *string `json:"PictureUrl,omitempty"`

	// Postcode:
	Postcode *string `json:"Postcode,omitempty"`

	// SocialSecurityNumber:
	SocialSecurityNumber *string `json:"SocialSecurityNumber,omitempty"`

	// StartDate:
	StartDate *Date `json:"StartDate,omitempty"`

	// State:
	State *string `json:"State,omitempty"`

	// Title:
	Title *string `json:"Title,omitempty"`
}

func (s *CRMContacts) GetIdentifier() GUID {
	return *s.ID
}

// List the Contacts entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *CRMContactsService) List(ctx context.Context, division int, all bool) ([]*CRMContacts, error) {
	var entities []*CRMContacts
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/Contacts?$select=*", division)
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

/* // Get the Contacts enitity, by ID.
func (s *CRMContactsService) Get(ctx context.Context, division int, id GUID) (*CRMContacts, error) {
	var entities []*CRMContacts
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/Contacts?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d Contacts entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
