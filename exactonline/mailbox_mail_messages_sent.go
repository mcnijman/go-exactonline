// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// MailboxMailMessagesSentService is responsible for communicating with
// the MailMessagesSent endpoint of the Mailbox service.
type MailboxMailMessagesSentService service

// MailboxMailMessagesSent:
// Service: Mailbox
// Entity: MailMessagesSent
// URL: /api/v1/{division}/mailbox/MailMessagesSent
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=MailboxMailMessagesSent
type MailboxMailMessagesSent struct {
	// ID:
	ID *GUID `json:"ID,omitempty"`

	// Bank:
	Bank *GUID `json:"Bank,omitempty"`

	// BankAccount:
	BankAccount *string `json:"BankAccount,omitempty"`

	// Created:
	Created *Date `json:"Created,omitempty"`

	// Creator:
	Creator *GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// ForDivision:
	ForDivision *int `json:"ForDivision,omitempty"`

	// Modified:
	Modified *Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Operation:
	Operation *int `json:"Operation,omitempty"`

	// OriginalMessage:
	OriginalMessage *GUID `json:"OriginalMessage,omitempty"`

	// OriginalMessageSubject:
	OriginalMessageSubject *string `json:"OriginalMessageSubject,omitempty"`

	// PartnerKey:
	PartnerKey *GUID `json:"PartnerKey,omitempty"`

	// Quantity:
	Quantity *float64 `json:"Quantity,omitempty"`

	// RecipientAccount:
	RecipientAccount *GUID `json:"RecipientAccount,omitempty"`

	// RecipientDeleted:
	RecipientDeleted *byte `json:"RecipientDeleted,omitempty"`

	// RecipientMailbox:
	RecipientMailbox *string `json:"RecipientMailbox,omitempty"`

	// RecipientMailboxDescription:
	RecipientMailboxDescription *string `json:"RecipientMailboxDescription,omitempty"`

	// RecipientMailboxID:
	RecipientMailboxID *GUID `json:"RecipientMailboxID,omitempty"`

	// RecipientStatus:
	RecipientStatus *int `json:"RecipientStatus,omitempty"`

	// RecipientStatusDescription:
	RecipientStatusDescription *string `json:"RecipientStatusDescription,omitempty"`

	// SenderAccount:
	SenderAccount *GUID `json:"SenderAccount,omitempty"`

	// SenderDateSent:
	SenderDateSent *Date `json:"SenderDateSent,omitempty"`

	// SenderDeleted:
	SenderDeleted *byte `json:"SenderDeleted,omitempty"`

	// SenderIPAddress:
	SenderIPAddress *string `json:"SenderIPAddress,omitempty"`

	// SenderMailbox:
	SenderMailbox *string `json:"SenderMailbox,omitempty"`

	// SenderMailboxDescription:
	SenderMailboxDescription *string `json:"SenderMailboxDescription,omitempty"`

	// SenderMailboxID:
	SenderMailboxID *GUID `json:"SenderMailboxID,omitempty"`

	// Subject:
	Subject *string `json:"Subject,omitempty"`

	// SynchronizationCode:
	SynchronizationCode *string `json:"SynchronizationCode,omitempty"`

	// Type:
	Type *int `json:"Type,omitempty"`
}

func (s *MailboxMailMessagesSent) GetIdentifier() GUID {
	return *s.ID
}

// List the MailMessagesSent entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *MailboxMailMessagesSentService) List(ctx context.Context, division int, all bool) ([]*MailboxMailMessagesSent, error) {
	var entities []*MailboxMailMessagesSent
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/mailbox/MailMessagesSent?$select=*", division)
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

/* // Get the MailMessagesSent enitity, by ID.
func (s *MailboxMailMessagesSentService) Get(ctx context.Context, division int, id GUID) (*MailboxMailMessagesSent, error) {
	var entities []*MailboxMailMessagesSent
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/mailbox/MailMessagesSent?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d MailMessagesSent entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
