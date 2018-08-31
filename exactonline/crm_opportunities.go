// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// CRMOpportunitiesService is responsible for communicating with
// the Opportunities endpoint of the CRM service.
type CRMOpportunitiesService service

// CRMOpportunities:
// Service: CRM
// Entity: Opportunities
// URL: /api/v1/{division}/crm/Opportunities
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=CRMOpportunities
type CRMOpportunities struct {
	// ID:
	ID *GUID `json:"ID,omitempty"`

	// Account:
	Account *GUID `json:"Account,omitempty"`

	// Accountant:
	Accountant *GUID `json:"Accountant,omitempty"`

	// AccountantCode:
	AccountantCode *string `json:"AccountantCode,omitempty"`

	// AccountantName:
	AccountantName *string `json:"AccountantName,omitempty"`

	// AccountCode:
	AccountCode *string `json:"AccountCode,omitempty"`

	// AccountName:
	AccountName *string `json:"AccountName,omitempty"`

	// ActionDate:
	ActionDate *Date `json:"ActionDate,omitempty"`

	// AmountDC:
	AmountDC *float64 `json:"AmountDC,omitempty"`

	// AmountFC:
	AmountFC *float64 `json:"AmountFC,omitempty"`

	// Campaign:
	Campaign *GUID `json:"Campaign,omitempty"`

	// CampaignDescription:
	CampaignDescription *string `json:"CampaignDescription,omitempty"`

	// Channel:
	Channel *int `json:"Channel,omitempty"`

	// ChannelDescription:
	ChannelDescription *string `json:"ChannelDescription,omitempty"`

	// CloseDate:
	CloseDate *Date `json:"CloseDate,omitempty"`

	// Created:
	Created *Date `json:"Created,omitempty"`

	// Creator:
	Creator *GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Currency:
	Currency *string `json:"Currency,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// LeadSource:
	LeadSource *GUID `json:"LeadSource,omitempty"`

	// LeadSourceDescription:
	LeadSourceDescription *string `json:"LeadSourceDescription,omitempty"`

	// Modified:
	Modified *Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Name:
	Name *string `json:"Name,omitempty"`

	// NextAction:
	NextAction *string `json:"NextAction,omitempty"`

	// Notes:
	Notes *string `json:"Notes,omitempty"`

	// OpportunityDepartmentCode:
	OpportunityDepartmentCode *int `json:"OpportunityDepartmentCode,omitempty"`

	// OpportunityDepartmentDescription:
	OpportunityDepartmentDescription *string `json:"OpportunityDepartmentDescription,omitempty"`

	// OpportunityStage:
	OpportunityStage *GUID `json:"OpportunityStage,omitempty"`

	// OpportunityStageDescription:
	OpportunityStageDescription *string `json:"OpportunityStageDescription,omitempty"`

	// OpportunityStatus:
	OpportunityStatus *int `json:"OpportunityStatus,omitempty"`

	// OpportunityType:
	OpportunityType *int `json:"OpportunityType,omitempty"`

	// OpportunityTypeDescription:
	OpportunityTypeDescription *string `json:"OpportunityTypeDescription,omitempty"`

	// Owner:
	Owner *GUID `json:"Owner,omitempty"`

	// OwnerFullName:
	OwnerFullName *string `json:"OwnerFullName,omitempty"`

	// Probability:
	Probability *float64 `json:"Probability,omitempty"`

	// Project:
	Project *GUID `json:"Project,omitempty"`

	// ProjectCode:
	ProjectCode *string `json:"ProjectCode,omitempty"`

	// ProjectDescription:
	ProjectDescription *string `json:"ProjectDescription,omitempty"`

	// RateFC:
	RateFC *float64 `json:"RateFC,omitempty"`

	// ReasonCode:
	ReasonCode *GUID `json:"ReasonCode,omitempty"`

	// ReasonCodeDescription:
	ReasonCodeDescription *string `json:"ReasonCodeDescription,omitempty"`

	// Reseller:
	Reseller *GUID `json:"Reseller,omitempty"`

	// ResellerCode:
	ResellerCode *string `json:"ResellerCode,omitempty"`

	// ResellerName:
	ResellerName *string `json:"ResellerName,omitempty"`

	// SalesType:
	SalesType *GUID `json:"SalesType,omitempty"`

	// SalesTypeDescription:
	SalesTypeDescription *string `json:"SalesTypeDescription,omitempty"`
}

func (s *CRMOpportunities) GetIdentifier() GUID {
	return *s.ID
}

// List the Opportunities entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *CRMOpportunitiesService) List(ctx context.Context, division int, all bool) ([]*CRMOpportunities, error) {
	var entities []*CRMOpportunities
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/Opportunities?$select=*", division)
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

/* // Get the Opportunities enitity, by ID.
func (s *CRMOpportunitiesService) Get(ctx context.Context, division int, id GUID) (*CRMOpportunities, error) {
	var entities []*CRMOpportunities
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/Opportunities?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d Opportunities entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
