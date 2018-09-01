// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package types

import (
	"encoding/json"
	"net/url"
)

// URL allows for unmarshalling the urls returned by Exact.
type URL struct {
	*url.URL
}

// UnmarshalJSON unmarshals the url to url.URL returned from the
// Exact Online API.
func (u *URL) UnmarshalJSON(b []byte) error {
	s := string(b[1 : len(b)-1]) // drop quotes
	j, err := url.Parse(s)
	if err == nil {
		u.URL = j
	}
	return err
}

// MarshalJSON marshals the url to a format expected by the
// Exact Online API.
func (u *URL) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}
