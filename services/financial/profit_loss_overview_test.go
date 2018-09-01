// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package financial

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestProfitLossOverviewEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/ProfitLossOverview?$select=*", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ProfitLossOverviewEndpoint.List returned error: %v, with url /api/v1/{division}/read/financial/ProfitLossOverview?$select=*", e)
	}
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/ProfitLossOverview?$skiptoken=foo", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ProfitLossOverviewEndpoint.List returned error: %v, with url /api/v1/{division}/read/financial/ProfitLossOverview?$skiptoken=foo", e2)
	}

	g := 100
	gs := strconv.Itoa(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "CurrentYear": `+gs+`}]}}`)
		}
	})

	entities, err := s.ProfitLossOverview.List(context.Background(), 0, true)
	if err != nil {
		t.Errorf("ProfitLossOverviewEndpoint.List returned error: %v", err)
	}

	want := []*ProfitLossOverview{{CurrentYear: &g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("ProfitLossOverviewEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestProfitLossOverviewEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/ProfitLossOverview?$select=*", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ProfitLossOverviewEndpoint.List returned error: %v, with url /api/v1/{division}/read/financial/ProfitLossOverview?$select=*", e)
	}
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/ProfitLossOverview?$skiptoken=foo", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ProfitLossOverviewEndpoint.List returned error: %v, with url /api/v1/{division}/read/financial/ProfitLossOverview?$skiptoken=foo", e2)
	}

	g := 100
	gs := strconv.Itoa(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "CurrentYear": `+gs+`}]}}`)
	})

	entities, err := s.ProfitLossOverview.List(context.Background(), 0, false)
	if err != nil {
		t.Errorf("ProfitLossOverviewEndpoint.List returned error: %v", err)
	}

	want := []*ProfitLossOverview{{CurrentYear: &g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("ProfitLossOverviewEndpoint.List returned %+v, want %+v", entities, want)
	}
}
