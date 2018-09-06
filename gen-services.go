// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// +build ignore

// This program generates Exact Online API endpoints. It can be invoked by running
// go generate
package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/PuerkitoBio/goquery"
)

const endpointsURL = "https://start.exactonline.nl/docs/HlpRestAPIResources.aspx"

var (
	typesMapping = map[string]string{
		"Binary":   "[]byte",
		"Boolean":  "bool",
		"Byte":     "byte",
		"DateTime": "types.Date",
		"Decimal":  "float64",
		"Double":   "float64",
		"Guid":     "types.GUID",
		"Int16":    "int",
		"Int32":    "int",
		"Int64":    "int64",
		"SByte":    "int",
		"String":   "string",
		"URL":      "types.URL",
	}

	needStringJSON = []struct {
		Service  string
		Endpoint string
		Field    string
	}{
		{"System", "Me", "Legislation"},
		{"System", "Divisions", "Hid"},
	}

	verbose = flag.Bool("v", false, "Print verbose log messages")

	serviceTmpl      = template.Must(template.New("service.tpl").ParseFiles("./templates/service.tpl"))
	serviceTestTmpl  = template.Must(template.New("service_test.tpl").ParseFiles("./templates/service_test.tpl"))
	endpointTmpl     = template.Must(template.New("endpoint.tpl").ParseFiles("./templates/endpoint.tpl"))
	endpointTestTmpl = template.Must(template.New("endpoint_test.tpl").ParseFiles("./templates/endpoint_test.tpl"))
)

type property struct {
	Name            string
	OriginalName    string
	Description     []string
	OriginalType    string
	Type            string
	IsPrimary       bool
	NeedsStringJSON bool
}

type endpoint struct {
	Name          string
	Description   string
	URL           string
	Docs          string
	HasWebhook    bool
	Methods       []string
	Service       string
	IsInBeta      bool
	Properties    []property
	ServiceEntity service
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

func (e endpoint) EndpointServiceName() string {
	return e.Name + "Endpoint"
}

func (e endpoint) FileName() string {
	return toSnake(e.Name) + ".go"
}

func (e endpoint) TestFileName() string {
	return toSnake(e.Name) + "_test.go"
}

func (e endpoint) PrimaryProperty() property {
	for _, p := range e.Properties {
		if p.IsPrimary {
			return p
		}
	}
	return e.Properties[0]
}

func (e endpoint) PackageName() string {
	return strings.Replace(toSnake(e.Service), "_", "", -1)
}

type service struct {
	Endpoints []endpoint
	Package   string
	Name      string
	Year      string
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

func servicesFromEndpoints(endpoints []endpoint) []service {
	m := map[string]service{}

	year := strconv.Itoa(time.Now().UTC().Year())

	for _, e := range endpoints {
		// Drop endpoints for now that don't conform to the default
		if !strings.Contains(e.URL, "/api/v1/") {
			continue
		}

		// Drop endpoints that don't support GET requests for now
		if !strings.Contains(strings.Join(e.Methods, " "), "GET") {
			continue
		}

		var s service
		if se, ok := m[e.Service]; ok {
			s = se
		}

		s.Endpoints = append(s.Endpoints, e)
		s.Name = e.Service
		s.Package = e.PackageName()
		s.Year = year

		m[e.Service] = s
	}

	var services []service

	for _, value := range m {
		services = append(services, value)
	}

	return services
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
		primary, _ := s.Attr("data-key")
		t, _ := s.Attr("data-type")

		nt := "json.RawMessage"
		if ntt, ok := typesMapping[strings.TrimPrefix(t, "Edm.")]; ok {
			nt = ntt
		}

		desc := strings.TrimSpace(s.Parent().Parent().Children().Eq(5).Text())
		properties = append(properties, property{
			OriginalName:    name,
			Name:            strings.Title(name),
			OriginalType:    t,
			Type:            nt,
			Description:     strings.Split(desc, "\n"),
			IsPrimary:       primary == "True", // strings.Contains(desc, "Primary key")
			NeedsStringJSON: needsStringJson(endpoint, strings.Title(name)),
		})
	})
	endpoint.Properties = properties
}

func needsStringJson(e *endpoint, propertyName string) bool {
	for _, n := range needStringJSON {
		if n.Service == e.Service &&
			n.Endpoint == e.Name &&
			n.Field == propertyName {
			return true
		}
	}
	return false
}

func main() {
	flag.Parse()
	endpoints := getEndpointsList()
	services := servicesFromEndpoints(endpoints)

	for _, s := range services {
		logf("------------------------")
		logf("Processing service %v...", s.Name)

		p := path.Join("./services/", s.Package)
		os.MkdirAll(p, os.ModePerm)

		writeFile(serviceTmpl, s, path.Join(p, "service.go"))
		writeFile(serviceTestTmpl, s, path.Join(p, "service_test.go"))

		for _, e := range s.Endpoints {
			logf("  Processing endpoint %v...", e.Name)

			getEndpointProperties(&e)
			e.ServiceEntity = s
			writeFile(endpointTmpl, e, path.Join(p, e.FileName()))
			writeFile(endpointTestTmpl, e, path.Join(p, e.TestFileName()))
		}
	}

	printCopyPasteDeclarations(services)

	// @TODO print copy paste declarations?
}

func writeFile(tpl *template.Template, data interface{}, filePath string) {
	cmd := exec.Command("goimports")
	cmd.Stderr = os.Stdout

	r, w := io.Pipe()
	cmd.Stdin = r

	var b bytes.Buffer
	cmd.Stdout = &b

	err := cmd.Start()
	dontPanic(err)

	err = tpl.Execute(w, data)
	dontPanic(err)

	err = w.Close()
	dontPanic(err)

	err = cmd.Wait()
	dontPanic(err)

	err = ioutil.WriteFile(filePath, b.Bytes(), 0644)
	dontPanic(err)
}

func printCopyPasteDeclarations(services []service) {
	fmt.Println("Service declarations")
	for _, s := range services {
		fmt.Printf("%s *%s.%sService\n", s.Name, s.Package, s.Name)
	}

	fmt.Println("")
	fmt.Println("Service creation")
	for _, s := range services {
		fmt.Printf("c.%s = %s.New%sService(c.client) \n", s.Name, s.Package, s.Name)
	}
}

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
