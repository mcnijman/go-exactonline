// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

//go:generate go run gen-services.go -v

package exactonline

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/services/accountancy"
	"github.com/mcnijman/go-exactonline/services/activities"
	"github.com/mcnijman/go-exactonline/services/assets"
	"github.com/mcnijman/go-exactonline/services/budget"
	"github.com/mcnijman/go-exactonline/services/bulk"
	"github.com/mcnijman/go-exactonline/services/cashflow"
	"github.com/mcnijman/go-exactonline/services/continuousmonitoring"
	"github.com/mcnijman/go-exactonline/services/crm"
	"github.com/mcnijman/go-exactonline/services/documents"
	"github.com/mcnijman/go-exactonline/services/financial"
	"github.com/mcnijman/go-exactonline/services/financialtransaction"
	"github.com/mcnijman/go-exactonline/services/general"
	"github.com/mcnijman/go-exactonline/services/generaljournalentry"
	"github.com/mcnijman/go-exactonline/services/hrm"
	"github.com/mcnijman/go-exactonline/services/inventory"
	"github.com/mcnijman/go-exactonline/services/logistics"
	"github.com/mcnijman/go-exactonline/services/mailbox"
	"github.com/mcnijman/go-exactonline/services/manufacturing"
	"github.com/mcnijman/go-exactonline/services/openingbalance"
	"github.com/mcnijman/go-exactonline/services/payroll"
	"github.com/mcnijman/go-exactonline/services/project"
	"github.com/mcnijman/go-exactonline/services/purchase"
	"github.com/mcnijman/go-exactonline/services/purchaseentry"
	"github.com/mcnijman/go-exactonline/services/purchaseorder"
	"github.com/mcnijman/go-exactonline/services/sales"
	"github.com/mcnijman/go-exactonline/services/salesentry"
	"github.com/mcnijman/go-exactonline/services/salesinvoice"
	"github.com/mcnijman/go-exactonline/services/salesorder"
	"github.com/mcnijman/go-exactonline/services/subscription"
	"github.com/mcnijman/go-exactonline/services/system"
	"github.com/mcnijman/go-exactonline/services/users"
	"github.com/mcnijman/go-exactonline/services/vat"
	"github.com/mcnijman/go-exactonline/services/webhooks"
	"github.com/mcnijman/go-exactonline/services/workflow"
	"github.com/mcnijman/go-exactonline/types"
	"golang.org/x/oauth2"
)

// A Client manages communication with the Exact Online API.
type Client struct {
	client *api.Client

	// Services used for talking to different parts of the Exact Online API
	Budget               *budget.BudgetService
	Bulk                 *bulk.BulkService
	ContinuousMonitoring *continuousmonitoring.ContinuousMonitoringService
	Documents            *documents.DocumentsService
	FinancialTransaction *financialtransaction.FinancialTransactionService
	General              *general.GeneralService
	Inventory            *inventory.InventoryService
	Accountancy          *accountancy.AccountancyService
	Users                *users.UsersService
	VAT                  *vat.VATService
	Workflow             *workflow.WorkflowService
	PurchaseEntry        *purchaseentry.PurchaseEntryService
	Payroll              *payroll.PayrollService
	Purchase             *purchase.PurchaseService
	SalesOrder           *salesorder.SalesOrderService
	Logistics            *logistics.LogisticsService
	CRM                  *crm.CRMService
	GeneralJournalEntry  *generaljournalentry.GeneralJournalEntryService
	OpeningBalance       *openingbalance.OpeningBalanceService
	Project              *project.ProjectService
	Webhooks             *webhooks.WebhooksService
	Cashflow             *cashflow.CashflowService
	SalesInvoice         *salesinvoice.SalesInvoiceService
	PurchaseOrder        *purchaseorder.PurchaseOrderService
	Sales                *sales.SalesService
	SalesEntry           *salesentry.SalesEntryService
	Mailbox              *mailbox.MailboxService
	Assets               *assets.AssetsService
	Financial            *financial.FinancialService
	HRM                  *hrm.HRMService
	Manufacturing        *manufacturing.ManufacturingService
	Subscription         *subscription.SubscriptionService
	System               *system.SystemService
	Activities           *activities.ActivitiesService
}

// NewClient returns a new Exact Online API client. Provide a http.Client that
// will perform the authentication for you (such as that provided by the
// golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client) *Client {
	restClient := api.NewClient(httpClient)

	c := &Client{client: restClient}

	// Services setup
	c.HRM = hrm.NewHRMService(c.client)
	c.VAT = vat.NewVATService(c.client)
	c.Webhooks = webhooks.NewWebhooksService(c.client)
	c.Activities = activities.NewActivitiesService(c.client)
	c.Documents = documents.NewDocumentsService(c.client)
	c.Payroll = payroll.NewPayrollService(c.client)
	c.Bulk = bulk.NewBulkService(c.client)
	c.FinancialTransaction = financialtransaction.NewFinancialTransactionService(c.client)
	c.GeneralJournalEntry = generaljournalentry.NewGeneralJournalEntryService(c.client)
	c.Manufacturing = manufacturing.NewManufacturingService(c.client)
	c.Workflow = workflow.NewWorkflowService(c.client)
	c.ContinuousMonitoring = continuousmonitoring.NewContinuousMonitoringService(c.client)
	c.Financial = financial.NewFinancialService(c.client)
	c.OpeningBalance = openingbalance.NewOpeningBalanceService(c.client)
	c.Project = project.NewProjectService(c.client)
	c.PurchaseEntry = purchaseentry.NewPurchaseEntryService(c.client)
	c.PurchaseOrder = purchaseorder.NewPurchaseOrderService(c.client)
	c.Subscription = subscription.NewSubscriptionService(c.client)
	c.Cashflow = cashflow.NewCashflowService(c.client)
	c.Mailbox = mailbox.NewMailboxService(c.client)
	c.Purchase = purchase.NewPurchaseService(c.client)
	c.Sales = sales.NewSalesService(c.client)
	c.System = system.NewSystemService(c.client)
	c.CRM = crm.NewCRMService(c.client)
	c.Logistics = logistics.NewLogisticsService(c.client)
	c.General = general.NewGeneralService(c.client)
	c.SalesInvoice = salesinvoice.NewSalesInvoiceService(c.client)
	c.SalesOrder = salesorder.NewSalesOrderService(c.client)
	c.Users = users.NewUsersService(c.client)
	c.Inventory = inventory.NewInventoryService(c.client)
	c.SalesEntry = salesentry.NewSalesEntryService(c.client)
	c.Budget = budget.NewBudgetService(c.client)
	c.Accountancy = accountancy.NewAccountancyService(c.client)
	c.Assets = assets.NewAssetsService(c.client)

	return c
}

// NewClientFromTokenSource is a wrapper around NewClient if you have a valid
// token source. It will create a http.Client from the oauth2.Tokensource.
// If no context is available you can use context.Background()
func NewClientFromTokenSource(ctx context.Context, tokenSource oauth2.TokenSource) *Client {
	httpClient := oauth2.NewClient(ctx, tokenSource)
	return NewClient(httpClient)
}

// SetBaseURL sets the base URL for communicating with the Exact Online API.
// If the URL does not have a trailing slash, one is added automatically.
// For each country, the Exact Online solution is deployed on a separate site.
// Because of this, the Exact Online server URL is country dependent.
// The Exact Online server URLs are:
//     - The Netherlands: https://start.exactonline.nl (default)
//     - Belgium: https://start.exactonline.be
//     - Germany: https://start.exactonline.de
//     - United Kingdom: https://start.exactonline.co.uk
//     - United States of America: https://start.exactonline.com
//     - Spain: https://start.exactonline.es
// Docs: https://support.exactonline.com/community/s/knowledge-base#All-All-DNO-Content-exact-online-sites
func (c *Client) SetBaseURL(baseURL string) error {
	baseEndpoint, err := url.Parse(baseURL)
	if err != nil {
		return err
	}
	if !strings.HasSuffix(baseEndpoint.Path, "/") {
		baseEndpoint.Path += "/"
	}
	c.client.BaseURL = baseEndpoint

	return nil
}

// SetUserAgent sets the useragent provided on every communication with the Exact Online API.
func (c *Client) SetUserAgent(userAgent string) {
	c.client.UserAgent = userAgent
}

// GetCurrentDivisionID fetches the last used division id of the user
func (c *Client) GetCurrentDivisionID(ctx context.Context) (int, error) {
	mes, err := c.System.Me.List(ctx, false)
	if err != nil {
		return 0, err
	}
	if len(mes) != 1 {
		return 0, fmt.Errorf("System.Me response is supposed to have 1 entity, got %d", len(mes))
	}
	return *mes[0].CurrentDivision, nil
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Date is a helper routine that allocates a new types.Date value
// to store v and returns a pointer to it.
func Date(v time.Time) *types.Date { return &types.Date{Time: v} }

// Float64 is a helper routine that allocates a new float64 value
// to store v and returns a pointer to it.
func Float64(v float64) *float64 { return &v }

// GUID is a helper routine that allocates a new types.GUID value
// to store v and returns a pointer to it.
func GUID(v uuid.UUID) *types.GUID { return &types.GUID{UUID: v} }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// URL is a helper routine that allocates a new types.URL value
// to store v and returns a pointer to it.
func URL(v *url.URL) *types.URL { return &types.URL{URL: v} }
