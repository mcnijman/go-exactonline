// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"testing"

	"github.com/mcnijman/go-exactonline/types"
)

// setup sets up a test HTTP server along with a exactonline.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	// We want to ensure that tests catch mistakes where the endpoint URL is
	// specified as absolute rather than relative. It only makes a difference
	// when there's a non-empty base URL path. So, use that. See issue #752.
	apiHandler := http.NewServeMux()
	apiHandler.Handle("/", mux)

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(apiHandler)

	// client is the GitHub client being tested and is
	// configured to use test server.
	client = NewClient(nil)
	url, _ := url.Parse(server.URL + "/")
	client.BaseURL = url

	return client, mux, server.URL, server.Close
}

func TestNewClient(t *testing.T) {
	c := NewClient(nil)

	if got, want := c.BaseURL.String(), defaultBaseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}
	if got, want := c.UserAgent, userAgent; got != want {
		t.Errorf("NewClient UserAgent is %v, want %v", got, want)
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient(nil)

	type T struct {
		GUID *types.GUID `json:","`
	}

	g := types.NewGUID()

	inURL, outURL := "/foo", defaultBaseURL+"foo"
	inBody, outBody := &T{GUID: &g}, `{"GUID":"`+g.String()+`"}`+"\n"
	req, e := c.NewRequest("GET", inURL, inBody)
	if e != nil {
		t.Errorf("NewRequest() error: %v", e)
	}

	// test that relative URL was expanded
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}

	// test that body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("NewRequest(%q) Body is %v, want %v", inBody, got, want)
	}

	// test that default user-agent is attached to the request
	if got, want := req.Header.Get("User-Agent"), c.UserAgent; got != want {
		t.Errorf("NewRequest() User-Agent is %v, want %v", got, want)
	}
}

func TestNewRequest_invalidJSON(t *testing.T) {
	c := NewClient(nil)

	type T struct {
		A map[interface{}]interface{}
	}
	_, err := c.NewRequest("GET", ".", &T{})

	if err == nil {
		t.Error("Expected error to be returned.")
	}
	if err, ok := err.(*json.UnsupportedTypeError); !ok {
		t.Errorf("Expected a JSON error; got %#v.", err)
	}
}

func TestNewRequest_badURL(t *testing.T) {
	c := NewClient(nil)
	_, err := c.NewRequest("GET", ":", nil)
	testURLParseError(t, err)
}

func TestNewRequest_emptyUserAgent(t *testing.T) {
	c := NewClient(nil)
	c.UserAgent = ""
	req, err := c.NewRequest("GET", ".", nil)
	if err != nil {
		t.Fatalf("NewRequest returned unexpected error: %v", err)
	}
	if _, ok := req.Header["User-Agent"]; ok {
		t.Fatal("constructed request contains unexpected User-Agent header")
	}
}

func TestNewRequest_emptyBody(t *testing.T) {
	c := NewClient(nil)
	req, err := c.NewRequest("GET", ".", nil)
	if err != nil {
		t.Fatalf("NewRequest returned unexpected error: %v", err)
	}
	if req.Body != nil {
		t.Fatalf("constructed request contains a non-nil Body")
	}
}

func TestNewRequest_errorForNoTrailingSlash(t *testing.T) {
	tests := []struct {
		rawurl    string
		wantError bool
	}{
		{rawurl: "https://example.com/api/v3", wantError: true},
		{rawurl: "https://example.com/api/v3/", wantError: false},
	}
	c := NewClient(nil)
	for _, test := range tests {
		u, err := url.Parse(test.rawurl)
		if err != nil {
			t.Fatalf("url.Parse returned unexpected error: %v.", err)
		}
		c.BaseURL = u
		if _, err := c.NewRequest(http.MethodGet, "test", nil); test.wantError && err == nil {
			t.Fatalf("Expected error to be returned.")
		} else if !test.wantError && err != nil {
			t.Fatalf("NewRequest returned unexpected error: %v.", err)
		}
	}
}

func TestDo(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	type foo struct {
		A string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{ "d": {"A":"a"} }`)
	})

	req, _ := client.NewRequest("GET", ".", nil)
	body := new(foo)
	_, err := client.Do(context.Background(), req, body)
	if err != nil {
		t.Errorf("Response body received an error = %v", err)
	}

	want := &foo{"a"}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Response body = %v, want %v", body, want)
	}
}

func TestDo_httpError(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	req, _ := client.NewRequest("GET", ".", nil)
	resp, err := client.Do(context.Background(), req, nil)

	if err == nil {
		t.Fatal("Expected HTTP 400 error, got no error.")
	}

	if resp.Response.StatusCode != 400 {
		t.Errorf("Expected HTTP 400 error, got %d status code.", resp.Response.StatusCode)
	}
}

func TestDo_redirectLoop(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	})

	req, _ := client.NewRequest("GET", ".", nil)
	_, err := client.Do(context.Background(), req, nil)

	if err == nil {
		t.Error("Expected error to be returned.")
	}
	if err, ok := err.(*url.Error); !ok {
		t.Errorf("Expected a URL error; got %#v.", err)
	}
}

func TestDo_noContent(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	var body json.RawMessage

	req, _ := client.NewRequest("GET", ".", nil)
	_, err := client.Do(context.Background(), req, &body)
	if err != nil {
		t.Fatalf("Do returned unexpected error: %v", err)
	}
}

func Test_handleResponseError500(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"error": {"code": "","message": {"lang": "","value": "Can't delete: Account 58 - Used in: Administrations"}}}`)
	})

	var body json.RawMessage

	req, _ := client.NewRequest("GET", ".", nil)
	_, err := client.Do(context.Background(), req, &body)
	if err == nil {
		t.Fatal("Do expected an error")
	}
}

func Test_handleResponseError400(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	errorCodes := []int{
		http.StatusNotFound,
		http.StatusForbidden,
		http.StatusBadRequest,
		http.StatusUnauthorized,
	}

	for _, code := range errorCodes {
		mux.HandleFunc("/"+strconv.Itoa(code), func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(code)
			fmt.Fprint(w, ``)
		})

		var body json.RawMessage

		req, _ := client.NewRequest("GET", "/"+strconv.Itoa(code), nil)
		_, err := client.Do(context.Background(), req, &body)
		if err == nil {
			t.Fatal("Do expected an error")
		}
	}
}

/* func Test_DoContextError(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		fmt.Fprint(w, `test`)
	})

	var body json.RawMessage

	ctxe := context.Background()
	ctx, cancel := context.WithTimeout(ctxe, 1*time.Microsecond)

	req, _ := client.NewRequest("GET", ".", nil)
	_, err := client.Do(ctx, req, &body)

	if err == nil {
		t.Fatal("Do expected an error")
	}

	if !reflect.DeepEqual(err, ctx.Err()) {
		t.Fatalf("err should be a context error, error = %+v, ctx.Err() = %+v", err, ctx.Err())
	}
	cancel()
} */

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %q, want %q", header, got, want)
	}
}

func testURLParseError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func testBody(t *testing.T, r *http.Request, want string) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Errorf("Error reading request body: %v", err)
	}
	if got := string(b); got != want {
		t.Errorf("request Body is %s, want %s", got, want)
	}
}

// Helper function to test that a value is marshalled to JSON as expected.
func testJSONMarshal(t *testing.T, v interface{}, want string) {
	j, err := json.Marshal(v)
	if err != nil {
		t.Errorf("Unable to marshal JSON for %v", v)
	}

	w := new(bytes.Buffer)
	err = json.Compact(w, []byte(want))
	if err != nil {
		t.Errorf("String is not valid json: %s", want)
	}

	if w.String() != string(j) {
		t.Errorf("json.Marshal(%q) returned %s, want %s", v, j, w)
	}

	// now go the other direction and make sure things unmarshal as expected
	u := reflect.ValueOf(v).Interface()
	if err := json.Unmarshal([]byte(want), u); err != nil {
		t.Errorf("Unable to unmarshal JSON for %v", want)
	}

	if !reflect.DeepEqual(v, u) {
		t.Errorf("json.Unmarshal(%q) returned %s, want %s", want, u, v)
	}
}

func TestClient_ResolvePathWithDivision(t *testing.T) {
	type fields struct {
		client    *http.Client
		BaseURL   *url.URL
		UserAgent string
	}
	type args struct {
		path     string
		division int
	}
	baseURL, _ := url.Parse("https://start.exactonline.nl/")
	u1, _ := url.Parse("https://start.exactonline.nl/foo-1000")
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *url.URL
		wantErr bool
	}{
		{"1", fields{http.DefaultClient, baseURL, "test useragen"}, args{"/foo-{division}", 1000}, u1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:    tt.fields.client,
				BaseURL:   tt.fields.BaseURL,
				UserAgent: tt.fields.UserAgent,
			}
			got, err := c.ResolvePathWithDivision(tt.args.path, tt.args.division)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ResolvePathWithDivision() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ResolvePathWithDivision() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddOdataKeyToURL(t *testing.T) {
	type args struct {
		u *url.URL
		k interface{}
	}

	g := types.NewGUID()
	in, _ := url.Parse("https://start.exactonline.nl/foo/bar")
	out, _ := url.Parse(fmt.Sprintf("https://start.exactonline.nl/foo/bar(guid'%s')", g.String()))
	in2, _ := url.Parse("https://start.exactonline.nl/foo/bar")
	out2, _ := url.Parse(fmt.Sprintf("https://start.exactonline.nl/foo/bar(guid'%s')", g.String()))

	i := 100
	in3, _ := url.Parse("https://start.exactonline.nl/foo/bar")
	out3, _ := url.Parse("https://start.exactonline.nl/foo/bar(100)")
	in4, _ := url.Parse("https://start.exactonline.nl/foo/bar")
	out4, _ := url.Parse("https://start.exactonline.nl/foo/bar(100)")

	tests := []struct {
		name    string
		args    args
		want    *url.URL
		wantErr bool
	}{
		{"1", args{in, g}, out, false},
		{"2", args{in2, &g}, out2, false},
		{"3", args{in3, i}, out3, false},
		{"4", args{in4, &i}, out4, false},
		{"5", args{nil, nil}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddOdataKeyToURL(tt.args.u, tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddOdataKeyToURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddOdataKeyToURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_UserHasRights(t *testing.T) {
	tests := []struct {
		name      string
		hasRights bool
		endpoint  string
		method    string
		wantErr   bool
	}{
		{"1", true, "subscription/SubscriptionLineTypes", "GET", false},
		{"2", false, "subscription/SubscriptionLineTypes", "POST", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, mux, _, teardown := setup()
			defer teardown()
			u, e := c.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
			if e != nil {
				t.Errorf("c.ResolvePathWithDivision in c.UserHasRights returned error: %v, with url /api/v1/{division}/users/UserHasRights", e)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", acceptHeader)
				q := r.URL.Query()
				if got, want := q.Get("endpoint"), fmt.Sprintf("'%s'", tt.endpoint); got != want {
					t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
				}
				if got, want := q.Get("method"), tt.method; got != want {
					t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
				}
				fmt.Fprint(w, `{ "d": { "UserHasRights": `+strconv.FormatBool(tt.hasRights)+` } }`)
			})

			got, err := c.UserHasRights(context.Background(), 0, tt.endpoint, tt.method)
			if err != nil && !tt.wantErr {
				t.Errorf("s.SubscriptionLineTypes.UserHasRights should not return an error = %v", err)
			}

			if got != tt.hasRights {
				t.Errorf("s.SubscriptionLineTypes.UserHasRights should return true, got: %v", got)
			}
		})
	}
}
