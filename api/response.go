// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// MetaData holds the uri and type of a result object.
type MetaData struct {
	URI  *url.URL `json:"uri"`
	Type *string  `json:"type"`
}

// ListResponse Holds the list response data.
type ListResponse struct {
	Data struct {
		Results json.RawMessage `json:"results,omitempty"`
		Next    string          `json:"__next,omitempty"`
	} `json:"d,omitempty"`
}

// ListResponseSliced Holds the list response data.
type ListResponseSliced struct {
	Data struct {
		Results []json.RawMessage `json:"results,omitempty"`
		Next    string            `json:"__next,omitempty"`
	} `json:"d,omitempty"`
}

// ErrorResponse Holds the json error response data.
// Most of the time these are validation errors.
type ErrorResponse struct {
	Response *http.Response
	Err      struct {
		Code    string `json:"code,omitempty"`
		Message struct {
			Lang  string `json:"lang,omitempty"`
			Value string `json:"value,omitempty"`
		}
	} `json:"error,omitempty"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("Error %s: %s", e.Err.Code, e.Err.Message)
}
