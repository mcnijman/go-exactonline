// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package activities

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

func ComplaintsPrimaryPropertySample() *types.GUID {
	v := types.NewGUID()
	return &v
}

func ComplaintsEntityWithPopulatedPrimaryProperty() *Complaints {
	return &Complaints{ID: ComplaintsPrimaryPropertySample()}
}

func ComplaintsStringOfPrimaryProperty(v *types.GUID) string {
	return v.String()
}

func ComplaintsStringJSONOfPrimaryProperty(v *types.GUID) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestComplaintsEndpoint_GetPrimary(t *testing.T) {
	var want types.GUID
	n := &Complaints{ID: &want}

	if got := n.GetPrimary(); !reflect.DeepEqual(*got, want) {
		t.Errorf("ComplaintsEndpoint.GetPrimary() failed, got: %v, want: %v", *got, want)
	}
}

func TestComplaintsEndpoint_UserHasRights(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ComplaintsEndpoint.List returned error: %v", e)
	}

	acceptHeaders := []string{"application/json"}

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))

		q := r.URL.Query()

		if got, want := q.Get("endpoint"), "'activities/Complaints'"; got != want {
			t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
		}

		if got, want := q.Get("method"), "GET"; got != want {
			t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
		}

		fmt.Fprint(w, `{ "d": { "UserHasRights": true } }`)
	})

	got, err := s.Complaints.UserHasRights(context.Background(), 0, "GET")
	if err != nil {
		t.Errorf("s.Complaints.UserHasRights should not return an error = %v", err)
	}

	if got != true {
		t.Errorf("s.Complaints.UserHasRights should return true, got: %v", got)
	}
}

func TestComplaintsEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/activities/Complaints", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ComplaintsEndpoint.List returned error: %v, with url /api/v1/{division}/activities/Complaints?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/activities/Complaints", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ComplaintsEndpoint.List returned error: %v, with url /api/v1/{division}/activities/Complaints?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := ComplaintsPrimaryPropertySample()
	gs := ComplaintsStringJSONOfPrimaryProperty(g)

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
		}
	})

	entities, err := s.Complaints.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("ComplaintsEndpoint.List returned error: %v", err)
	}

	want := []*Complaints{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("ComplaintsEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestComplaintsEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/activities/Complaints", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ComplaintsEndpoint.List returned error: %v, with url /api/v1/{division}/activities/Complaints", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/activities/Complaints", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ComplaintsEndpoint.List returned error: %v, with url /api/v1/{division}/activities/Complaints", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := ComplaintsPrimaryPropertySample()
	gs := ComplaintsStringJSONOfPrimaryProperty(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
	})

	entities, err := s.Complaints.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("ComplaintsEndpoint.List returned error: %v", err)
	}

	want := []*Complaints{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("ComplaintsEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestComplaintsEndpoint_Get(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	s1 := ComplaintsPrimaryPropertySample()
	type args struct {
		ctx      context.Context
		division int
		id       *types.GUID
	}
	tests := []struct {
		name    string
		args    args
		want    *Complaints
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, s1},
			&Complaints{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/activities/Complaints", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in ComplaintsEndpoint.Delete() returned error: %v, with url /api/v1/{division}/activities/Complaints", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in ComplaintsEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.want)
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.Complaints.Get(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComplaintsEndpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComplaintsEndpoint.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplaintsEndpoint_New(t *testing.T) {
	s, _, _, teardown := setup()
	defer teardown()
	got := s.Complaints.New()
	want := &Complaints{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ComplaintsEndpoint.New() expected to return %v, got %v", want, got)
	}
}

func TestComplaintsEndpoint_Create(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *Complaints
	}
	tests := []struct {
		name    string
		args    args
		want    *Complaints
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &Complaints{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&Complaints{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/activities/Complaints", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in ComplaintsEndpoint.Create returned error: %v, with url /api/v1/{division}/activities/Complaints", e)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "POST")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				testBody(t, r, `{"__metadata":{"uri":"https://start.exactonline.nl"}}`+"\n")
				fmt.Fprint(w, `{ "d": { "__metadata": { "uri": "https://start.exactonline.nl"}}}`)
			})

			got, err := s.Complaints.Create(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComplaintsEndpoint.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComplaintsEndpoint.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
