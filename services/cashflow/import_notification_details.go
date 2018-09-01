// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package cashflow

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// ImportNotificationDetailsEndpoint is responsible for communicating with
// the ImportNotificationDetails endpoint of the Cashflow service.
type ImportNotificationDetailsEndpoint service

// ImportNotificationDetails:
// Service: Cashflow
// Entity: ImportNotificationDetails
// URL: /api/v1/{division}/cashflow/ImportNotificationDetails
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=CashflowImportNotificationDetails
type ImportNotificationDetails struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// CashflowImportNotification: ID of the notification these details belong to.
	CashflowImportNotification *types.GUID `json:"CashflowImportNotification,omitempty"`

	// CashflowTransactionFeed: ID of the cashflow transaction feed related to this notification.
	CashflowTransactionFeed *types.GUID `json:"CashflowTransactionFeed,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Message: Termed message
	Message *string `json:"Message,omitempty"`

	// ResponseCode: Response code
	ResponseCode *int `json:"ResponseCode,omitempty"`

	// ResponseCodeArguments: Additional information about the response
	ResponseCodeArguments *string `json:"ResponseCodeArguments,omitempty"`
}

func (s *ImportNotificationDetails) GetIdentifier() types.GUID {
	return *s.ID
}

// List the ImportNotificationDetails entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ImportNotificationDetailsEndpoint) List(ctx context.Context, division int, all bool) ([]*ImportNotificationDetails, error) {
	var entities []*ImportNotificationDetails
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/cashflow/ImportNotificationDetails?$select=*", division)
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
