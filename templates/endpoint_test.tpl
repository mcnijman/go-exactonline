// Copyright {{.ServiceEntity.Year}} The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package {{.ServiceEntity.Package}}

{{ if (.HasMethod "GET")}}
func Test{{.EndpointServiceName}}_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

  opts1 := api.NewListOptions()
  opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("{{.URL}}", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in {{.EndpointServiceName}}.List returned error: %v, with url {{.URL}}?$select=*", e)
	}
  api.AddListOptionsToURL(u, opts1)

  opts2 := api.NewListOptions()
  opts2.Select.Add("*")
  opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("{{.URL}}", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in {{.EndpointServiceName}}.List returned error: %v, with url {{.URL}}?$skiptoken=foo", e2)
	}
  api.AddListOptionsToURL(u, opts2)

	{{ if eq .PrimaryProperty.Type "types.GUID"}}
	g := types.NewGUID()
	gs := g.String()
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "` + u2.String() +  `", "results": [{ "{{.PrimaryProperty.Name}}": "`+ gs + `"}]}}`)
		}
	})
	{{ end }}
	{{ if eq .PrimaryProperty.Type "int"}}
	g := 100
	gs := strconv.Itoa(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "` + u2.String() +  `", "results": [{ "{{.PrimaryProperty.Name}}": `+ gs + `}]}}`)
		}
	})
	{{end}}
	{{ if eq .PrimaryProperty.Type "int64"}}
	g := int64(100)
	gs := strconv.Itoa(int(g))
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "` + u2.String() +  `", "results": [{ "{{.PrimaryProperty.Name}}": `+ gs + `}]}}`)
		}
	})
	{{end}}
	{{ if eq .PrimaryProperty.Type "string"}}
	g := "str"
	gs := "str"
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "` + u2.String() +  `", "results": [{ "{{.PrimaryProperty.Name}}": "`+ gs + `"}]}}`)
		}
	})
	{{end}}

	entities, err := s.{{.Name}}.List(context.Background(),{{ if .NeedsDivision}} 0,{{end}} true, opts1)
	if err != nil {
		t.Errorf("{{.EndpointServiceName}}.List returned error: %v", err)
	}

	want := {{"[]"}}*{{.Name}}{{"{{"}}{{.PrimaryProperty.Name}}: &g{{"}}"}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("{{.EndpointServiceName}}.List returned %+v, want %+v", entities, want)
	}
}

func Test{{.EndpointServiceName}}_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

  opts1 := api.NewListOptions()
  opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("{{.URL}}", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in {{.EndpointServiceName}}.List returned error: %v, with url {{.URL}}?$select=*", e)
	}
  api.AddListOptionsToURL(u, opts1)

  opts2 := api.NewListOptions()
  opts2.Select.Add("*")
  opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("{{.URL}}", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in {{.EndpointServiceName}}.List returned error: %v, with url {{.URL}}?$skiptoken=foo", e2)
	}
  api.AddListOptionsToURL(u2, opts2)
	{{ if eq .PrimaryProperty.Type "types.GUID"}}
	g := types.NewGUID()
	gs := g.String()
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "` + u2.String() +  `", "results": [{ "{{.PrimaryProperty.Name}}": "` + gs +  `"}]}}`)
	})
	{{ end }}
	{{ if eq .PrimaryProperty.Type "int"}}
	g := 100
	gs := strconv.Itoa(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "` + u2.String() +  `", "results": [{ "{{.PrimaryProperty.Name}}": ` + gs +  `}]}}`)
	})
	{{end}}
	{{ if eq .PrimaryProperty.Type "int64"}}
	g := int64(100)
	gs := strconv.Itoa(int(g))
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "` + u2.String() +  `", "results": [{ "{{.PrimaryProperty.Name}}": ` + gs +  `}]}}`)
	})
	{{end}}
	{{ if eq .PrimaryProperty.Type "string"}}
	g := "100"
	gs := "100"
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "` + u2.String() +  `", "results": [{ "{{.PrimaryProperty.Name}}": "` + gs +  `"}]}}`)
	})
	{{end}}

	entities, err := s.{{.Name}}.List(context.Background(),{{ if .NeedsDivision}} 0,{{end}} false, opts1)
	if err != nil {
		t.Errorf("{{.EndpointServiceName}}.List returned error: %v", err)
	}

	want := {{"[]"}}*{{.Name}}{{"{{"}}{{.PrimaryProperty.Name}}: &g{{"}}"}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("{{.EndpointServiceName}}.List returned %+v, want %+v", entities, want)
	}
}
{{end}}
