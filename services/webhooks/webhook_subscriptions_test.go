// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package webhooks

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

func WebhookSubscriptionsPrimaryPropertySample() *types.GUID {
	v := types.NewGUID()
	return &v
}

func WebhookSubscriptionsEntityWithPopulatedPrimaryProperty() *WebhookSubscriptions {
	return &WebhookSubscriptions{ID: WebhookSubscriptionsPrimaryPropertySample()}
}

func WebhookSubscriptionsStringOfPrimaryProperty(v *types.GUID) string {
	return v.String()
}

func WebhookSubscriptionsStringJSONOfPrimaryProperty(v *types.GUID) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestWebhookSubscriptionsEndpoint_GetPrimary(t *testing.T) {
	var want types.GUID
	n := &WebhookSubscriptions{ID: &want}

	if got := n.GetPrimary(); !reflect.DeepEqual(*got, want) {
		t.Errorf("WebhookSubscriptionsEndpoint.GetPrimary() failed, got: %v, want: %v", *got, want)
	}
}

func TestWebhookSubscriptionsEndpoint_UserHasRights(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in WebhookSubscriptionsEndpoint.List returned error: %v", e)
	}

	acceptHeaders := []string{"application/json"}

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))

		q := r.URL.Query()

		if got, want := q.Get("endpoint"), "'webhooks/WebhookSubscriptions'"; got != want {
			t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
		}

		if got, want := q.Get("method"), "GET"; got != want {
			t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
		}

		fmt.Fprint(w, `{ "d": { "UserHasRights": true } }`)
	})

	got, err := s.WebhookSubscriptions.UserHasRights(context.Background(), 0, "GET")
	if err != nil {
		t.Errorf("s.WebhookSubscriptions.UserHasRights should not return an error = %v", err)
	}

	if got != true {
		t.Errorf("s.WebhookSubscriptions.UserHasRights should return true, got: %v", got)
	}
}

func TestWebhookSubscriptionsEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/webhooks/WebhookSubscriptions", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in WebhookSubscriptionsEndpoint.List returned error: %v, with url /api/v1/{division}/webhooks/WebhookSubscriptions?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/webhooks/WebhookSubscriptions", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in WebhookSubscriptionsEndpoint.List returned error: %v, with url /api/v1/{division}/webhooks/WebhookSubscriptions?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := WebhookSubscriptionsPrimaryPropertySample()
	gs := WebhookSubscriptionsStringJSONOfPrimaryProperty(g)

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
		}
	})

	entities, err := s.WebhookSubscriptions.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("WebhookSubscriptionsEndpoint.List returned error: %v", err)
	}

	want := []*WebhookSubscriptions{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("WebhookSubscriptionsEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestWebhookSubscriptionsEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/webhooks/WebhookSubscriptions", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in WebhookSubscriptionsEndpoint.List returned error: %v, with url /api/v1/{division}/webhooks/WebhookSubscriptions", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/webhooks/WebhookSubscriptions", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in WebhookSubscriptionsEndpoint.List returned error: %v, with url /api/v1/{division}/webhooks/WebhookSubscriptions", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := WebhookSubscriptionsPrimaryPropertySample()
	gs := WebhookSubscriptionsStringJSONOfPrimaryProperty(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
	})

	entities, err := s.WebhookSubscriptions.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("WebhookSubscriptionsEndpoint.List returned error: %v", err)
	}

	want := []*WebhookSubscriptions{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("WebhookSubscriptionsEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestWebhookSubscriptionsEndpoint_Get(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	s1 := WebhookSubscriptionsPrimaryPropertySample()
	type args struct {
		ctx      context.Context
		division int
		id       *types.GUID
	}
	tests := []struct {
		name    string
		args    args
		want    *WebhookSubscriptions
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, s1},
			&WebhookSubscriptions{ID: s1, MetaData: &api.MetaData{URI: &types.URL{&url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/webhooks/WebhookSubscriptions", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in WebhookSubscriptionsEndpoint.Delete() returned error: %v, with url /api/v1/{division}/webhooks/WebhookSubscriptions", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in WebhookSubscriptionsEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.want)
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.WebhookSubscriptions.Get(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("WebhookSubscriptionsEndpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WebhookSubscriptionsEndpoint.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWebhookSubscriptionsEndpoint_New(t *testing.T) {
	s, _, _, teardown := setup()
	defer teardown()
	got := s.WebhookSubscriptions.New()
	want := &WebhookSubscriptions{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WebhookSubscriptionsEndpoint.New() expected to return %v, got %v", want, got)
	}
}

func TestWebhookSubscriptionsEndpoint_Create(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *WebhookSubscriptions
	}
	tests := []struct {
		name    string
		args    args
		want    *WebhookSubscriptions
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &WebhookSubscriptions{MetaData: &api.MetaData{URI: &types.URL{&url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&WebhookSubscriptions{MetaData: &api.MetaData{URI: &types.URL{&url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/webhooks/WebhookSubscriptions", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in WebhookSubscriptionsEndpoint.Create returned error: %v, with url /api/v1/{division}/webhooks/WebhookSubscriptions", e)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "POST")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				testBody(t, r, `{"__metadata":{"uri":"https://start.exactonline.nl"}}`+"\n")
				fmt.Fprint(w, `{ "d": { "__metadata": { "uri": "https://start.exactonline.nl"}}}`)
			})

			got, err := s.WebhookSubscriptions.Create(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("WebhookSubscriptionsEndpoint.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WebhookSubscriptionsEndpoint.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWebhookSubscriptionsEndpoint_Update(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *WebhookSubscriptions
	}
	s1 := WebhookSubscriptionsPrimaryPropertySample()
	tests := []struct {
		name    string
		args    args
		want    *WebhookSubscriptions
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &WebhookSubscriptions{ID: s1, MetaData: &api.MetaData{URI: &types.URL{&url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&WebhookSubscriptions{ID: s1, MetaData: &api.MetaData{URI: &types.URL{&url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/webhooks/WebhookSubscriptions", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in WebhookSubscriptionsEndpoint.Update returned error: %v, with url /api/v1/{division}/webhooks/WebhookSubscriptions", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.entity.GetPrimary())
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in WebhookSubscriptionsEndpoint.Update returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "PUT")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.args.entity)
				testBody(t, r, string(b)+"\n")
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.WebhookSubscriptions.Update(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("WebhookSubscriptionsEndpoint.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WebhookSubscriptionsEndpoint.Update() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestWebhookSubscriptionsEndpoint_Delete(t *testing.T) {
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
			args{context.Background(), 0, WebhookSubscriptionsPrimaryPropertySample()},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/webhooks/WebhookSubscriptions", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in WebhookSubscriptionsEndpoint.Delete() returned error: %v, with url /api/v1/{division}/webhooks/WebhookSubscriptions", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in WebhookSubscriptionsEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "DELETE")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testBody(t, r, "")
				w.WriteHeader(http.StatusNoContent)
			})

			err := s.WebhookSubscriptions.Delete(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("WebhookSubscriptionsEndpoint.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
