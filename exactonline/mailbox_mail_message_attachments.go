// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// MailboxMailMessageAttachmentsService is responsible for communicating with
// the MailMessageAttachments endpoint of the Mailbox service.
type MailboxMailMessageAttachmentsService service

// MailboxMailMessageAttachments:
// Service: Mailbox
// Entity: MailMessageAttachments
// URL: /api/v1/{division}/mailbox/MailMessageAttachments
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=MailboxMailMessageAttachments
type MailboxMailMessageAttachments struct {
	// ID:
	ID *GUID `json:",omitempty"`

	// Attachment:
	Attachment *[]byte `json:",omitempty"`

	// AttachmentFileExtension:
	AttachmentFileExtension *string `json:",omitempty"`

	// AttachmentFileName:
	AttachmentFileName *string `json:",omitempty"`

	// FileSize:
	FileSize *int64 `json:",omitempty"`

	// MailMessageID:
	MailMessageID *GUID `json:",omitempty"`

	// RecipientAccount:
	RecipientAccount *GUID `json:",omitempty"`

	// SenderAccount:
	SenderAccount *GUID `json:",omitempty"`

	// Type:
	Type *int `json:",omitempty"`

	// TypeDescription:
	TypeDescription *string `json:",omitempty"`

	// Url:
	Url *string `json:",omitempty"`
}

func (s *MailboxMailMessageAttachments) GetIdentifier() GUID {
	return *s.ID
}

// List the MailMessageAttachments entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *MailboxMailMessageAttachmentsService) List(ctx context.Context, division int, all bool) ([]*MailboxMailMessageAttachments, error) {
	var entities []*MailboxMailMessageAttachments
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/mailbox/MailMessageAttachments?$select=*", division)
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