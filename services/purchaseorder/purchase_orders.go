// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package purchaseorder

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// PurchaseOrdersEndpoint is responsible for communicating with
// the PurchaseOrders endpoint of the PurchaseOrder service.
type PurchaseOrdersEndpoint service

// PurchaseOrders:
// Service: PurchaseOrder
// Entity: PurchaseOrders
// URL: /api/v1/{division}/purchaseorder/PurchaseOrders
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=PurchaseOrderPurchaseOrders
type PurchaseOrders struct {
	// PurchaseOrderID:
	PurchaseOrderID *types.GUID `json:"PurchaseOrderID,omitempty"`

	// AmountDC:
	AmountDC *float64 `json:"AmountDC,omitempty"`

	// AmountFC:
	AmountFC *float64 `json:"AmountFC,omitempty"`

	// Created:
	Created *types.Date `json:"Created,omitempty"`

	// Creator:
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Currency:
	Currency *string `json:"Currency,omitempty"`

	// DeliveryAccount:
	DeliveryAccount *types.GUID `json:"DeliveryAccount,omitempty"`

	// DeliveryAccountCode:
	DeliveryAccountCode *string `json:"DeliveryAccountCode,omitempty"`

	// DeliveryAccountName:
	DeliveryAccountName *string `json:"DeliveryAccountName,omitempty"`

	// DeliveryAddress:
	DeliveryAddress *types.GUID `json:"DeliveryAddress,omitempty"`

	// DeliveryContact:
	DeliveryContact *types.GUID `json:"DeliveryContact,omitempty"`

	// DeliveryContactPersonFullName:
	DeliveryContactPersonFullName *string `json:"DeliveryContactPersonFullName,omitempty"`

	// Description:
	Description *string `json:"Description,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// Document:
	Document *types.GUID `json:"Document,omitempty"`

	// DocumentSubject:
	DocumentSubject *string `json:"DocumentSubject,omitempty"`

	// DropShipment:
	DropShipment *bool `json:"DropShipment,omitempty"`

	// ExchangeRate:
	ExchangeRate *float64 `json:"ExchangeRate,omitempty"`

	// InvoiceStatus:
	InvoiceStatus *int `json:"InvoiceStatus,omitempty"`

	// Modified:
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// OrderDate:
	OrderDate *types.Date `json:"OrderDate,omitempty"`

	// OrderNumber:
	OrderNumber *int `json:"OrderNumber,omitempty"`

	// OrderStatus:
	OrderStatus *int `json:"OrderStatus,omitempty"`

	// PaymentCondition:
	PaymentCondition *string `json:"PaymentCondition,omitempty"`

	// PaymentConditionDescription:
	PaymentConditionDescription *string `json:"PaymentConditionDescription,omitempty"`

	// PurchaseAgent:
	PurchaseAgent *types.GUID `json:"PurchaseAgent,omitempty"`

	// PurchaseAgentFullName:
	PurchaseAgentFullName *string `json:"PurchaseAgentFullName,omitempty"`

	// PurchaseOrderLines:
	PurchaseOrderLines *[]byte `json:"PurchaseOrderLines,omitempty"`

	// ReceiptDate:
	ReceiptDate *types.Date `json:"ReceiptDate,omitempty"`

	// ReceiptStatus:
	ReceiptStatus *int `json:"ReceiptStatus,omitempty"`

	// Remarks:
	Remarks *string `json:"Remarks,omitempty"`

	// SalesOrder:
	SalesOrder *types.GUID `json:"SalesOrder,omitempty"`

	// SalesOrderNumber:
	SalesOrderNumber *int `json:"SalesOrderNumber,omitempty"`

	// ShippingMethod:
	ShippingMethod *types.GUID `json:"ShippingMethod,omitempty"`

	// ShippingMethodDescription:
	ShippingMethodDescription *string `json:"ShippingMethodDescription,omitempty"`

	// Source:
	Source *int `json:"Source,omitempty"`

	// Supplier:
	Supplier *types.GUID `json:"Supplier,omitempty"`

	// SupplierCode:
	SupplierCode *string `json:"SupplierCode,omitempty"`

	// SupplierContact:
	SupplierContact *types.GUID `json:"SupplierContact,omitempty"`

	// SupplierContactPersonFullName:
	SupplierContactPersonFullName *string `json:"SupplierContactPersonFullName,omitempty"`

	// SupplierName:
	SupplierName *string `json:"SupplierName,omitempty"`

	// VATAmount:
	VATAmount *float64 `json:"VATAmount,omitempty"`

	// Warehouse:
	Warehouse *types.GUID `json:"Warehouse,omitempty"`

	// WarehouseCode:
	WarehouseCode *string `json:"WarehouseCode,omitempty"`

	// WarehouseDescription:
	WarehouseDescription *string `json:"WarehouseDescription,omitempty"`

	// YourRef:
	YourRef *string `json:"YourRef,omitempty"`
}

func (s *PurchaseOrders) GetIdentifier() types.GUID {
	return *s.PurchaseOrderID
}

// List the PurchaseOrders entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *PurchaseOrdersEndpoint) List(ctx context.Context, division int, all bool) ([]*PurchaseOrders, error) {
	var entities []*PurchaseOrders
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/purchaseorder/PurchaseOrders?$select=*", division)
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
