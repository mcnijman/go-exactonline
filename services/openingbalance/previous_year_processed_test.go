// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package openingbalance

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

func PreviousYearProcessedPrimaryPropertySample() *int {
	v := 100
	return &v
}

func PreviousYearProcessedEntityWithPopulatedPrimaryProperty() *PreviousYearProcessed {
	return &PreviousYearProcessed{Division: PreviousYearProcessedPrimaryPropertySample()}
}

func PreviousYearProcessedStringOfPrimaryProperty(v *int) string {
	return strconv.Itoa(*v)
}

func PreviousYearProcessedStringJSONOfPrimaryProperty(v *int) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestPreviousYearProcessedEndpoint_GetPrimary(t *testing.T) {
	var want int
	n := &PreviousYearProcessed{Division: &want}

	if got := n.GetPrimary(); !reflect.DeepEqual(*got, want) {
		t.Errorf("PreviousYearProcessedEndpoint.GetPrimary() failed, got: %v, want: %v", *got, want)
	}
}

func TestPreviousYearProcessedEndpoint_UserHasRights(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in PreviousYearProcessedEndpoint.List returned error: %v", e)
	}

	acceptHeaders := []string{"application/json"}

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))

		q := r.URL.Query()

		if got, want := q.Get("endpoint"), "'PreviousYear/Processed'"; got != want {
			t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
		}

		if got, want := q.Get("method"), "GET"; got != want {
			t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
		}

		fmt.Fprint(w, `{ "d": { "UserHasRights": true } }`)
	})

	got, err := s.PreviousYearProcessed.UserHasRights(context.Background(), 0, "GET")
	if err != nil {
		t.Errorf("s.PreviousYearProcessed.UserHasRights should not return an error = %v", err)
	}

	if got != true {
		t.Errorf("s.PreviousYearProcessed.UserHasRights should return true, got: %v", got)
	}
}

func TestPreviousYearProcessedEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/openingbalance/PreviousYear/Processed", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in PreviousYearProcessedEndpoint.List returned error: %v, with url /api/v1/{division}/openingbalance/PreviousYear/Processed?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/openingbalance/PreviousYear/Processed", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in PreviousYearProcessedEndpoint.List returned error: %v, with url /api/v1/{division}/openingbalance/PreviousYear/Processed?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := PreviousYearProcessedPrimaryPropertySample()
	gs := PreviousYearProcessedStringJSONOfPrimaryProperty(g)

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "Division": `+gs+`}]}}`)
		}
	})

	entities, err := s.PreviousYearProcessed.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("PreviousYearProcessedEndpoint.List returned error: %v", err)
	}

	want := []*PreviousYearProcessed{{Division: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("PreviousYearProcessedEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestPreviousYearProcessedEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/openingbalance/PreviousYear/Processed", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in PreviousYearProcessedEndpoint.List returned error: %v, with url /api/v1/{division}/openingbalance/PreviousYear/Processed", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/openingbalance/PreviousYear/Processed", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in PreviousYearProcessedEndpoint.List returned error: %v, with url /api/v1/{division}/openingbalance/PreviousYear/Processed", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := PreviousYearProcessedPrimaryPropertySample()
	gs := PreviousYearProcessedStringJSONOfPrimaryProperty(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "Division": `+gs+`}]}}`)
	})

	entities, err := s.PreviousYearProcessed.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("PreviousYearProcessedEndpoint.List returned error: %v", err)
	}

	want := []*PreviousYearProcessed{{Division: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("PreviousYearProcessedEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestPreviousYearProcessedEndpoint_Get(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	s1 := PreviousYearProcessedPrimaryPropertySample()
	type args struct {
		ctx      context.Context
		division int
		id       *int
	}
	tests := []struct {
		name    string
		args    args
		want    *PreviousYearProcessed
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, s1},
			&PreviousYearProcessed{Division: s1, MetaData: &api.MetaData{URI: &types.URL{&url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/openingbalance/PreviousYear/Processed", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in PreviousYearProcessedEndpoint.Delete() returned error: %v, with url /api/v1/{division}/openingbalance/PreviousYear/Processed", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in PreviousYearProcessedEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.want)
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.PreviousYearProcessed.Get(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("PreviousYearProcessedEndpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PreviousYearProcessedEndpoint.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
