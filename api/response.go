// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/sebas7dk/go-exactonline/types"
)

type data struct {
	Data json.RawMessage `json:"d"`
}

type listData struct {
	Results json.RawMessage `json:"results"`
	Next    string          `json:"__next"`
}

// Response is a Exact Online API response. This wraps the standard http.Response
// returned from Exact Online and provides convenient access to things like
// pagination links.
type Response struct {
	*http.Response

	Data     json.RawMessage
	NextPage *url.URL
}

// UnmarshalJSON unmarshals the JSON response
func (r *Response) UnmarshalJSON(j []byte) error {
	var d data
	err := json.Unmarshal(j, &d)
	if err == io.EOF {
		return nil // ignore EOF errors caused by empty response body
	}
	if err != nil {
		return fmt.Errorf("Response.UnmarshalJSON() error: %v", err)
	}
	j = d.Data

	if strings.HasPrefix(string(j), "[") {
		j = append([]byte(`{"results":`), j...)
		j = append(j, []byte("}")...)
	}
	if !strings.Contains(string(j), `"results":`) &&
		!strings.Contains(string(j), `"__next":`) {
		r.Data = j
		return nil
	}

	var l listData
	err = json.Unmarshal(j, &l)
	if err == io.EOF {
		return nil // ignore EOF errors caused by empty response body
	}
	if err != nil {
		return fmt.Errorf("Response.UnmarshalJSON() error: %v", err)
	}

	r.Data = l.Results
	if l.Next != "" {
		r.NextPage, err = url.Parse(l.Next)
	}

	if err != nil {
		return fmt.Errorf("Response.UnmarshalJSON() parse url error: %v", err)
	}

	return nil
}

// ListResponse Holds the list response data.
type ListResponse struct {
	Data struct {
		Results json.RawMessage `json:"results,omitempty"`
		Next    string          `json:"__next,omitempty"`
	} `json:"d,omitempty"`
}

// ErrorResponse Holds the json error response data.
// Most of the time these are validation errors.
type ErrorResponse struct {
	*http.Response
	Err struct {
		Code    string `json:"code,omitempty"`
		Message struct {
			Lang  string `json:"lang,omitempty"`
			Value string `json:"value,omitempty"`
		}
	} `json:"error,omitempty"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("ErrorResponse %s: %s", e.Err.Code, e.Err.Message)
}

// MetaData holds the uri and type of a result object.
type MetaData struct {
	URI  *types.URL `json:"uri,omitempty"`
	Type *string    `json:"type,omitempty"`
}
