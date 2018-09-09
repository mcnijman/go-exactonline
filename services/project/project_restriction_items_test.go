// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package project

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

func ProjectRestrictionItemsPrimaryPropertySample() *types.GUID {
	v := types.NewGUID()
	return &v
}

func ProjectRestrictionItemsEntityWithPopulatedPrimaryProperty() *ProjectRestrictionItems {
	return &ProjectRestrictionItems{ID: ProjectRestrictionItemsPrimaryPropertySample()}
}

func ProjectRestrictionItemsStringOfPrimaryProperty(v *types.GUID) string {
	return v.String()
}

func ProjectRestrictionItemsStringJSONOfPrimaryProperty(v *types.GUID) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestProjectRestrictionItemsEndpoint_GetPrimary(t *testing.T) {
	var want types.GUID
	n := &ProjectRestrictionItems{ID: &want}

	if got := n.GetPrimary(); !reflect.DeepEqual(*got, want) {
		t.Errorf("ProjectRestrictionItemsEndpoint.GetPrimary() failed, got: %v, want: %v", *got, want)
	}
}

func TestProjectRestrictionItemsEndpoint_UserHasRights(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ProjectRestrictionItemsEndpoint.List returned error: %v", e)
	}

	acceptHeaders := []string{"application/json"}

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))

		q := r.URL.Query()

		if got, want := q.Get("endpoint"), "'project/ProjectRestrictionItems'"; got != want {
			t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
		}

		if got, want := q.Get("method"), "GET"; got != want {
			t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
		}

		fmt.Fprint(w, `{ "d": { "UserHasRights": true } }`)
	})

	got, err := s.ProjectRestrictionItems.UserHasRights(context.Background(), 0, "GET")
	if err != nil {
		t.Errorf("s.ProjectRestrictionItems.UserHasRights should not return an error = %v", err)
	}

	if got != true {
		t.Errorf("s.ProjectRestrictionItems.UserHasRights should return true, got: %v", got)
	}
}

func TestProjectRestrictionItemsEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectRestrictionItems", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ProjectRestrictionItemsEndpoint.List returned error: %v, with url /api/v1/{division}/project/ProjectRestrictionItems?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectRestrictionItems", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ProjectRestrictionItemsEndpoint.List returned error: %v, with url /api/v1/{division}/project/ProjectRestrictionItems?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := ProjectRestrictionItemsPrimaryPropertySample()
	gs := ProjectRestrictionItemsStringJSONOfPrimaryProperty(g)

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
		}
	})

	entities, err := s.ProjectRestrictionItems.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("ProjectRestrictionItemsEndpoint.List returned error: %v", err)
	}

	want := []*ProjectRestrictionItems{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("ProjectRestrictionItemsEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestProjectRestrictionItemsEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectRestrictionItems", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ProjectRestrictionItemsEndpoint.List returned error: %v, with url /api/v1/{division}/project/ProjectRestrictionItems", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectRestrictionItems", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ProjectRestrictionItemsEndpoint.List returned error: %v, with url /api/v1/{division}/project/ProjectRestrictionItems", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := ProjectRestrictionItemsPrimaryPropertySample()
	gs := ProjectRestrictionItemsStringJSONOfPrimaryProperty(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
	})

	entities, err := s.ProjectRestrictionItems.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("ProjectRestrictionItemsEndpoint.List returned error: %v", err)
	}

	want := []*ProjectRestrictionItems{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("ProjectRestrictionItemsEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestProjectRestrictionItemsEndpoint_Get(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	s1 := ProjectRestrictionItemsPrimaryPropertySample()
	type args struct {
		ctx      context.Context
		division int
		id       *types.GUID
	}
	tests := []struct {
		name    string
		args    args
		want    *ProjectRestrictionItems
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, s1},
			&ProjectRestrictionItems{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectRestrictionItems", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in ProjectRestrictionItemsEndpoint.Delete() returned error: %v, with url /api/v1/{division}/project/ProjectRestrictionItems", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in ProjectRestrictionItemsEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.want)
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.ProjectRestrictionItems.Get(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectRestrictionItemsEndpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProjectRestrictionItemsEndpoint.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectRestrictionItemsEndpoint_New(t *testing.T) {
	s, _, _, teardown := setup()
	defer teardown()
	got := s.ProjectRestrictionItems.New()
	want := &ProjectRestrictionItems{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ProjectRestrictionItemsEndpoint.New() expected to return %v, got %v", want, got)
	}
}

func TestProjectRestrictionItemsEndpoint_Create(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *ProjectRestrictionItems
	}
	tests := []struct {
		name    string
		args    args
		want    *ProjectRestrictionItems
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &ProjectRestrictionItems{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&ProjectRestrictionItems{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectRestrictionItems", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in ProjectRestrictionItemsEndpoint.Create returned error: %v, with url /api/v1/{division}/project/ProjectRestrictionItems", e)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "POST")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				testBody(t, r, `{"__metadata":{"uri":"https://start.exactonline.nl"}}`+"\n")
				fmt.Fprint(w, `{ "d": { "__metadata": { "uri": "https://start.exactonline.nl"}}}`)
			})

			got, err := s.ProjectRestrictionItems.Create(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectRestrictionItemsEndpoint.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProjectRestrictionItemsEndpoint.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectRestrictionItemsEndpoint_Update(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *ProjectRestrictionItems
	}
	s1 := ProjectRestrictionItemsPrimaryPropertySample()
	tests := []struct {
		name    string
		args    args
		want    *ProjectRestrictionItems
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &ProjectRestrictionItems{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&ProjectRestrictionItems{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectRestrictionItems", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in ProjectRestrictionItemsEndpoint.Update returned error: %v, with url /api/v1/{division}/project/ProjectRestrictionItems", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.entity.GetPrimary())
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in ProjectRestrictionItemsEndpoint.Update returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "PUT")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.args.entity)
				testBody(t, r, string(b)+"\n")
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.ProjectRestrictionItems.Update(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectRestrictionItemsEndpoint.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProjectRestrictionItemsEndpoint.Update() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestProjectRestrictionItemsEndpoint_Delete(t *testing.T) {
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
			args{context.Background(), 0, ProjectRestrictionItemsPrimaryPropertySample()},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectRestrictionItems", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in ProjectRestrictionItemsEndpoint.Delete() returned error: %v, with url /api/v1/{division}/project/ProjectRestrictionItems", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in ProjectRestrictionItemsEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "DELETE")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testBody(t, r, "")
				w.WriteHeader(http.StatusNoContent)
			})

			err := s.ProjectRestrictionItems.Delete(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectRestrictionItemsEndpoint.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
