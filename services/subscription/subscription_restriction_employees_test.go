// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package subscription

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

func SubscriptionRestrictionEmployeesPrimaryPropertySample() *types.GUID {
	v := types.NewGUID()
	return &v
}

func SubscriptionRestrictionEmployeesEntityWithPopulatedPrimaryProperty() *SubscriptionRestrictionEmployees {
	return &SubscriptionRestrictionEmployees{ID: SubscriptionRestrictionEmployeesPrimaryPropertySample()}
}

func SubscriptionRestrictionEmployeesStringOfPrimaryProperty(v *types.GUID) string {
	return v.String()
}

func SubscriptionRestrictionEmployeesStringJSONOfPrimaryProperty(v *types.GUID) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestSubscriptionRestrictionEmployeesEndpoint_GetPrimary(t *testing.T) {
	var want types.GUID
	n := &SubscriptionRestrictionEmployees{ID: &want}

	if got := n.GetPrimary(); !reflect.DeepEqual(*got, want) {
		t.Errorf("SubscriptionRestrictionEmployeesEndpoint.GetPrimary() failed, got: %v, want: %v", *got, want)
	}
}

func TestSubscriptionRestrictionEmployeesEndpoint_UserHasRights(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SubscriptionRestrictionEmployeesEndpoint.List returned error: %v", e)
	}

	acceptHeaders := []string{"application/json"}

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))

		q := r.URL.Query()

		if got, want := q.Get("endpoint"), "'subscription/SubscriptionRestrictionEmployees'"; got != want {
			t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
		}

		if got, want := q.Get("method"), "GET"; got != want {
			t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
		}

		fmt.Fprint(w, `{ "d": { "UserHasRights": true } }`)
	})

	got, err := s.SubscriptionRestrictionEmployees.UserHasRights(context.Background(), 0, "GET")
	if err != nil {
		t.Errorf("s.SubscriptionRestrictionEmployees.UserHasRights should not return an error = %v", err)
	}

	if got != true {
		t.Errorf("s.SubscriptionRestrictionEmployees.UserHasRights should return true, got: %v", got)
	}
}

func TestSubscriptionRestrictionEmployeesEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/SubscriptionRestrictionEmployees", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SubscriptionRestrictionEmployeesEndpoint.List returned error: %v, with url /api/v1/{division}/subscription/SubscriptionRestrictionEmployees?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/SubscriptionRestrictionEmployees", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SubscriptionRestrictionEmployeesEndpoint.List returned error: %v, with url /api/v1/{division}/subscription/SubscriptionRestrictionEmployees?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := SubscriptionRestrictionEmployeesPrimaryPropertySample()
	gs := SubscriptionRestrictionEmployeesStringJSONOfPrimaryProperty(g)

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
		}
	})

	entities, err := s.SubscriptionRestrictionEmployees.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("SubscriptionRestrictionEmployeesEndpoint.List returned error: %v", err)
	}

	want := []*SubscriptionRestrictionEmployees{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("SubscriptionRestrictionEmployeesEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestSubscriptionRestrictionEmployeesEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/SubscriptionRestrictionEmployees", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SubscriptionRestrictionEmployeesEndpoint.List returned error: %v, with url /api/v1/{division}/subscription/SubscriptionRestrictionEmployees", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/SubscriptionRestrictionEmployees", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SubscriptionRestrictionEmployeesEndpoint.List returned error: %v, with url /api/v1/{division}/subscription/SubscriptionRestrictionEmployees", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := SubscriptionRestrictionEmployeesPrimaryPropertySample()
	gs := SubscriptionRestrictionEmployeesStringJSONOfPrimaryProperty(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
	})

	entities, err := s.SubscriptionRestrictionEmployees.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("SubscriptionRestrictionEmployeesEndpoint.List returned error: %v", err)
	}

	want := []*SubscriptionRestrictionEmployees{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("SubscriptionRestrictionEmployeesEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestSubscriptionRestrictionEmployeesEndpoint_Get(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	s1 := SubscriptionRestrictionEmployeesPrimaryPropertySample()
	type args struct {
		ctx      context.Context
		division int
		id       *types.GUID
	}
	tests := []struct {
		name    string
		args    args
		want    *SubscriptionRestrictionEmployees
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, s1},
			&SubscriptionRestrictionEmployees{ID: s1, MetaData: &api.MetaData{URI: &types.URL{&url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/SubscriptionRestrictionEmployees", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in SubscriptionRestrictionEmployeesEndpoint.Delete() returned error: %v, with url /api/v1/{division}/subscription/SubscriptionRestrictionEmployees", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in SubscriptionRestrictionEmployeesEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.want)
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.SubscriptionRestrictionEmployees.Get(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubscriptionRestrictionEmployeesEndpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubscriptionRestrictionEmployeesEndpoint.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscriptionRestrictionEmployeesEndpoint_New(t *testing.T) {
	s, _, _, teardown := setup()
	defer teardown()
	got := s.SubscriptionRestrictionEmployees.New()
	want := &SubscriptionRestrictionEmployees{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SubscriptionRestrictionEmployeesEndpoint.New() expected to return %v, got %v", want, got)
	}
}

func TestSubscriptionRestrictionEmployeesEndpoint_Create(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *SubscriptionRestrictionEmployees
	}
	tests := []struct {
		name    string
		args    args
		want    *SubscriptionRestrictionEmployees
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &SubscriptionRestrictionEmployees{MetaData: &api.MetaData{URI: &types.URL{&url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&SubscriptionRestrictionEmployees{MetaData: &api.MetaData{URI: &types.URL{&url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/SubscriptionRestrictionEmployees", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in SubscriptionRestrictionEmployeesEndpoint.Create returned error: %v, with url /api/v1/{division}/subscription/SubscriptionRestrictionEmployees", e)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "POST")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				testBody(t, r, `{"__metadata":{"uri":"https://start.exactonline.nl"}}`+"\n")
				fmt.Fprint(w, `{ "d": { "__metadata": { "uri": "https://start.exactonline.nl"}}}`)
			})

			got, err := s.SubscriptionRestrictionEmployees.Create(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubscriptionRestrictionEmployeesEndpoint.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubscriptionRestrictionEmployeesEndpoint.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscriptionRestrictionEmployeesEndpoint_Delete(t *testing.T) {
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
			args{context.Background(), 0, SubscriptionRestrictionEmployeesPrimaryPropertySample()},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/SubscriptionRestrictionEmployees", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in SubscriptionRestrictionEmployeesEndpoint.Delete() returned error: %v, with url /api/v1/{division}/subscription/SubscriptionRestrictionEmployees", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in SubscriptionRestrictionEmployeesEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "DELETE")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testBody(t, r, "")
				w.WriteHeader(http.StatusNoContent)
			})

			err := s.SubscriptionRestrictionEmployees.Delete(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubscriptionRestrictionEmployeesEndpoint.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
