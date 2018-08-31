// go:generate go run gen-services.go -v

package exactonline

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
	"sync"

	"golang.org/x/oauth2"
)

const (
	defaultBaseURL = "https://start.exactonline.nl/"
	userAgent      = "go-exactonline"
)

// A Client manages communication with the Exact Online API.
type Client struct {
	clientMu sync.Mutex
	client   *http.Client

	// BaseURL for API requests. Defaults to the Dutch API. See more available base urls in the API documentation. @TODO
	BaseURL *url.URL

	// UserAgent used when communicating with the Exact Online API.
	UserAgent string

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the Exact Online API
	AccountancyAccountInvolvedAccounts *AccountancyAccountInvolvedAccountsService
	AccountancyAccountOwners           *AccountancyAccountOwnersService
}

type service struct {
	client *Client
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
	c.common.client = c
	c.AccountancyAccountInvolvedAccounts = (*AccountancyAccountInvolvedAccountsService)(&c.common)
	c.AccountancyAccountOwners = (*AccountancyAccountOwnersService)(&c.common)
	return c
}

// NewClientFromTokenSource is a wrapper around NewClient if you have a valid
// token source. If no context is available you can use context.Background()
func NewClientFromTokenSource(ctx context.Context, tokenSource oauth2.TokenSource) *Client {
	httpClient := oauth2.NewClient(ctx, tokenSource)
	return NewClient(httpClient)
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
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	return c.BaseURL.Parse(strings.Replace(path, "{division}", strconv.Itoa(division), 1))
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
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
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
	resp, err := c.client.Do(req)
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
	res, err := c.Do(ctx, req, v)
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
	if r.StatusCode == 500 {
		var e InternalServerErrorResponse
		f := json.NewDecoder(r.Body).Decode(e)
		if f != nil {
			return fmt.Errorf("%s: ListRequestAndDo for %s, also encountered an error "+
				"Unmarshalling the error response", r.Status, u)
		}
		return fmt.Errorf("%s: ListRequestAndDo for %s, with message %s", r.Status,
			u, e.Error.Message.Value)
	}

	if r.StatusCode == 400 || r.StatusCode == 401 || r.StatusCode == 403 ||
		r.StatusCode == 404 {
		return fmt.Errorf("%s: ListRequestAndDo for %s", r.Status, u)
	}

	return nil
}
