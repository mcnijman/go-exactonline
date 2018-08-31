// +build ignore

// This program generates Exact Online API endpoints. It can be invoked by running
// go generate
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
)

const endpointsURL = "https://start.exactonline.nl/docs/HlpRestAPIResources.aspx"

var (
	typesMapping = map[string]string{
		"Binary":   "[]byte",
		"Boolean":  "bool",
		"Byte":     "byte",
		"DateTime": "Date",
		"Decimal":  "float64",
		"Double":   "float64",
		"Guid":     "GUID",
		"Int16":    "int",
		"Int32":    "int",
		"Int64":    "int64",
		"SByte":    "int",
		"String":   "string",
		"URL":      "url.URL",
	}
	verbose = flag.Bool("v", false, "Print verbose log messages")

	sourceTmpl     = template.Must(template.New("source").Parse(source))
	sourceTmplTest = template.Must(template.New("source").Parse(test))
)

type property struct {
	Name         string
	OriginalName string
	Description  []string
	OriginalType string
	Type         string
	IsPrimary    bool
}

type endpoint struct {
	Name        string
	Description string
	URL         string
	Docs        string
	HasWebhook  bool
	Methods     []string
	Service     string
	IsInBeta    bool
	Properties  []property
}

func (e endpoint) HasMethod(method string) bool {
	for _, m := range e.Methods {
		if m == method {
			return true
		}
	}
	return false
}

func (e endpoint) NeedsDivision() bool {
	return strings.Contains(e.URL, "{division}")
}

func (e endpoint) EntityName() string {
	return e.Service + e.Name
}

func (e endpoint) ServiceName() string {
	return e.EntityName() + "Service"
}

func (e endpoint) PrimaryProperty() property {
	for _, p := range e.Properties {
		if p.IsPrimary {
			return p
		}
	}
	return e.Properties[0]
}

type templateData struct {
	filename string
	Year     int
	Package  string
	Imports  map[string]string
	Endpoint endpoint
}

func getEndpointsList() []endpoint {
	logf("Fetching endpoints list %v...", endpointsURL)
	res, err := http.Get(endpointsURL)
	dontPanic(err)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("getEndpointsList: status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	dontPanic(err)

	var endpoints []endpoint

	doc.Find("tr.filter").Each(func(i int, s *goquery.Selection) {
		class, _ := s.Attr("class")
		c := s.Children()
		docs, _ := c.Eq(1).Find("a").Attr("href")
		endpoints = append(endpoints, endpoint{
			Name:       strings.Replace(c.Eq(1).Find("a").Text(), "/", "", -1),
			URL:        c.Eq(2).Text(),
			Docs:       docs,
			HasWebhook: strings.Contains(class, "HasWebhook"),
			Methods:    strings.Split(c.Eq(3).Text(), ", "),
			Service:    c.Eq(0).Text(),
			IsInBeta:   strings.Contains(c.Eq(1).Text(), "BETA"),
		})
	})

	return endpoints
}

func getEndpointProperties(endpoint *endpoint) {
	uri := "https://start.exactonline.nl/docs/" + endpoint.Docs
	res, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("%s: status code error: %d %s", uri, res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	dontPanic(err)

	var properties []property
	doc.Find(".showget").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		prim, _ := s.Attr("data-key")
		t, _ := s.Attr("data-type")

		nt := "[]byte"
		if ntt, ok := typesMapping[strings.TrimPrefix(t, "Edm.")]; ok {
			nt = ntt
		}

		desc := strings.TrimSpace(s.Parent().Parent().Children().Eq(5).Text())
		properties = append(properties, property{
			OriginalName: name,
			Name:         strings.Title(name),
			OriginalType: t,
			Type:         nt,
			Description:  strings.Split(desc, "\n"),
			IsPrimary:    prim == "True", // strings.Contains(desc, "Primary key")
		})
	})
	endpoint.Properties = properties
}

func printCopyPasteDeclarations(endpoints []endpoint) {
	fmt.Println("Service declarations")
	for _, e := range endpoints {
		fmt.Printf("%s *%s \n", e.EntityName(), e.ServiceName())
	}

	fmt.Println("")
	fmt.Println("Service creation")
	for _, e := range endpoints {
		fmt.Printf("c.%s = (*%s)(&c.common) \n", e.EntityName(), e.ServiceName())
	}
}

func main() {
	flag.Parse()
	endpoints := getEndpointsList()
	var filtered []endpoint

	for _, e := range endpoints {
		logf("Processing %v...", e.Name)

		// Drop endpoints for now that don't conform to the default
		if !strings.Contains(e.URL, "/api/v1/") {
			continue
		}

		// Drop endpoints that don't support GET requests for now
		if !strings.Contains(strings.Join(e.Methods, " "), "GET") {
			continue
		}

		getEndpointProperties(&e)

		filtered = append(filtered, e)

		s := &templateData{
			filename: toSnake(e.EntityName()) + ".go",
			Year:     2018,
			Package:  "exactonline",
			Endpoint: e,
			Imports: map[string]string{
				"1": "context",
				// "2": "fmt",
			},
		}

		var buf bytes.Buffer
		err := sourceTmpl.Execute(&buf, s)
		dontPanic(err)

		if e.Name == "AccountDocumentsCount" {
			logf("%+v \n", e.PrimaryProperty())
			fmt.Print(string(buf.Bytes()))
			logf("%+v \n", e)
		}

		clean, err := format.Source(buf.Bytes())
		dontPanic(err)

		logf("Writing %v...", s.filename)
		err = ioutil.WriteFile(s.filename, clean, 0644)
		dontPanic(err)

		/* t := &templateData{
			filename: toSnake(e.EntityName()) + "_test.go",
			Year:     2018,
			Package:  "exactonline",
			Endpoint: e,
			Imports: map[string]string{
				"1": "testing",
				"2": "net/http",
				"3": "strings",
				"4": "context",
				"5": "reflect",
				"6": "fmt",
			},
		}

		var buf2 bytes.Buffer
		err = sourceTmplTest.Execute(&buf2, t)
		dontPanic(err)

		clean2, err := format.Source(buf2.Bytes())
		dontPanic(err)

		logf("Writing %v...", t.filename)
		err = ioutil.WriteFile(t.filename, clean2, 0644)
		dontPanic(err) */
	}

	printCopyPasteDeclarations(filtered)
}

const source = `
// Copyright {{.Year}} The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package {{.Package}}
{{with .Imports}}
import (
  {{- range . -}}
  "{{.}}"
  {{end -}}
)
{{end}}

{{with .Endpoint}}

// {{.ServiceName}} is responsible for communicating with
// the {{.Name}} endpoint of the {{.Service}} service.
type {{.ServiceName}} service

// {{.EntityName}}: {{.Description}}
// Service: {{.Service}}
// Entity: {{.Name}}
// URL: {{.URL}}
// HasWebhook: {{.HasWebhook}}
// IsInBeta: {{.IsInBeta}}
// Methods: {{range .Methods}}{{.}} {{end}}
// Endpoint docs: https://start.exactonline.nl/docs/{{.Docs}}
type {{.EntityName}} struct {
{{- range .Properties -}}
	{{if .Name}}
	// {{.Name}}:
	{{- range $i, $e := .Description -}}
	{{if $i}} //{{end}} {{ $e }}
	{{- end}}
	{{.Name}} *{{.Type}}  ` + "`" + `json:"{{.OriginalName}},omitempty"` + "`" + `
	{{end -}}
{{end}}
}

func (s *{{.EntityName}}) GetIdentifier() {{.PrimaryProperty.Type}} {
	return *s.{{.PrimaryProperty.Name}}
}

{{ if (.HasMethod "GET")}}
// List the {{.Name}} entities{{ if .NeedsDivision }} in the provided divison{{end}}.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *{{.ServiceName}}) List(ctx context.Context, division int, all bool) ([]*{{.Service}}{{.Name}}, error) {
	var entities []*{{.Service}}{{.Name}}
	u, err := s.client.ResolvePathWithDivision("{{.URL}}?$select=*", division)
	if err != nil {
		return nil, err
	}
	if all {
		err = s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err = s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}

/* // Get the {{.Name}} enitity, by {{.PrimaryProperty.Name}}.
func (s *{{.ServiceName}}) Get(ctx context.Context, division int, id {{.PrimaryProperty.Type}}) (*{{.Service}}{{.Name}}, error) {
	var entities []*{{.Service}}{{.Name}}
	u, err := s.client.ResolvePathWithDivision("{{.URL}}?$select=*", division)
	if err != nil {
		return nil, err
	}

	if _, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities); err != nil {
		return nil, err
	}

	if len(entities) != 1 {
		return nil, fmt.Errorf("Returned %d {{.Name}} entities, expected 1", len(entities))
	}

	return entities[0], nil
} */
{{end}}
{{end}}
`

const test = `
// Copyright {{.Year}} The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package {{.Package}}
{{with .Imports}}
import (
  {{- range . -}}
  "{{.}}"
  {{end -}}
)
{{end}}

{{with .Endpoint}}
{{ if (.HasMethod "GET")}}
func Test{{.ServiceName}}_List_all(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	u, e := client.ResolvePathWithDivision("{{.URL}}?$select=*", 0)
	u2, e := client.ResolvePathWithDivision("{{.URL}}?$skiptoken=foo", 0)
	if e != nil {
		t.Errorf("client.ResolvePathWithDivision in {{.ServiceName}}.List returned error: %v, with url {{.URL}}", e)
	}
	g := NewGUID()
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, ` + "`" + `{ "d": { "__next": "", "results": []}}` + "`" + `)
		} else {
			fmt.Fprint(w, ` + "`" + `{ "d": { "__next": "` + "` + u2.String() + `" + `", "results": [{ "{{.PrimaryProperty.Name}}": "` + "` + g.String() + `" + `"}]}}` + "`" + `)
		}
	})

	entities, err := client.{{.EntityName}}.List(context.Background(), 0, true)
	if err != nil {
		t.Errorf("{{.ServiceName}}.List returned error: %v", err)
	}

	want := {{"[]"}}*{{.EntityName}}{{"{{"}}{{.PrimaryProperty.Name}}: &g{{"}}"}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("{{.ServiceName}}.List returned %+v, want %+v", entities, want)
	}
}

func Test{{.ServiceName}}_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	u, e := client.ResolvePathWithDivision("{{.URL}}?$select=*", 0)
	u2, e := client.ResolvePathWithDivision("{{.URL}}?$skiptoken=foo", 0)
	if e != nil {
		t.Errorf("client.ResolvePathWithDivision in {{.ServiceName}}.List returned error: %v, with url {{.URL}}", e)
	}
	g := NewGUID()
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, ` + "`" + `{ "d": { "__next": "` + "` + u2.String() + `" + `", "results": [{ "{{.PrimaryProperty.Name}}": "` + "` + g.String() + `" + `"}]}}` + "`" + `)
	})

	entities, err := client.{{.EntityName}}.List(context.Background(), 0, false)
	if err != nil {
		t.Errorf("{{.ServiceName}}.List returned error: %v", err)
	}

	want := {{"[]"}}*{{.EntityName}}{{"{{"}}{{.PrimaryProperty.Name}}: &g{{"}}"}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("{{.ServiceName}}.List returned %+v, want %+v", entities, want)
	}
}
{{end}}
{{end}}
`

/* const services = `
// Copyright {{.Year}} The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
{{- range .Endpoints -}}
{{.ServiceName}} {{}}
{{end}}
` */

func toSnake(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}

func logf(fmt string, args ...interface{}) {
	if *verbose {
		log.Printf(fmt, args...)
	}
}

func dontPanic(e error) {
	if e != nil {
		panic(e)
	}
}
