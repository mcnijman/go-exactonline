// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package cashflow

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

func PaymentConditionsPrimaryPropertySample() *types.GUID {
	v := types.NewGUID()
	return &v
}

func PaymentConditionsEntityWithPopulatedPrimaryProperty() *PaymentConditions {
	return &PaymentConditions{ID: PaymentConditionsPrimaryPropertySample()}
}

func PaymentConditionsStringOfPrimaryProperty(v *types.GUID) string {
	return v.String()
}

func PaymentConditionsStringJSONOfPrimaryProperty(v *types.GUID) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestPaymentConditionsEndpoint_GetPrimary(t *testing.T) {
	var want types.GUID
	n := &PaymentConditions{ID: &want}

	if got := n.GetPrimary(); !reflect.DeepEqual(*got, want) {
		t.Errorf("PaymentConditionsEndpoint.GetPrimary() failed, got: %v, want: %v", *got, want)
	}
}

func TestPaymentConditionsEndpoint_UserHasRights(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in PaymentConditionsEndpoint.List returned error: %v", e)
	}

	acceptHeaders := []string{"application/json"}

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))

		q := r.URL.Query()

		if got, want := q.Get("endpoint"), "'cashflow/PaymentConditions'"; got != want {
			t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
		}

		if got, want := q.Get("method"), "GET"; got != want {
			t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
		}

		fmt.Fprint(w, `{ "d": { "UserHasRights": true } }`)
	})

	got, err := s.PaymentConditions.UserHasRights(context.Background(), 0, "GET")
	if err != nil {
		t.Errorf("s.PaymentConditions.UserHasRights should not return an error = %v", err)
	}

	if got != true {
		t.Errorf("s.PaymentConditions.UserHasRights should return true, got: %v", got)
	}
}

func TestPaymentConditionsEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/cashflow/PaymentConditions", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in PaymentConditionsEndpoint.List returned error: %v, with url /api/v1/{division}/cashflow/PaymentConditions?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/cashflow/PaymentConditions", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in PaymentConditionsEndpoint.List returned error: %v, with url /api/v1/{division}/cashflow/PaymentConditions?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := PaymentConditionsPrimaryPropertySample()
	gs := PaymentConditionsStringJSONOfPrimaryProperty(g)

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
		}
	})

	entities, err := s.PaymentConditions.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("PaymentConditionsEndpoint.List returned error: %v", err)
	}

	want := []*PaymentConditions{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("PaymentConditionsEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestPaymentConditionsEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/cashflow/PaymentConditions", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in PaymentConditionsEndpoint.List returned error: %v, with url /api/v1/{division}/cashflow/PaymentConditions", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/cashflow/PaymentConditions", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in PaymentConditionsEndpoint.List returned error: %v, with url /api/v1/{division}/cashflow/PaymentConditions", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := PaymentConditionsPrimaryPropertySample()
	gs := PaymentConditionsStringJSONOfPrimaryProperty(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
	})

	entities, err := s.PaymentConditions.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("PaymentConditionsEndpoint.List returned error: %v", err)
	}

	want := []*PaymentConditions{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("PaymentConditionsEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestPaymentConditionsEndpoint_Get(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	s1 := PaymentConditionsPrimaryPropertySample()
	type args struct {
		ctx      context.Context
		division int
		id       *types.GUID
	}
	tests := []struct {
		name    string
		args    args
		want    *PaymentConditions
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, s1},
			&PaymentConditions{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/cashflow/PaymentConditions", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in PaymentConditionsEndpoint.Delete() returned error: %v, with url /api/v1/{division}/cashflow/PaymentConditions", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in PaymentConditionsEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.want)
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.PaymentConditions.Get(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("PaymentConditionsEndpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentConditionsEndpoint.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentConditionsEndpoint_New(t *testing.T) {
	s, _, _, teardown := setup()
	defer teardown()
	got := s.PaymentConditions.New()
	want := &PaymentConditions{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("PaymentConditionsEndpoint.New() expected to return %v, got %v", want, got)
	}
}

func TestPaymentConditionsEndpoint_Create(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *PaymentConditions
	}
	tests := []struct {
		name    string
		args    args
		want    *PaymentConditions
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &PaymentConditions{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&PaymentConditions{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/cashflow/PaymentConditions", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in PaymentConditionsEndpoint.Create returned error: %v, with url /api/v1/{division}/cashflow/PaymentConditions", e)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "POST")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				testBody(t, r, `{"__metadata":{"uri":"https://start.exactonline.nl"}}`+"\n")
				fmt.Fprint(w, `{ "d": { "__metadata": { "uri": "https://start.exactonline.nl"}}}`)
			})

			got, err := s.PaymentConditions.Create(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("PaymentConditionsEndpoint.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentConditionsEndpoint.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
