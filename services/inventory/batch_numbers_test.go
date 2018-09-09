// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package inventory

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

func BatchNumbersPrimaryPropertySample() *types.GUID {
	v := types.NewGUID()
	return &v
}

func BatchNumbersEntityWithPopulatedPrimaryProperty() *BatchNumbers {
	return &BatchNumbers{ID: BatchNumbersPrimaryPropertySample()}
}

func BatchNumbersStringOfPrimaryProperty(v *types.GUID) string {
	return v.String()
}

func BatchNumbersStringJSONOfPrimaryProperty(v *types.GUID) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestBatchNumbersEndpoint_GetPrimary(t *testing.T) {
	var want types.GUID
	n := &BatchNumbers{ID: &want}

	if got := n.GetPrimary(); !reflect.DeepEqual(*got, want) {
		t.Errorf("BatchNumbersEndpoint.GetPrimary() failed, got: %v, want: %v", *got, want)
	}
}

func TestBatchNumbersEndpoint_UserHasRights(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in BatchNumbersEndpoint.List returned error: %v", e)
	}

	acceptHeaders := []string{"application/json"}

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))

		q := r.URL.Query()

		if got, want := q.Get("endpoint"), "'inventory/BatchNumbers'"; got != want {
			t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
		}

		if got, want := q.Get("method"), "GET"; got != want {
			t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
		}

		fmt.Fprint(w, `{ "d": { "UserHasRights": true } }`)
	})

	got, err := s.BatchNumbers.UserHasRights(context.Background(), 0, "GET")
	if err != nil {
		t.Errorf("s.BatchNumbers.UserHasRights should not return an error = %v", err)
	}

	if got != true {
		t.Errorf("s.BatchNumbers.UserHasRights should return true, got: %v", got)
	}
}

func TestBatchNumbersEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/inventory/BatchNumbers", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in BatchNumbersEndpoint.List returned error: %v, with url /api/v1/{division}/inventory/BatchNumbers?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/inventory/BatchNumbers", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in BatchNumbersEndpoint.List returned error: %v, with url /api/v1/{division}/inventory/BatchNumbers?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := BatchNumbersPrimaryPropertySample()
	gs := BatchNumbersStringJSONOfPrimaryProperty(g)

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
		}
	})

	entities, err := s.BatchNumbers.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("BatchNumbersEndpoint.List returned error: %v", err)
	}

	want := []*BatchNumbers{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("BatchNumbersEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestBatchNumbersEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/inventory/BatchNumbers", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in BatchNumbersEndpoint.List returned error: %v, with url /api/v1/{division}/inventory/BatchNumbers", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/inventory/BatchNumbers", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in BatchNumbersEndpoint.List returned error: %v, with url /api/v1/{division}/inventory/BatchNumbers", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := BatchNumbersPrimaryPropertySample()
	gs := BatchNumbersStringJSONOfPrimaryProperty(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
	})

	entities, err := s.BatchNumbers.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("BatchNumbersEndpoint.List returned error: %v", err)
	}

	want := []*BatchNumbers{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("BatchNumbersEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestBatchNumbersEndpoint_Get(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	s1 := BatchNumbersPrimaryPropertySample()
	type args struct {
		ctx      context.Context
		division int
		id       *types.GUID
	}
	tests := []struct {
		name    string
		args    args
		want    *BatchNumbers
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, s1},
			&BatchNumbers{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/inventory/BatchNumbers", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in BatchNumbersEndpoint.Delete() returned error: %v, with url /api/v1/{division}/inventory/BatchNumbers", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in BatchNumbersEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.want)
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.BatchNumbers.Get(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("BatchNumbersEndpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BatchNumbersEndpoint.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
