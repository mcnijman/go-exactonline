// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package system

import (
	"context"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

// MeEndpoint is responsible for communicating with
// the Me endpoint of the System service.
type MeEndpoint service

// Me:
// Service: System
// Entity: Me
// URL: /api/v1/current/Me
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SystemSystemMe
type Me struct {
	// UserID: Primary key
	UserID *types.GUID `json:"UserID,omitempty"`

	// CurrentDivision: Division number that is currently used in the API. You should use a division number in the url
	CurrentDivision *int `json:"CurrentDivision,omitempty"`

	// DivisionCustomer: Owner account of the division
	DivisionCustomer *types.GUID `json:"DivisionCustomer,omitempty"`

	// DivisionCustomerCode: Owner account code of the division
	DivisionCustomerCode *string `json:"DivisionCustomerCode,omitempty"`

	// DivisionCustomerName: Owner account name of the division
	DivisionCustomerName *string `json:"DivisionCustomerName,omitempty"`

	// DivisionCustomerSiretNumber: Owner account SIRET Number of the division for French legislation
	DivisionCustomerSiretNumber *string `json:"DivisionCustomerSiretNumber,omitempty"`

	// DivisionCustomerVatNumber: Owner account VAT Number of the division
	DivisionCustomerVatNumber *string `json:"DivisionCustomerVatNumber,omitempty"`

	// Email: Email address of the user
	Email *string `json:"Email,omitempty"`

	// EmployeeID: Employee ID
	EmployeeID *types.GUID `json:"EmployeeID,omitempty"`

	// FirstName: First name
	FirstName *string `json:"FirstName,omitempty"`

	// FullName: Full name of the user
	FullName *string `json:"FullName,omitempty"`

	// Gender: Gender: M=Male, V=Female, O=Unknown
	Gender *string `json:"Gender,omitempty"`

	// Initials: Initials
	Initials *string `json:"Initials,omitempty"`

	// Language: Language spoken by this user
	Language *string `json:"Language,omitempty"`

	// LanguageCode: Language (culture) that is used in Exact Online
	LanguageCode *string `json:"LanguageCode,omitempty"`

	// LastName: Last name
	LastName *string `json:"LastName,omitempty"`

	// Legislation: Legislation
	Legislation *int64 `json:"Legislation,omitempty"`

	// MiddleName: Middle name
	MiddleName *string `json:"MiddleName,omitempty"`

	// Mobile: Mobile phone
	Mobile *string `json:"Mobile,omitempty"`

	// Nationality: Nationality
	Nationality *string `json:"Nationality,omitempty"`

	// Phone: Phone number
	Phone *string `json:"Phone,omitempty"`

	// PhoneExtension: Phone number extension
	PhoneExtension *string `json:"PhoneExtension,omitempty"`

	// PictureUrl: Url that can be used to retrieve the picture of the user
	PictureUrl *string `json:"PictureUrl,omitempty"`

	// ServerTime: The current date and time in Exact Online
	ServerTime *string `json:"ServerTime,omitempty"`

	// ServerUtcOffset: The time difference with UTC in seconds
	ServerUtcOffset *float64 `json:"ServerUtcOffset,omitempty"`

	// ThumbnailPicture: Binary thumbnail picture of this user
	ThumbnailPicture *[]byte `json:"ThumbnailPicture,omitempty"`

	// ThumbnailPictureFormat: File type of the picture
	ThumbnailPictureFormat *string `json:"ThumbnailPictureFormat,omitempty"`

	// Title: Title
	Title *string `json:"Title,omitempty"`

	// UserName: Login name of the user
	UserName *string `json:"UserName,omitempty"`
}

// List the Me entities.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *MeEndpoint) List(ctx context.Context, all bool, o *api.ListOptions) ([]*Me, error) {
	var entities []*Me
	u, _ := s.client.ResolveURL("/api/v1/current/Me") // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
