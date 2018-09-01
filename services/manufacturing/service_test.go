// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package manufacturing

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/mcnijman/go-exactonline/api"
)

func TestNewManufacturingService(t *testing.T) {
	c := api.NewClient(nil)
	s := NewManufacturingService(c)

	if !reflect.DeepEqual(c, s.client) {
		t.Error("Clients are supposed to be the same")
	}

	if s.BillOfMaterialMaterials == nil {
		t.Error("Property BillOfMaterialMaterials should not be nil")
	}
	if s.BillOfMaterialVersions == nil {
		t.Error("Property BillOfMaterialVersions should not be nil")
	}
	if s.ByProductReceipts == nil {
		t.Error("Property ByProductReceipts should not be nil")
	}
	if s.ByProductReversals == nil {
		t.Error("Property ByProductReversals should not be nil")
	}
	if s.MaterialIssues == nil {
		t.Error("Property MaterialIssues should not be nil")
	}
	if s.MaterialReversals == nil {
		t.Error("Property MaterialReversals should not be nil")
	}
	if s.OperationResources == nil {
		t.Error("Property OperationResources should not be nil")
	}
	if s.Operations == nil {
		t.Error("Property Operations should not be nil")
	}
	if s.ProductionAreas == nil {
		t.Error("Property ProductionAreas should not be nil")
	}
	if s.ShopOrderMaterialPlans == nil {
		t.Error("Property ShopOrderMaterialPlans should not be nil")
	}
	if s.ShopOrderReceipts == nil {
		t.Error("Property ShopOrderReceipts should not be nil")
	}
	if s.ShopOrderReversals == nil {
		t.Error("Property ShopOrderReversals should not be nil")
	}
	if s.ShopOrderRoutingStepPlans == nil {
		t.Error("Property ShopOrderRoutingStepPlans should not be nil")
	}
	if s.ShopOrders == nil {
		t.Error("Property ShopOrders should not be nil")
	}
	if s.StageForDeliveryReceipts == nil {
		t.Error("Property StageForDeliveryReceipts should not be nil")
	}
	if s.StageForDeliveryReversals == nil {
		t.Error("Property StageForDeliveryReversals should not be nil")
	}
	if s.SubOrderReceipts == nil {
		t.Error("Property SubOrderReceipts should not be nil")
	}
	if s.SubOrderReversals == nil {
		t.Error("Property SubOrderReversals should not be nil")
	}
	if s.TimeTransactions == nil {
		t.Error("Property TimeTransactions should not be nil")
	}
	if s.Workcenters == nil {
		t.Error("Property Workcenters should not be nil")
	}
}

// setup sets up a test HTTP server along with a exactonline.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (service *ManufacturingService, mux *http.ServeMux, serverURL string, teardown func()) {
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
	client := api.NewClient(nil)
	url, _ := url.Parse(server.URL + "/")
	client.BaseURL = url
	service = NewManufacturingService(client)

	return service, mux, server.URL, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %q, want %q", header, got, want)
	}
}

func testURLParseError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func testBody(t *testing.T, r *http.Request, want string) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Errorf("Error reading request body: %v", err)
	}
	if got := string(b); got != want {
		t.Errorf("request Body is %s, want %s", got, want)
	}
}

// Helper function to test that a value is marshalled to JSON as expected.
func testJSONMarshal(t *testing.T, v interface{}, want string) {
	j, err := json.Marshal(v)
	if err != nil {
		t.Errorf("Unable to marshal JSON for %v", v)
	}

	w := new(bytes.Buffer)
	err = json.Compact(w, []byte(want))
	if err != nil {
		t.Errorf("String is not valid json: %s", want)
	}

	if w.String() != string(j) {
		t.Errorf("json.Marshal(%q) returned %s, want %s", v, j, w)
	}

	// now go the other direction and make sure things unmarshal as expected
	u := reflect.ValueOf(v).Interface()
	if err := json.Unmarshal([]byte(want), u); err != nil {
		t.Errorf("Unable to unmarshal JSON for %v", want)
	}

	if !reflect.DeepEqual(v, u) {
		t.Errorf("json.Unmarshal(%q) returned %s, want %s", want, u, v)
	}
}
