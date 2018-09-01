// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package types

import (
	"encoding/json"
	"strings"

	"github.com/gofrs/uuid"
)

// NewGUID generates a new GUID.
func NewGUID() GUID {
	return GUID{UUID: uuid.Must(uuid.NewV4())}
}

// GUID allows for unmarshalling the urls returned by Exact.
type GUID struct {
	uuid.UUID
}

// UnmarshalJSON unmarshals the guid to uuid.UUID returned from the
// Exact Online API.
func (g *GUID) UnmarshalJSON(b []byte) error {
	s := []byte(strings.Replace(string(b), `"`, "", -1))
	err := (&g.UUID).UnmarshalText(s)
	return err
}

// MarshalJSON marshals the url to a format expected by the
// Exact Online API.
func (g *GUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.String())
}

func (g *GUID) String() string {
	if g.UUID == uuid.Nil {
		return ""
	}
	return g.UUID.String()
}
