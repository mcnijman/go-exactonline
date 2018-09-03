// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package project

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

func TestTimeAndBillingEntryRecentActivitiesAndExpensesEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/TimeAndBillingEntryRecentActivitiesAndExpenses", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in TimeAndBillingEntryRecentActivitiesAndExpensesEndpoint.List returned error: %v, with url /api/v1/{division}/read/project/TimeAndBillingEntryRecentActivitiesAndExpenses?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/TimeAndBillingEntryRecentActivitiesAndExpenses", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in TimeAndBillingEntryRecentActivitiesAndExpensesEndpoint.List returned error: %v, with url /api/v1/{division}/read/project/TimeAndBillingEntryRecentActivitiesAndExpenses?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u, opts2)

	g := types.NewGUID()
	gs := g.String()
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": "`+gs+`"}]}}`)
		}
	})

	entities, err := s.TimeAndBillingEntryRecentActivitiesAndExpenses.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("TimeAndBillingEntryRecentActivitiesAndExpensesEndpoint.List returned error: %v", err)
	}

	want := []*TimeAndBillingEntryRecentActivitiesAndExpenses{{ID: &g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("TimeAndBillingEntryRecentActivitiesAndExpensesEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestTimeAndBillingEntryRecentActivitiesAndExpensesEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/TimeAndBillingEntryRecentActivitiesAndExpenses", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in TimeAndBillingEntryRecentActivitiesAndExpensesEndpoint.List returned error: %v, with url /api/v1/{division}/read/project/TimeAndBillingEntryRecentActivitiesAndExpenses?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/TimeAndBillingEntryRecentActivitiesAndExpenses", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in TimeAndBillingEntryRecentActivitiesAndExpensesEndpoint.List returned error: %v, with url /api/v1/{division}/read/project/TimeAndBillingEntryRecentActivitiesAndExpenses?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := types.NewGUID()
	gs := g.String()
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": "`+gs+`"}]}}`)
	})

	entities, err := s.TimeAndBillingEntryRecentActivitiesAndExpenses.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("TimeAndBillingEntryRecentActivitiesAndExpensesEndpoint.List returned error: %v", err)
	}

	want := []*TimeAndBillingEntryRecentActivitiesAndExpenses{{ID: &g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("TimeAndBillingEntryRecentActivitiesAndExpensesEndpoint.List returned %+v, want %+v", entities, want)
	}
}
