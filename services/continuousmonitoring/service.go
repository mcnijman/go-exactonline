// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package continuousmonitoring

import "github.com/mcnijman/go-exactonline/api"

type service struct {
	client *api.Client
}

// ContinuousMonitoringService is responsible for communication with the ContinuousMonitoring
// endpoints of the Exact Online API.
type ContinuousMonitoringService struct {
	client *api.Client

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Endpoints available under this service
	IndicatorBalances                *IndicatorBalancesEndpoint
	IndicatorDeviatingAmountEntereds *IndicatorDeviatingAmountEnteredsEndpoint
	IndicatorDifferenceByPeriods     *IndicatorDifferenceByPeriodsEndpoint
	IndicatorDifferentVatCodes       *IndicatorDifferentVatCodesEndpoint
	IndicatorGLAccounts              *IndicatorGLAccountsEndpoint
	IndicatorLiquidities             *IndicatorLiquiditiesEndpoint
	IndicatorSignals                 *IndicatorSignalsEndpoint
	IndicatorStates                  *IndicatorStatesEndpoint
	IndicatorUsageOfJournals         *IndicatorUsageOfJournalsEndpoint
}

// NewContinuousMonitoringService creates a new initialized instance of the
// ContinuousMonitoringService.
func NewContinuousMonitoringService(apiClient *api.Client) *ContinuousMonitoringService {
	s := &ContinuousMonitoringService{client: apiClient}

	s.common.client = apiClient

	s.IndicatorBalances = (*IndicatorBalancesEndpoint)(&s.common)
	s.IndicatorDeviatingAmountEntereds = (*IndicatorDeviatingAmountEnteredsEndpoint)(&s.common)
	s.IndicatorDifferenceByPeriods = (*IndicatorDifferenceByPeriodsEndpoint)(&s.common)
	s.IndicatorDifferentVatCodes = (*IndicatorDifferentVatCodesEndpoint)(&s.common)
	s.IndicatorGLAccounts = (*IndicatorGLAccountsEndpoint)(&s.common)
	s.IndicatorLiquidities = (*IndicatorLiquiditiesEndpoint)(&s.common)
	s.IndicatorSignals = (*IndicatorSignalsEndpoint)(&s.common)
	s.IndicatorStates = (*IndicatorStatesEndpoint)(&s.common)
	s.IndicatorUsageOfJournals = (*IndicatorUsageOfJournalsEndpoint)(&s.common)

	return s
}
