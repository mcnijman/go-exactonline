// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package types

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/gofrs/uuid"
)

func TestNewGUID(t *testing.T) {
	g := NewGUID()
	if reflect.TypeOf(g) != reflect.TypeOf(GUID{}) {
		t.Errorf("GUID has an incorrect type: got %s, want GUID", reflect.TypeOf(g))
	}

	if g.UUID.Version() != uuid.V4 {
		t.Errorf("GUID has an incorrect incorrect uuid version: got %s, want V4", string(g.UUID.Version()))
	}
}

func TestGUID_unMarshalJSON(t *testing.T) {
	g := NewGUID()
	b := []byte(`"` + g.UUID.String() + `"`)

	var u GUID
	json.Unmarshal(b, &u)

	if !reflect.DeepEqual(u, g) {
		t.Errorf("Unmarshalled URI failed: got %s, want %s", u.UUID.String(), g.UUID.String())
	}

}

func TestGUID_MarshalJSON(t *testing.T) {

}
