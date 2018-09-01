// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package salesentry

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/mcnijman/go-exactonline/types"
)

func TestSalesEntryLinesEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/salesentry/SalesEntryLines?$select=*", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SalesEntryLinesEndpoint.List returned error: %v, with url /api/v1/{division}/salesentry/SalesEntryLines?$select=*", e)
	}
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/salesentry/SalesEntryLines?$skiptoken=foo", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SalesEntryLinesEndpoint.List returned error: %v, with url /api/v1/{division}/salesentry/SalesEntryLines?$skiptoken=foo", e2)
	}

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

	entities, err := s.SalesEntryLines.List(context.Background(), 0, true)
	if err != nil {
		t.Errorf("SalesEntryLinesEndpoint.List returned error: %v", err)
	}

	want := []*SalesEntryLines{{ID: &g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("SalesEntryLinesEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestSalesEntryLinesEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/salesentry/SalesEntryLines?$select=*", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SalesEntryLinesEndpoint.List returned error: %v, with url /api/v1/{division}/salesentry/SalesEntryLines?$select=*", e)
	}
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/salesentry/SalesEntryLines?$skiptoken=foo", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SalesEntryLinesEndpoint.List returned error: %v, with url /api/v1/{division}/salesentry/SalesEntryLines?$skiptoken=foo", e2)
	}

	g := types.NewGUID()
	gs := g.String()
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": "`+gs+`"}]}}`)
	})

	entities, err := s.SalesEntryLines.List(context.Background(), 0, false)
	if err != nil {
		t.Errorf("SalesEntryLinesEndpoint.List returned error: %v", err)
	}

	want := []*SalesEntryLines{{ID: &g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("SalesEntryLinesEndpoint.List returned %+v, want %+v", entities, want)
	}
}
