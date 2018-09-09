// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package salesinvoice

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

func SalesInvoicesPrimaryPropertySample() *types.GUID {
	v := types.NewGUID()
	return &v
}

func SalesInvoicesEntityWithPopulatedPrimaryProperty() *SalesInvoices {
	return &SalesInvoices{InvoiceID: SalesInvoicesPrimaryPropertySample()}
}

func SalesInvoicesStringOfPrimaryProperty(v *types.GUID) string {
	return v.String()
}

func SalesInvoicesStringJSONOfPrimaryProperty(v *types.GUID) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestSalesInvoicesEndpoint_GetPrimary(t *testing.T) {
	var want types.GUID
	n := &SalesInvoices{InvoiceID: &want}

	if got := n.GetPrimary(); !reflect.DeepEqual(*got, want) {
		t.Errorf("SalesInvoicesEndpoint.GetPrimary() failed, got: %v, want: %v", *got, want)
	}
}

func TestSalesInvoicesEndpoint_UserHasRights(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SalesInvoicesEndpoint.List returned error: %v", e)
	}

	acceptHeaders := []string{"application/json"}

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))

		q := r.URL.Query()

		if got, want := q.Get("endpoint"), "'salesinvoice/SalesInvoices'"; got != want {
			t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
		}

		if got, want := q.Get("method"), "GET"; got != want {
			t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
		}

		fmt.Fprint(w, `{ "d": { "UserHasRights": true } }`)
	})

	got, err := s.SalesInvoices.UserHasRights(context.Background(), 0, "GET")
	if err != nil {
		t.Errorf("s.SalesInvoices.UserHasRights should not return an error = %v", err)
	}

	if got != true {
		t.Errorf("s.SalesInvoices.UserHasRights should return true, got: %v", got)
	}
}

func TestSalesInvoicesEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/salesinvoice/SalesInvoices", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SalesInvoicesEndpoint.List returned error: %v, with url /api/v1/{division}/salesinvoice/SalesInvoices?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/salesinvoice/SalesInvoices", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SalesInvoicesEndpoint.List returned error: %v, with url /api/v1/{division}/salesinvoice/SalesInvoices?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := SalesInvoicesPrimaryPropertySample()
	gs := SalesInvoicesStringJSONOfPrimaryProperty(g)

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "InvoiceID": `+gs+`}]}}`)
		}
	})

	entities, err := s.SalesInvoices.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("SalesInvoicesEndpoint.List returned error: %v", err)
	}

	want := []*SalesInvoices{{InvoiceID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("SalesInvoicesEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestSalesInvoicesEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/salesinvoice/SalesInvoices", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SalesInvoicesEndpoint.List returned error: %v, with url /api/v1/{division}/salesinvoice/SalesInvoices", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/salesinvoice/SalesInvoices", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SalesInvoicesEndpoint.List returned error: %v, with url /api/v1/{division}/salesinvoice/SalesInvoices", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := SalesInvoicesPrimaryPropertySample()
	gs := SalesInvoicesStringJSONOfPrimaryProperty(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "InvoiceID": `+gs+`}]}}`)
	})

	entities, err := s.SalesInvoices.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("SalesInvoicesEndpoint.List returned error: %v", err)
	}

	want := []*SalesInvoices{{InvoiceID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("SalesInvoicesEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestSalesInvoicesEndpoint_Get(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	s1 := SalesInvoicesPrimaryPropertySample()
	type args struct {
		ctx      context.Context
		division int
		id       *types.GUID
	}
	tests := []struct {
		name    string
		args    args
		want    *SalesInvoices
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, s1},
			&SalesInvoices{InvoiceID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/salesinvoice/SalesInvoices", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in SalesInvoicesEndpoint.Delete() returned error: %v, with url /api/v1/{division}/salesinvoice/SalesInvoices", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in SalesInvoicesEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.want)
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.SalesInvoices.Get(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("SalesInvoicesEndpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesInvoicesEndpoint.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSalesInvoicesEndpoint_New(t *testing.T) {
	s, _, _, teardown := setup()
	defer teardown()
	got := s.SalesInvoices.New()
	want := &SalesInvoices{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SalesInvoicesEndpoint.New() expected to return %v, got %v", want, got)
	}
}

func TestSalesInvoicesEndpoint_Create(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *SalesInvoices
	}
	tests := []struct {
		name    string
		args    args
		want    *SalesInvoices
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &SalesInvoices{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&SalesInvoices{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/salesinvoice/SalesInvoices", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in SalesInvoicesEndpoint.Create returned error: %v, with url /api/v1/{division}/salesinvoice/SalesInvoices", e)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "POST")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				testBody(t, r, `{"__metadata":{"uri":"https://start.exactonline.nl"}}`+"\n")
				fmt.Fprint(w, `{ "d": { "__metadata": { "uri": "https://start.exactonline.nl"}}}`)
			})

			got, err := s.SalesInvoices.Create(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("SalesInvoicesEndpoint.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesInvoicesEndpoint.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSalesInvoicesEndpoint_Update(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *SalesInvoices
	}
	s1 := SalesInvoicesPrimaryPropertySample()
	tests := []struct {
		name    string
		args    args
		want    *SalesInvoices
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &SalesInvoices{InvoiceID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&SalesInvoices{InvoiceID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/salesinvoice/SalesInvoices", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in SalesInvoicesEndpoint.Update returned error: %v, with url /api/v1/{division}/salesinvoice/SalesInvoices", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.entity.GetPrimary())
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in SalesInvoicesEndpoint.Update returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "PUT")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.args.entity)
				testBody(t, r, string(b)+"\n")
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.SalesInvoices.Update(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("SalesInvoicesEndpoint.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesInvoicesEndpoint.Update() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestSalesInvoicesEndpoint_Delete(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		id       *types.GUID
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, SalesInvoicesPrimaryPropertySample()},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/salesinvoice/SalesInvoices", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in SalesInvoicesEndpoint.Delete() returned error: %v, with url /api/v1/{division}/salesinvoice/SalesInvoices", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in SalesInvoicesEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "DELETE")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testBody(t, r, "")
				w.WriteHeader(http.StatusNoContent)
			})

			err := s.SalesInvoices.Delete(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("SalesInvoicesEndpoint.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
