// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package exactonline

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
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

// setup sets up a test HTTP server along with a exactonline.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	// We want to ensure that tests catch mistakes where the endpoint URL is
	// specified as absolute rather than relative. It only makes a difference
	// when there's a non-empty base URL path. So, use that. See issue #752.
	apiHandler := http.NewServeMux()
	apiHandler.Handle("/", mux)

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(apiHandler)

	// client is the GitHub client being tested and is
	// configured to use test server.
	client = NewClient(nil)
	client.SetBaseURL(server.URL + "/")
	return client, mux, server.URL, server.Close
}

func TestNewClient(t *testing.T) {
	type args struct {
		httpClient *http.Client
	}
	a := api.NewClient(nil)
	b := &http.Client{Timeout: 1 * time.Millisecond}
	c := api.NewClient(b)
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{"default", args{http.DefaultClient}, &Client{
			client:               api.NewClient(nil),
			HRM:                  hrm.NewHRMService(a),
			VAT:                  vat.NewVATService(a),
			Webhooks:             webhooks.NewWebhooksService(a),
			Activities:           activities.NewActivitiesService(a),
			Documents:            documents.NewDocumentsService(a),
			Payroll:              payroll.NewPayrollService(a),
			Bulk:                 bulk.NewBulkService(a),
			FinancialTransaction: financialtransaction.NewFinancialTransactionService(a),
			GeneralJournalEntry:  generaljournalentry.NewGeneralJournalEntryService(a),
			Manufacturing:        manufacturing.NewManufacturingService(a),
			Workflow:             workflow.NewWorkflowService(a),
			ContinuousMonitoring: continuousmonitoring.NewContinuousMonitoringService(a),
			Financial:            financial.NewFinancialService(a),
			OpeningBalance:       openingbalance.NewOpeningBalanceService(a),
			Project:              project.NewProjectService(a),
			PurchaseEntry:        purchaseentry.NewPurchaseEntryService(a),
			PurchaseOrder:        purchaseorder.NewPurchaseOrderService(a),
			Subscription:         subscription.NewSubscriptionService(a),
			Cashflow:             cashflow.NewCashflowService(a),
			Mailbox:              mailbox.NewMailboxService(a),
			Purchase:             purchase.NewPurchaseService(a),
			Sales:                sales.NewSalesService(a),
			System:               system.NewSystemService(a),
			CRM:                  crm.NewCRMService(a),
			Logistics:            logistics.NewLogisticsService(a),
			General:              general.NewGeneralService(a),
			SalesInvoice:         salesinvoice.NewSalesInvoiceService(a),
			SalesOrder:           salesorder.NewSalesOrderService(a),
			Users:                users.NewUsersService(a),
			Inventory:            inventory.NewInventoryService(a),
			SalesEntry:           salesentry.NewSalesEntryService(a),
			Budget:               budget.NewBudgetService(a),
			Accountancy:          accountancy.NewAccountancyService(a),
			Assets:               assets.NewAssetsService(a),
		}},
		{"custom", args{b}, &Client{
			client:               api.NewClient(b),
			HRM:                  hrm.NewHRMService(c),
			VAT:                  vat.NewVATService(c),
			Webhooks:             webhooks.NewWebhooksService(c),
			Activities:           activities.NewActivitiesService(c),
			Documents:            documents.NewDocumentsService(c),
			Payroll:              payroll.NewPayrollService(c),
			Bulk:                 bulk.NewBulkService(c),
			FinancialTransaction: financialtransaction.NewFinancialTransactionService(c),
			GeneralJournalEntry:  generaljournalentry.NewGeneralJournalEntryService(c),
			Manufacturing:        manufacturing.NewManufacturingService(c),
			Workflow:             workflow.NewWorkflowService(c),
			ContinuousMonitoring: continuousmonitoring.NewContinuousMonitoringService(c),
			Financial:            financial.NewFinancialService(c),
			OpeningBalance:       openingbalance.NewOpeningBalanceService(c),
			Project:              project.NewProjectService(c),
			PurchaseEntry:        purchaseentry.NewPurchaseEntryService(c),
			PurchaseOrder:        purchaseorder.NewPurchaseOrderService(c),
			Subscription:         subscription.NewSubscriptionService(c),
			Cashflow:             cashflow.NewCashflowService(c),
			Mailbox:              mailbox.NewMailboxService(c),
			Purchase:             purchase.NewPurchaseService(c),
			Sales:                sales.NewSalesService(c),
			System:               system.NewSystemService(c),
			CRM:                  crm.NewCRMService(c),
			Logistics:            logistics.NewLogisticsService(c),
			General:              general.NewGeneralService(c),
			SalesInvoice:         salesinvoice.NewSalesInvoiceService(c),
			SalesOrder:           salesorder.NewSalesOrderService(c),
			Users:                users.NewUsersService(c),
			Inventory:            inventory.NewInventoryService(c),
			SalesEntry:           salesentry.NewSalesEntryService(c),
			Budget:               budget.NewBudgetService(c),
			Accountancy:          accountancy.NewAccountancyService(c),
			Assets:               assets.NewAssetsService(c),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.httpClient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClientFromTokenSource(t *testing.T) {
	type args struct {
		ctx         context.Context
		tokenSource oauth2.TokenSource
	}
	ctx := context.Background()
	tok := oauth2.StaticTokenSource(&oauth2.Token{})
	c := oauth2.NewClient(ctx, tok)
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{"1", args{ctx, tok}, NewClient(c)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClientFromTokenSource(tt.args.ctx, tt.args.tokenSource); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClientFromTokenSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SetBaseURL(t *testing.T) {
	type fields struct {
		client               *api.Client
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
	type args struct {
		baseURL string
		u       *url.URL
	}
	c := api.NewClient(nil)
	uri := "https://start.exactonline.nl/"
	u, _ := url.Parse(uri)
	uri2 := "https://start.exactonline.nl"
	u2, _ := url.Parse(uri2)
	uri3 := ":foo"
	u3, _ := url.Parse(uri2)
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		wantEqual bool
	}{
		{"1", fields{client: c}, args{uri, u}, false, true},
		{"2", fields{client: c}, args{uri2, u2}, false, false},
		{"3", fields{client: c}, args{uri3, u3}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:               tt.fields.client,
				Budget:               tt.fields.Budget,
				Bulk:                 tt.fields.Bulk,
				ContinuousMonitoring: tt.fields.ContinuousMonitoring,
				Documents:            tt.fields.Documents,
				FinancialTransaction: tt.fields.FinancialTransaction,
				General:              tt.fields.General,
				Inventory:            tt.fields.Inventory,
				Accountancy:          tt.fields.Accountancy,
				Users:                tt.fields.Users,
				VAT:                  tt.fields.VAT,
				Workflow:             tt.fields.Workflow,
				PurchaseEntry:        tt.fields.PurchaseEntry,
				Payroll:              tt.fields.Payroll,
				Purchase:             tt.fields.Purchase,
				SalesOrder:           tt.fields.SalesOrder,
				Logistics:            tt.fields.Logistics,
				CRM:                  tt.fields.CRM,
				GeneralJournalEntry:  tt.fields.GeneralJournalEntry,
				OpeningBalance:       tt.fields.OpeningBalance,
				Project:              tt.fields.Project,
				Webhooks:             tt.fields.Webhooks,
				Cashflow:             tt.fields.Cashflow,
				SalesInvoice:         tt.fields.SalesInvoice,
				PurchaseOrder:        tt.fields.PurchaseOrder,
				Sales:                tt.fields.Sales,
				SalesEntry:           tt.fields.SalesEntry,
				Mailbox:              tt.fields.Mailbox,
				Assets:               tt.fields.Assets,
				Financial:            tt.fields.Financial,
				HRM:                  tt.fields.HRM,
				Manufacturing:        tt.fields.Manufacturing,
				Subscription:         tt.fields.Subscription,
				System:               tt.fields.System,
				Activities:           tt.fields.Activities,
			}
			if err := c.SetBaseURL(tt.args.baseURL); (err != nil) != tt.wantErr {
				t.Errorf("Client.SetBaseURL() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got, want := c.client.BaseURL, tt.args.u; reflect.DeepEqual(got, want) != tt.wantEqual {
				t.Errorf("Client.SetBaseURL() got = %v, want %v", got, want)
			}
		})
	}
}

func TestClient_SetUserAgent(t *testing.T) {
	type fields struct {
		client               *api.Client
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
	type args struct {
		userAgent string
	}
	c := api.NewClient(nil)
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"1", fields{client: c}, args{"test user agent"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:               tt.fields.client,
				Budget:               tt.fields.Budget,
				Bulk:                 tt.fields.Bulk,
				ContinuousMonitoring: tt.fields.ContinuousMonitoring,
				Documents:            tt.fields.Documents,
				FinancialTransaction: tt.fields.FinancialTransaction,
				General:              tt.fields.General,
				Inventory:            tt.fields.Inventory,
				Accountancy:          tt.fields.Accountancy,
				Users:                tt.fields.Users,
				VAT:                  tt.fields.VAT,
				Workflow:             tt.fields.Workflow,
				PurchaseEntry:        tt.fields.PurchaseEntry,
				Payroll:              tt.fields.Payroll,
				Purchase:             tt.fields.Purchase,
				SalesOrder:           tt.fields.SalesOrder,
				Logistics:            tt.fields.Logistics,
				CRM:                  tt.fields.CRM,
				GeneralJournalEntry:  tt.fields.GeneralJournalEntry,
				OpeningBalance:       tt.fields.OpeningBalance,
				Project:              tt.fields.Project,
				Webhooks:             tt.fields.Webhooks,
				Cashflow:             tt.fields.Cashflow,
				SalesInvoice:         tt.fields.SalesInvoice,
				PurchaseOrder:        tt.fields.PurchaseOrder,
				Sales:                tt.fields.Sales,
				SalesEntry:           tt.fields.SalesEntry,
				Mailbox:              tt.fields.Mailbox,
				Assets:               tt.fields.Assets,
				Financial:            tt.fields.Financial,
				HRM:                  tt.fields.HRM,
				Manufacturing:        tt.fields.Manufacturing,
				Subscription:         tt.fields.Subscription,
				System:               tt.fields.System,
				Activities:           tt.fields.Activities,
			}
			c.SetUserAgent(tt.args.userAgent)

			if got, want := c.client.UserAgent, tt.args.userAgent; got != want {
				t.Errorf("Client.SetUserAgent(%v) = %v, want %v", tt.args.userAgent, got, want)
			}
		})
	}
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}
func TestClient_GetCurrentDivisionID(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	want := 100

	mux.HandleFunc("/api/v1/current/Me", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{ "d": { "__next": "", "results": [{ "CurrentDivision": 100 }]}}`)
	})

	ctx := context.Background()

	got, err := c.GetCurrentDivisionID(ctx)
	if err != nil {
		t.Errorf("Client.GetCurrentDivisionID() error = %v", err)
	}

	if got != want {
		t.Errorf("Client.GetCurrentDivisionID() = %v, want %v", got, want)
	}
}

func TestClient_GetCurrentDivisionID_error_count(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/v1/current/Me", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{ "d": { "__next": "", "results": [{ "CurrentDivision": 100 }, { "CurrentDivision": 200 }]}}`)
	})

	ctx := context.Background()

	_, err := c.GetCurrentDivisionID(ctx)
	if err == nil {
		t.Error("Client.GetCurrentDivisionID() want error but got nil")
	}
}

func TestClient_GetCurrentDivisionID_error(t *testing.T) {
	c, _, _, teardown := setup()
	defer teardown()

	ctx := context.Background()

	_, err := c.GetCurrentDivisionID(ctx)
	if err == nil {
		t.Error("Client.GetCurrentDivisionID() want error but got nil")
	}
}

func TestBool(t *testing.T) {
	type args struct {
		v bool
	}
	b := true
	tests := []struct {
		name string
		args args
		want *bool
	}{
		{"1", args{b}, &b},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bool(tt.args.v); *got != *tt.want {
				t.Errorf("Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	type args struct {
		v int
	}
	i := 1000
	tests := []struct {
		name string
		args args
		want *int
	}{
		{"1", args{i}, &i},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.v); *got != *tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	type args struct {
		v int64
	}
	i := int64(1000)
	tests := []struct {
		name string
		args args
		want *int64
	}{
		{"1", args{i}, &i},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64(tt.args.v); *got != *tt.want {
				t.Errorf("Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	type args struct {
		v string
	}
	s := "teststring"
	tests := []struct {
		name string
		args args
		want *string
	}{
		{"1", args{s}, &s},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.v); *got != *tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate(t *testing.T) {
	type args struct {
		v time.Time
	}
	d := time.Now()
	tests := []struct {
		name string
		args args
		want *types.Date
	}{
		{"1", args{d}, &types.Date{Time: d}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Date(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Date() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	type args struct {
		v float64
	}
	f := float64(10.12)
	tests := []struct {
		name string
		args args
		want *float64
	}{
		{"1", args{f}, &f},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64(tt.args.v); *got != *tt.want {
				t.Errorf("Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGUID(t *testing.T) {
	type args struct {
		v uuid.UUID
	}
	u := uuid.Must(uuid.NewV4())
	tests := []struct {
		name string
		args args
		want *types.GUID
	}{
		{"1", args{u}, &types.GUID{UUID: u}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GUID(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURL(t *testing.T) {
	type args struct {
		v *url.URL
	}
	u, _ := url.Parse("https://start.exactonline.nl")
	tests := []struct {
		name string
		args args
		want *types.URL
	}{
		{"1", args{u}, &types.URL{URL: u}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := URL(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URL() = %v, want %v", got, tt.want)
			}
		})
	}
}
