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

func TasksPrimaryPropertySample() *types.GUID {
	v := types.NewGUID()
	return &v
}

func TasksEntityWithPopulatedPrimaryProperty() *Tasks {
	return &Tasks{ID: TasksPrimaryPropertySample()}
}

func TasksStringOfPrimaryProperty(v *types.GUID) string {
	return v.String()
}

func TasksStringJSONOfPrimaryProperty(v *types.GUID) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestTasksEndpoint_GetPrimary(t *testing.T) {
	var want types.GUID
	n := &Tasks{ID: &want}

	if got := n.GetPrimary(); !reflect.DeepEqual(*got, want) {
		t.Errorf("TasksEndpoint.GetPrimary() failed, got: %v, want: %v", *got, want)
	}
}

func TestTasksEndpoint_UserHasRights(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in TasksEndpoint.List returned error: %v", e)
	}

	acceptHeaders := []string{"application/json"}

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))

		q := r.URL.Query()

		if got, want := q.Get("endpoint"), "'activities/Tasks'"; got != want {
			t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
		}

		if got, want := q.Get("method"), "GET"; got != want {
			t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
		}

		fmt.Fprint(w, `{ "d": { "UserHasRights": true } }`)
	})

	got, err := s.Tasks.UserHasRights(context.Background(), 0, "GET")
	if err != nil {
		t.Errorf("s.Tasks.UserHasRights should not return an error = %v", err)
	}

	if got != true {
		t.Errorf("s.Tasks.UserHasRights should return true, got: %v", got)
	}
}

func TestTasksEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/activities/Tasks", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in TasksEndpoint.List returned error: %v, with url /api/v1/{division}/activities/Tasks?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/activities/Tasks", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in TasksEndpoint.List returned error: %v, with url /api/v1/{division}/activities/Tasks?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := TasksPrimaryPropertySample()
	gs := TasksStringJSONOfPrimaryProperty(g)

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
		}
	})

	entities, err := s.Tasks.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("TasksEndpoint.List returned error: %v", err)
	}

	want := []*Tasks{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("TasksEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestTasksEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/activities/Tasks", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in TasksEndpoint.List returned error: %v, with url /api/v1/{division}/activities/Tasks", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/activities/Tasks", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in TasksEndpoint.List returned error: %v, with url /api/v1/{division}/activities/Tasks", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := TasksPrimaryPropertySample()
	gs := TasksStringJSONOfPrimaryProperty(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
	})

	entities, err := s.Tasks.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("TasksEndpoint.List returned error: %v", err)
	}

	want := []*Tasks{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("TasksEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestTasksEndpoint_Get(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	s1 := TasksPrimaryPropertySample()
	type args struct {
		ctx      context.Context
		division int
		id       *types.GUID
	}
	tests := []struct {
		name    string
		args    args
		want    *Tasks
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, s1},
			&Tasks{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/activities/Tasks", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in TasksEndpoint.Delete() returned error: %v, with url /api/v1/{division}/activities/Tasks", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in TasksEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.want)
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.Tasks.Get(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TasksEndpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TasksEndpoint.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTasksEndpoint_New(t *testing.T) {
	s, _, _, teardown := setup()
	defer teardown()
	got := s.Tasks.New()
	want := &Tasks{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TasksEndpoint.New() expected to return %v, got %v", want, got)
	}
}

func TestTasksEndpoint_Create(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *Tasks
	}
	tests := []struct {
		name    string
		args    args
		want    *Tasks
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &Tasks{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&Tasks{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/activities/Tasks", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in TasksEndpoint.Create returned error: %v, with url /api/v1/{division}/activities/Tasks", e)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "POST")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				testBody(t, r, `{"__metadata":{"uri":"https://start.exactonline.nl"}}`+"\n")
				fmt.Fprint(w, `{ "d": { "__metadata": { "uri": "https://start.exactonline.nl"}}}`)
			})

			got, err := s.Tasks.Create(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("TasksEndpoint.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TasksEndpoint.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
