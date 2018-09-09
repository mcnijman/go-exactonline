// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package financial

import (
	"context"

	"github.com/mcnijman/go-exactonline/api"
)

// ProfitLossOverviewEndpoint is responsible for communicating with
// the ProfitLossOverview endpoint of the Financial service.
type ProfitLossOverviewEndpoint service

// ProfitLossOverview:
// Service: Financial
// Entity: ProfitLossOverview
// URL: /api/v1/{division}/read/financial/ProfitLossOverview
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadFinancialProfitLossOverview
type ProfitLossOverview struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// CurrentYear: Primary key, Current year
	CurrentYear *int `json:"CurrentYear,omitempty"`

	// CostsCurrentPeriod: Costs in current period
	CostsCurrentPeriod *float64 `json:"CostsCurrentPeriod,omitempty"`

	// CostsCurrentYear: Costs in current year
	CostsCurrentYear *float64 `json:"CostsCurrentYear,omitempty"`

	// CostsPreviousYear: Costs in previous year
	CostsPreviousYear *float64 `json:"CostsPreviousYear,omitempty"`

	// CostsPreviousYearPeriod: Costs in period of previous year
	CostsPreviousYearPeriod *float64 `json:"CostsPreviousYearPeriod,omitempty"`

	// CurrencyCode: Currency code
	CurrencyCode *string `json:"CurrencyCode,omitempty"`

	// CurrentPeriod: Current period
	CurrentPeriod *int `json:"CurrentPeriod,omitempty"`

	// PreviousYear: Previous year
	PreviousYear *int `json:"PreviousYear,omitempty"`

	// PreviousYearPeriod: Period in previous year
	PreviousYearPeriod *int `json:"PreviousYearPeriod,omitempty"`

	// ResultCurrentPeriod: Results of current period
	ResultCurrentPeriod *float64 `json:"ResultCurrentPeriod,omitempty"`

	// ResultCurrentYear:
	ResultCurrentYear *float64 `json:"ResultCurrentYear,omitempty"`

	// ResultPreviousYear:
	ResultPreviousYear *float64 `json:"ResultPreviousYear,omitempty"`

	// ResultPreviousYearPeriod: Results of period in previous year
	ResultPreviousYearPeriod *float64 `json:"ResultPreviousYearPeriod,omitempty"`

	// RevenueCurrentPeriod: Revenue in current period
	RevenueCurrentPeriod *float64 `json:"RevenueCurrentPeriod,omitempty"`

	// RevenueCurrentYear: Revenue in current year
	RevenueCurrentYear *float64 `json:"RevenueCurrentYear,omitempty"`

	// RevenuePreviousYear: Revenue in previous year
	RevenuePreviousYear *float64 `json:"RevenuePreviousYear,omitempty"`

	// RevenuePreviousYearPeriod: Revenue in period of previous year
	RevenuePreviousYearPeriod *float64 `json:"RevenuePreviousYearPeriod,omitempty"`
}

func (e *ProfitLossOverview) GetPrimary() *int {
	return e.CurrentYear
}

func (s *ProfitLossOverviewEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "financial/ProfitLossOverview", method)
}

// List the ProfitLossOverview entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ProfitLossOverviewEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*ProfitLossOverview, error) {
	var entities []*ProfitLossOverview
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/ProfitLossOverview", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the ProfitLossOverview entitiy in the provided division.
func (s *ProfitLossOverviewEndpoint) Get(ctx context.Context, division int, id *int) (*ProfitLossOverview, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/ProfitLossOverview", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &ProfitLossOverview{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
