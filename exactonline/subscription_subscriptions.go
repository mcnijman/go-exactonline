// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// SubscriptionSubscriptionsService is responsible for communicating with
// the Subscriptions endpoint of the Subscription service.
type SubscriptionSubscriptionsService service

// SubscriptionSubscriptions:
// Service: Subscription
// Entity: Subscriptions
// URL: /api/v1/{division}/subscription/Subscriptions
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SubscriptionSubscriptions
type SubscriptionSubscriptions struct {
	// EntryID: Primary key
	EntryID *GUID `json:"EntryID,omitempty"`

	// BlockEntry: Indicates if subscription is blocked for time cost entry
	BlockEntry *bool `json:"BlockEntry,omitempty"`

	// CancellationDate: Date of cancellation
	CancellationDate *Date `json:"CancellationDate,omitempty"`

	// Classification: Reference to Classification
	Classification *GUID `json:"Classification,omitempty"`

	// ClassificationCode: Code of Classification
	ClassificationCode *string `json:"ClassificationCode,omitempty"`

	// ClassificationDescription: Description of Classification
	ClassificationDescription *string `json:"ClassificationDescription,omitempty"`

	// Created: Creation date
	Created *Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Currency: Currency code
	Currency *string `json:"Currency,omitempty"`

	// CustomerPONumber: Purchase order number of customer
	CustomerPONumber *string `json:"CustomerPONumber,omitempty"`

	// Description: Description
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// EndDate: End date
	EndDate *Date `json:"EndDate,omitempty"`

	// InvoiceDay: Invoice Day
	InvoiceDay *byte `json:"InvoiceDay,omitempty"`

	// InvoicedTo: Invoice date
	InvoicedTo *Date `json:"InvoicedTo,omitempty"`

	// InvoiceTo: Reference to invoice account
	InvoiceTo *GUID `json:"InvoiceTo,omitempty"`

	// InvoiceToContactPerson: Reference to contact person of invoice account
	InvoiceToContactPerson *GUID `json:"InvoiceToContactPerson,omitempty"`

	// InvoiceToContactPersonFullName: Name of contact person of invoice account
	InvoiceToContactPersonFullName *string `json:"InvoiceToContactPersonFullName,omitempty"`

	// InvoiceToName: Name of invoice account
	InvoiceToName *string `json:"InvoiceToName,omitempty"`

	// InvoicingStartDate: Invoicing start date
	InvoicingStartDate *Date `json:"InvoicingStartDate,omitempty"`

	// Modified: Last modified date
	Modified *Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Remarks
	Notes *string `json:"Notes,omitempty"`

	// Number: Number
	Number *int `json:"Number,omitempty"`

	// OrderedBy: Reference to order account
	OrderedBy *GUID `json:"OrderedBy,omitempty"`

	// OrderedByContactPerson: Reference of contact person of order account
	OrderedByContactPerson *GUID `json:"OrderedByContactPerson,omitempty"`

	// OrderedByContactPersonFullName: Name of contact person of order account
	OrderedByContactPersonFullName *string `json:"OrderedByContactPersonFullName,omitempty"`

	// OrderedByName: Name of order account
	OrderedByName *string `json:"OrderedByName,omitempty"`

	// PaymentCondition: Payment condition
	PaymentCondition *string `json:"PaymentCondition,omitempty"`

	// PaymentConditionDescription: Description of PaymentCondition
	PaymentConditionDescription *string `json:"PaymentConditionDescription,omitempty"`

	// Printed: Indicates if subscription is printed
	Printed *bool `json:"Printed,omitempty"`

	// ReasonCancelled: Reference to reason cancelled
	ReasonCancelled *GUID `json:"ReasonCancelled,omitempty"`

	// ReasonCancelledCode: Code of ReasonCancelled
	ReasonCancelledCode *string `json:"ReasonCancelledCode,omitempty"`

	// ReasonCancelledDescription: Description of ReasonCancelled
	ReasonCancelledDescription *string `json:"ReasonCancelledDescription,omitempty"`

	// StartDate: Start date
	StartDate *Date `json:"StartDate,omitempty"`

	// SubscriptionLines: Collection of subscription lines
	SubscriptionLines *[]byte `json:"SubscriptionLines,omitempty"`

	// SubscriptionRestrictionEmployees: Collection of restriction employees
	SubscriptionRestrictionEmployees *[]byte `json:"SubscriptionRestrictionEmployees,omitempty"`

	// SubscriptionRestrictionItems: Collection of restriction items
	SubscriptionRestrictionItems *[]byte `json:"SubscriptionRestrictionItems,omitempty"`

	// SubscriptionType: Reference to subscription type
	SubscriptionType *GUID `json:"SubscriptionType,omitempty"`

	// SubscriptionTypeCode: Code of SubscriptionType
	SubscriptionTypeCode *string `json:"SubscriptionTypeCode,omitempty"`

	// SubscriptionTypeDescription: Description of SubscriptionType
	SubscriptionTypeDescription *string `json:"SubscriptionTypeDescription,omitempty"`
}

func (s *SubscriptionSubscriptions) GetIdentifier() GUID {
	return *s.EntryID
}

// List the Subscriptions entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SubscriptionSubscriptionsService) List(ctx context.Context, division int, all bool) ([]*SubscriptionSubscriptions, error) {
	var entities []*SubscriptionSubscriptions
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/Subscriptions?$select=*", division)
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

/* // Get the Subscriptions enitity, by EntryID.
func (s *SubscriptionSubscriptionsService) Get(ctx context.Context, division int, id GUID) (*SubscriptionSubscriptions, error) {
	var entities []*SubscriptionSubscriptions
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/Subscriptions?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d Subscriptions entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
