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
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	contentType    = "application/json"
	acceptHeader   = "application/json"
	defaultBaseURL = "https://start.exactonline.nl/"
	userAgent      = "github.com/mcnijman/go-exactonline"
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
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL) // #nosec

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	return c
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
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)
	resp, err := c.client.Do(req) // #nosec G107
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
	defer resp.Body.Close()

	if err := handleResponseError(resp, req.URL.String()); err != nil {
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, copyErr := io.Copy(w, resp.Body)
			if copyErr != nil {
				err = copyErr
			}
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return resp, err
}

// NewRequestAndDo combines NewRequest and Do methods
func (c *Client) NewRequestAndDo(ctx context.Context, method, urlStr string, body, v interface{}) (*http.Request, *http.Response, error) {
	req, e := c.NewRequest(method, urlStr, body)
	if e != nil {
		return req, nil, e
	}
	res, err := c.Do(ctx, req, v) // #nosec G107
	return req, res, err
}

// ListRequestAndDo combines NewRequestAndDo and unmarshalls in general ListResponse
func (c *Client) ListRequestAndDo(ctx context.Context, urlStr string, v interface{}) (*ListResponse, *http.Request, *http.Response, error) {
	var listResponse ListResponse
	req, res, err := c.NewRequestAndDo(ctx, "GET", urlStr, nil, &listResponse)
	if err != nil {
		return nil, nil, nil, err
	}

	if v != nil {
		err = json.Unmarshal(listResponse.Data.Results, v)
	}

	return &listResponse, req, res, err
}

// ListRequestAndDoAll requests all paginated pages ListRequestAndDo
func (c *Client) ListRequestAndDoAll(ctx context.Context, urlStr string, v interface{}) error {
	var s []json.RawMessage
	f, _, _, err := c.ListRequestAndDo(ctx, urlStr, &s)
	if err != nil {
		return err
	}

	var next = f.Data.Next
	for next != "" {
		var i []json.RawMessage
		l, _, _, rErr := c.ListRequestAndDo(ctx, next, &i)
		if rErr != nil {
			return rErr
		}
		s = append(s, i...)
		next = l.Data.Next
	}

	err = unmarshalRawMessages(s, v)
	return err
}

func unmarshalRawMessages(m []json.RawMessage, v interface{}) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, v)
	return err
}

func handleResponseError(r *http.Response, u string) error {
	if r.StatusCode == http.StatusInternalServerError {
		var e InternalServerErrorResponse
		f := json.NewDecoder(r.Body).Decode(e)
		if f != nil {
			return fmt.Errorf("%s: ListRequestAndDo for %s, also encountered an error "+
				"Unmarshalling the error response", r.Status, u)
		}
		return fmt.Errorf("%s: ListRequestAndDo for %s, with message %s", r.Status,
			u, e.Error.Message.Value)
	}

	if r.StatusCode == http.StatusBadRequest || r.StatusCode == http.StatusUnauthorized ||
		r.StatusCode == http.StatusForbidden || r.StatusCode == http.StatusNotFound {
		return fmt.Errorf("%s: ListRequestAndDo for %s", r.Status, u)
	}

	return nil
}
