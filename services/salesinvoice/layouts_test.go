// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package salesinvoice

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/mcnijman/go-exactonline/types"
)

func TestLayoutsEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/salesinvoice/Layouts?$select=*", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in LayoutsEndpoint.List returned error: %v, with url /api/v1/{division}/salesinvoice/Layouts?$select=*", e)
	}
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/salesinvoice/Layouts?$skiptoken=foo", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in LayoutsEndpoint.List returned error: %v, with url /api/v1/{division}/salesinvoice/Layouts?$skiptoken=foo", e2)
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

	entities, err := s.Layouts.List(context.Background(), 0, true)
	if err != nil {
		t.Errorf("LayoutsEndpoint.List returned error: %v", err)
	}

	want := []*Layouts{{ID: &g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("LayoutsEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestLayoutsEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/salesinvoice/Layouts?$select=*", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in LayoutsEndpoint.List returned error: %v, with url /api/v1/{division}/salesinvoice/Layouts?$select=*", e)
	}
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/salesinvoice/Layouts?$skiptoken=foo", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in LayoutsEndpoint.List returned error: %v, with url /api/v1/{division}/salesinvoice/Layouts?$skiptoken=foo", e2)
	}

	g := types.NewGUID()
	gs := g.String()
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": "`+gs+`"}]}}`)
	})

	entities, err := s.Layouts.List(context.Background(), 0, false)
	if err != nil {
		t.Errorf("LayoutsEndpoint.List returned error: %v", err)
	}

	want := []*Layouts{{ID: &g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("LayoutsEndpoint.List returned %+v, want %+v", entities, want)
	}
}
