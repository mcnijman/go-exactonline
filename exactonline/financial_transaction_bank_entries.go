// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// FinancialTransactionBankEntriesService is responsible for communicating with
// the BankEntries endpoint of the FinancialTransaction service.
type FinancialTransactionBankEntriesService service

// FinancialTransactionBankEntries:
// Service: FinancialTransaction
// Entity: BankEntries
// URL: /api/v1/{division}/financialtransaction/BankEntries
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=FinancialTransactionBankEntries
type FinancialTransactionBankEntries struct {
	// EntryID:
	EntryID *GUID `json:"EntryID,omitempty"`

	// BankEntryLines:
	BankEntryLines *[]byte `json:"BankEntryLines,omitempty"`

	// BankStatementDocument:
	BankStatementDocument *GUID `json:"BankStatementDocument,omitempty"`

	// BankStatementDocumentNumber:
	BankStatementDocumentNumber *int `json:"BankStatementDocumentNumber,omitempty"`

	// BankStatementDocumentSubject:
	BankStatementDocumentSubject *string `json:"BankStatementDocumentSubject,omitempty"`

	// ClosingBalanceFC:
	ClosingBalanceFC *float64 `json:"ClosingBalanceFC,omitempty"`

	// Created:
	Created *Date `json:"Created,omitempty"`

	// Currency:
	Currency *string `json:"Currency,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// EntryNumber:
	EntryNumber *int `json:"EntryNumber,omitempty"`

	// FinancialPeriod:
	FinancialPeriod *int `json:"FinancialPeriod,omitempty"`

	// FinancialYear:
	FinancialYear *int `json:"FinancialYear,omitempty"`

	// JournalCode:
	JournalCode *string `json:"JournalCode,omitempty"`

	// JournalDescription:
	JournalDescription *string `json:"JournalDescription,omitempty"`

	// Modified:
	Modified *Date `json:"Modified,omitempty"`

	// OpeningBalanceFC:
	OpeningBalanceFC *float64 `json:"OpeningBalanceFC,omitempty"`

	// Status:
	Status *int `json:"Status,omitempty"`

	// StatusDescription:
	StatusDescription *string `json:"StatusDescription,omitempty"`
}

func (s *FinancialTransactionBankEntries) GetIdentifier() GUID {
	return *s.EntryID
}

// List the BankEntries entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *FinancialTransactionBankEntriesService) List(ctx context.Context, division int, all bool) ([]*FinancialTransactionBankEntries, error) {
	var entities []*FinancialTransactionBankEntries
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/financialtransaction/BankEntries?$select=*", division)
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

/* // Get the BankEntries enitity, by EntryID.
func (s *FinancialTransactionBankEntriesService) Get(ctx context.Context, division int, id GUID) (*FinancialTransactionBankEntries, error) {
	var entities []*FinancialTransactionBankEntries
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/financialtransaction/BankEntries?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d BankEntries entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
