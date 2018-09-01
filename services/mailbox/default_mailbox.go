// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package mailbox

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// DefaultMailboxEndpoint is responsible for communicating with
// the DefaultMailbox endpoint of the Mailbox service.
type DefaultMailboxEndpoint service

// DefaultMailbox:
// Service: Mailbox
// Entity: DefaultMailbox
// URL: /api/v1/{division}/read/mailbox/DefaultMailbox
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadMailboxDefaultMailbox
type DefaultMailbox struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// Description: Extra description of the mailbox
	Description *string `json:"Description,omitempty"`

	// ForDivision: Only used when this mailbox is used for one specific administration, for example invoices to this mailbox will only be booked in this administration
	ForDivision *int `json:"ForDivision,omitempty"`

	// IsScanServiceMailbox: Indicates whether this service is used for messages returned by the scanning service
	IsScanServiceMailbox *bool `json:"IsScanServiceMailbox,omitempty"`

	// Mailbox: E-mail address-like format, for example johndoe@exactonline.nl
	Mailbox *string `json:"Mailbox,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ValidFrom: Date that this mailbox became valid
	ValidFrom *types.Date `json:"ValidFrom,omitempty"`

	// ValidTo: Date that this mailbox will not be valid anymore
	ValidTo *types.Date `json:"ValidTo,omitempty"`
}

func (s *DefaultMailbox) GetIdentifier() types.GUID {
	return *s.ID
}

// List the DefaultMailbox entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *DefaultMailboxEndpoint) List(ctx context.Context, division int, all bool) ([]*DefaultMailbox, error) {
	var entities []*DefaultMailbox
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/read/mailbox/DefaultMailbox?$select=*", division)
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
