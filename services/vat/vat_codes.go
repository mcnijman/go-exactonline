// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package vat

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

// VATCodesEndpoint is responsible for communicating with
// the VATCodes endpoint of the VAT service.
type VATCodesEndpoint service

// VATCodes:
// Service: VAT
// Entity: VATCodes
// URL: /api/v1/{division}/vat/VATCodes
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=VATVATCodes
type VATCodes struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Account: Tax account
	Account *types.GUID `json:"Account,omitempty"`

	// AccountCode: Code of Account
	AccountCode *string `json:"AccountCode,omitempty"`

	// AccountName: Name of Account
	AccountName *string `json:"AccountName,omitempty"`

	// CalculationBasis: Indicates how to calculate the tax. 0 = based on the gross amount, 1 = based on the gross amount &#43; another tax
	CalculationBasis *byte `json:"CalculationBasis,omitempty"`

	// Charged: Indicates if transactions using the VAT code are transactions of the domestic VAT charging regulation (such as those for subcontractors) or transactions that are registered within the EU. If Charged=1 and linked to a purchase invoice, both a line for the VAT to pay and a line for the VAT to claim are being created
	Charged *bool `json:"Charged,omitempty"`

	// Code: VAT code
	Code *string `json:"Code,omitempty"`

	// Country: Obsolete
	Country *string `json:"Country,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description of the VAT code
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// EUSalesListing: Used in all legislations except France. Indicates if and how transactions using the VAT code appear on the ICT return (EU sales list). L = Listing goods, N = No listing, S = Listing services, T = Triangulation
	EUSalesListing *string `json:"EUSalesListing,omitempty"`

	// GLDiscountPurchase: Indicates the purchase discount GL account linked to the VAT codes for German legislation
	GLDiscountPurchase *types.GUID `json:"GLDiscountPurchase,omitempty"`

	// GLDiscountPurchaseCode: Code of GLDiscountPurchase
	GLDiscountPurchaseCode *string `json:"GLDiscountPurchaseCode,omitempty"`

	// GLDiscountPurchaseDescription: Description of GLDiscountPurchase
	GLDiscountPurchaseDescription *string `json:"GLDiscountPurchaseDescription,omitempty"`

	// GLDiscountSales: Indicates the sales discount GL account linked to the VAT codes for German legislation
	GLDiscountSales *types.GUID `json:"GLDiscountSales,omitempty"`

	// GLDiscountSalesCode: Code of GLDiscountSales
	GLDiscountSalesCode *string `json:"GLDiscountSalesCode,omitempty"`

	// GLDiscountSalesDescription: Description of GLDiscountSales
	GLDiscountSalesDescription *string `json:"GLDiscountSalesDescription,omitempty"`

	// GLToClaim: G/L account that is used to book the VAT to claim. If you enter purchases with a VAT code, the VAT amount to be claimed is entered to this VAT account. Must be of type VAT
	GLToClaim *types.GUID `json:"GLToClaim,omitempty"`

	// GLToClaimCode: Code of GLToClaim
	GLToClaimCode *string `json:"GLToClaimCode,omitempty"`

	// GLToClaimDescription: Description of GLToClaim
	GLToClaimDescription *string `json:"GLToClaimDescription,omitempty"`

	// GLToPay: G/L account that is used to book the VAT to pay. If you enter sales with a VAT code, the VAT amount to be paid is entered to this VAT account. Must be of type VAT
	GLToPay *types.GUID `json:"GLToPay,omitempty"`

	// GLToPayCode: Code of GLToPay
	GLToPayCode *string `json:"GLToPayCode,omitempty"`

	// GLToPayDescription: Description of GLToPay
	GLToPayDescription *string `json:"GLToPayDescription,omitempty"`

	// IntraStat: Used in all legislations except France. Indicates if intrastat is used
	IntraStat *bool `json:"IntraStat,omitempty"`

	// IntrastatType: Used in France legislation only. Indicates if and how transactions using the VAT code appear on the DEB/DES return. L = Goods, N = Empty, S = Services
	IntrastatType *string `json:"IntrastatType,omitempty"`

	// IsBlocked: Indicates if the VAT code may still be used
	IsBlocked *bool `json:"IsBlocked,omitempty"`

	// LegalText: Legal description for VAT code to print in the total block of the invoice
	LegalText *string `json:"LegalText,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: User name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Percentage: Percentage of the VAT code
	Percentage *float64 `json:"Percentage,omitempty"`

	// TaxReturnType: Indicates what type of Taxcode it is: can be VAT, IncomeTax
	TaxReturnType *int `json:"TaxReturnType,omitempty"`

	// Type: Indicates how the VAT amount should be calculated in relation to the invoice amount. B = VAT 0% (Only base amount), E = Excluding, I = Including, N = No VAT
	Type *string `json:"Type,omitempty"`

	// VatDocType: Field in VAT code maintenance to calculate different VATs depending on the selected document type. P = purchase invoice, F = freelance invoice, E = expense voucher. The field is valid for witholding tax type
	VatDocType *string `json:"VatDocType,omitempty"`

	// VatMargin: The VAT margin scheme is used for the trade of secondhand goods which are purchased without VAT (for example when a company buys a secondhand good from a private person). In the VAT margin scheme, the VAT is not calculated based on the sales price. Instead of that, the VAT is calculated based on the margin (gross sales price minus the gross purchase price)
	VatMargin *byte `json:"VatMargin,omitempty"`

	// VATPartialRatio: Partial ratio explains which part of the VAT the company has to pay. Used in some branches where the sellers have a bad reputation, so the buyers have to take over the VAT-liability
	VATPartialRatio *int `json:"VATPartialRatio,omitempty"`

	// VATPercentages: VAT percentages. You can have several VAT percentages, with start and end dates
	VATPercentages *json.RawMessage `json:"VATPercentages,omitempty"`

	// VATTransactionType: Indicates the type of transactions for which the VAT code may be used. B = Both, P = Purchase, S = Sales
	VATTransactionType *string `json:"VATTransactionType,omitempty"`
}

func (e *VATCodes) GetPrimary() *types.GUID {
	return e.ID
}

func (s *VATCodesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "vat/VATCodes", method)
}

// List the VATCodes entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *VATCodesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*VATCodes, error) {
	var entities []*VATCodes
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/vat/VATCodes", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the VATCodes entitiy in the provided division.
func (s *VATCodesEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*VATCodes, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/vat/VATCodes", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &VATCodes{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty VATCodes entity
func (s *VATCodesEndpoint) New() *VATCodes {
	return &VATCodes{}
}

// Create the VATCodes entity in the provided division.
func (s *VATCodesEndpoint) Create(ctx context.Context, division int, entity *VATCodes) (*VATCodes, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/vat/VATCodes", division) // #nosec
	e := &VATCodes{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the VATCodes entity in the provided division.
func (s *VATCodesEndpoint) Update(ctx context.Context, division int, entity *VATCodes) (*VATCodes, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/vat/VATCodes", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &VATCodes{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the VATCodes entity in the provided division.
func (s *VATCodesEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/vat/VATCodes", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return err
	}

	_, r, requestError := s.client.NewRequestAndDo(ctx, "DELETE", u.String(), nil, nil)
	if requestError != nil {
		return requestError
	}

	if r.StatusCode != http.StatusNoContent {
		body, _ := ioutil.ReadAll(r.Body) // #nosec
		return fmt.Errorf("Failed with status %v and body %v", r.StatusCode, body)
	}

	return nil
}
