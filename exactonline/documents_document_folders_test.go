// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestDocumentsDocumentFoldersService_List_all(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	u, e := client.ResolvePathWithDivision("/api/v1/{division}/documents/DocumentFolders?$select=*", 0)
	u2, e := client.ResolvePathWithDivision("/api/v1/{division}/documents/DocumentFolders?$skiptoken=foo", 0)
	if e != nil {
		t.Errorf("client.ResolvePathWithDivision in DocumentsDocumentFoldersService.List returned error: %v, with url /api/v1/{division}/documents/DocumentFolders", e)
	}

	g := NewGUID()
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

	/* g := NewGUID()
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "` + u2.String() + `", "results": [{ "ID": "` + g.String() + `"}]}}`)
		}
	}) */

	entities, err := client.DocumentsDocumentFolders.List(context.Background(), 0, true)
	if err != nil {
		t.Errorf("DocumentsDocumentFoldersService.List returned error: %v", err)
	}

	want := []*DocumentsDocumentFolders{{ID: &g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("DocumentsDocumentFoldersService.List returned %+v, want %+v", entities, want)
	}
}

func TestDocumentsDocumentFoldersService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	u, e := client.ResolvePathWithDivision("/api/v1/{division}/documents/DocumentFolders?$select=*", 0)
	u2, e := client.ResolvePathWithDivision("/api/v1/{division}/documents/DocumentFolders?$skiptoken=foo", 0)
	if e != nil {
		t.Errorf("client.ResolvePathWithDivision in DocumentsDocumentFoldersService.List returned error: %v, with url /api/v1/{division}/documents/DocumentFolders", e)
	}

	g := NewGUID()
	gs := g.String()
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": "`+gs+`"}]}}`)
	})

	entities, err := client.DocumentsDocumentFolders.List(context.Background(), 0, false)
	if err != nil {
		t.Errorf("DocumentsDocumentFoldersService.List returned error: %v", err)
	}

	want := []*DocumentsDocumentFolders{{ID: &g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("DocumentsDocumentFoldersService.List returned %+v, want %+v", entities, want)
	}
}