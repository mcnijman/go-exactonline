// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package crm

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/mcnijman/go-exactonline/types"
)

func TestReasonCodesEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/ReasonCodes?$select=*", 0)
	u2, e := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/ReasonCodes?$skiptoken=foo", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ReasonCodesEndpoint.List returned error: %v, with url /api/v1/{division}/crm/ReasonCodes", e)
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

	entities, err := s.ReasonCodes.List(context.Background(), 0, true)
	if err != nil {
		t.Errorf("ReasonCodesEndpoint.List returned error: %v", err)
	}

	want := []*ReasonCodes{{ID: &g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("ReasonCodesEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestReasonCodesEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/ReasonCodes?$select=*", 0)
	u2, e := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/ReasonCodes?$skiptoken=foo", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ReasonCodesEndpoint.List returned error: %v, with url /api/v1/{division}/crm/ReasonCodes", e)
	}

	g := types.NewGUID()
	gs := g.String()
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": "`+gs+`"}]}}`)
	})

	entities, err := s.ReasonCodes.List(context.Background(), 0, false)
	if err != nil {
		t.Errorf("ReasonCodesEndpoint.List returned error: %v", err)
	}

	want := []*ReasonCodes{{ID: &g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("ReasonCodesEndpoint.List returned %+v, want %+v", entities, want)
	}
}
