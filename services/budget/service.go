// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package budget

import "github.com/mcnijman/go-exactonline/api"

type service struct {
	client *api.Client
}

// BudgetService is responsible for communication with the Budget
// endpoints of the Exact Online API.
type BudgetService struct {
	client *api.Client

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Endpoints available under this service
	Budgets         *BudgetsEndpoint
	BudgetScenarios *BudgetScenariosEndpoint
}

// NewBudgetService creates a new initialized instance of the
// BudgetService.
func NewBudgetService(apiClient *api.Client) *BudgetService {
	s := &BudgetService{client: apiClient}

	s.common.client = apiClient

	s.Budgets = (*BudgetsEndpoint)(&s.common)
	s.BudgetScenarios = (*BudgetScenariosEndpoint)(&s.common)

	return s
}
