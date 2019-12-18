// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/sebas7dk/go-exactonline/types"
)

const (
	contentType  = "application/json"
	acceptHeader = "application/json"
	userAgent    = "github.com/sebas7dk/go-exactonline"
)

// A Client manages communication with the Exact Online API.
type Client struct {
	client *http.Client

	// BaseURL for API requests. Defaults to the Dutch API. See more available base urls in the API documentation. @TODO
	BaseURL *url.URL

	// UserAgent used when communicating with the Exact Online API.
	UserAgent string
}

// NewClient returns a new Exact Online API client. Provide a http.Client that
// will perform the authentication for you (such as that provided by the
// golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client, url string) (*Client, error) {
	if httpClient == nil {
		return fmt.Errorf("A valid http client is required")
	}

	baseURL, err := url.Parse(url) // #nosec
	if err != nil {
		return nil, err
	}

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	return c, nill
}

// ResolveURL will either return either a resolved path or a valid absolute URI
func (c *Client) ResolveURL(urlStr string) (*url.URL, error) {
	if abs, err := url.Parse(urlStr); err == nil && abs.IsAbs() {
		return abs, nil
	}

	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	return c.BaseURL.Parse(urlStr)
}

// ResolvePathWithDivision will resolve the base url for paths that need a division prefix
func (c *Client) ResolvePathWithDivision(path string, division int) (*url.URL, error) {
	s := strings.Replace(path, "{division}", strconv.Itoa(division), 1)
	return c.ResolveURL(s)
}

// NewRequest creates an API request. An absolute URL must be provided in url.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, parseErr := c.ResolveURL(urlStr)
	if parseErr != nil {
		return nil, parseErr
	}
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return req, err
	}
	if body != nil {
		req.Header.Set("Content-Type", contentType)
	}
	req.Header.Set("Accept", acceptHeader)
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req = req.WithContext(ctx)
	res, err := c.client.Do(req) // #nosec G107
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}
	defer res.Body.Close()
	response := &Response{Response: res}

	err = checkResponse(res, req.URL.String())
	if err != nil {
		return response, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil || len(body) == 0 {
		return response, err
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, copyErr := io.Copy(w, res.Body)
			if copyErr != nil {
				err = copyErr
			}
		} else {
			decErr := json.Unmarshal(response.Data, v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return response, err
}

// NewRequestAndDo combines NewRequest and Do methods
func (c *Client) NewRequestAndDo(ctx context.Context, method, urlStr string, body, v interface{}) (*http.Request, *Response, error) {
	req, e := c.NewRequest(method, urlStr, body)
	if e != nil {
		return req, nil, e
	}
	res, err := c.Do(ctx, req, v) // #nosec G107
	return req, res, err
}

// ListRequestAndDoAll requests all paginated pages ListRequestAndDo
func (c *Client) ListRequestAndDoAll(ctx context.Context, urlStr string, v interface{}) error {
	var s []json.RawMessage
	_, f, err := c.NewRequestAndDo(ctx, "GET", urlStr, nil, &s)
	if err != nil {
		return err
	}

	var next = f.NextPage
	for next != nil {
		var i []json.RawMessage
		_, l, rErr := c.NewRequestAndDo(ctx, "GET", next.String(), nil, &i)
		if rErr != nil {
			return rErr
		}
		s = append(s, i...)
		next = l.NextPage
	}

	err = unmarshalRawMessages(s, v)
	return err
}

// UserHasRights checks for the given endpoints if the user has permissions to request that method
// at that endpoint.
func (c *Client) UserHasRights(ctx context.Context, division int, endpoint, method string) (bool, error) {
	u, _ := c.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", division) // #nosec
	q := u.Query()
	q.Add("endpoint", fmt.Sprintf("'%s'", endpoint))
	q.Add("method", method)
	u.RawQuery = q.Encode()

	v := struct {
		UserHasRights bool `json:"UserHasRights,"`
	}{}
	_, _, err := c.NewRequestAndDo(ctx, "GET", u.String(), nil, &v)

	return v.UserHasRights, err
}

func unmarshalRawMessages(m []json.RawMessage, v interface{}) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, v)
	return err
}

func checkResponse(r *http.Response, u string) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		if err := json.Unmarshal(data, errorResponse); err != nil {
			return fmt.Errorf("%s: for %s, also encountered an error "+
				"Unmarshalling the error response", r.Status, u)
		}
	}

	return errorResponse
}

// AddOdataKeyToURL will add the odata getter key to the url path
func AddOdataKeyToURL(u *url.URL, k interface{}) (*url.URL, error) {
	v := reflect.ValueOf(k)

	if k == nil ||
		v == reflect.Zero(reflect.TypeOf(k)) ||
		(v.Kind() == reflect.Ptr && v.IsNil()) {
		return nil, errors.New("Cannot add Nil value to URL")
	}

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Type().String() == "types.GUID" {
		g := v.Interface().(types.GUID)
		return u.Parse(fmt.Sprintf("%s(guid'%s')", u.Path, g.String()))
	}

	return u.Parse(fmt.Sprintf("%s(%v)", u.Path, v.Interface()))
}
