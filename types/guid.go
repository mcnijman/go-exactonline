// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package types

import (
	"encoding/json"
	"fmt"
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
	if err != nil {
		return fmt.Errorf("GUID.UnmarshalJSON() error: %v", err)
	}
	return nil
}

// MarshalJSON marshals the url to a format expected by the
// Exact Online API.
func (g *GUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.String())
}

func (g *GUID) String() string {
	if !g.IsSet() {
		return ""
	}
	return g.UUID.String()
}

// IsSet checks if the GUID/uuid actually exists
func (g *GUID) IsSet() bool {
	return g.UUID != uuid.Nil
}
