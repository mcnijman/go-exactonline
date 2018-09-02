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

func TestGUID_UnMarshalJSON(t *testing.T) {
	g := NewGUID()
	b := []byte(`"` + g.UUID.String() + `"`)

	var u GUID
	json.Unmarshal(b, &u)

	if !reflect.DeepEqual(u, g) {
		t.Errorf("Unmarshalled GUID failed: got %s, want %s", u.UUID.String(), g.UUID.String())
	}

}

func TestGUID_MarshalJSON(t *testing.T) {
	g := NewGUID()
	got, err := json.Marshal(g)

	if err != nil {
		t.Errorf("URL.MarshalJSON() error = %v", err)
	}

	want := []byte(`"` + g.UUID.String() + `"`)

	if string(got) != string(want) {
		t.Errorf("Marshalling GUID failed: got %v, want %v", got, want)
	}

	b, err := g.MarshalJSON()
	if err != nil {
		t.Errorf("Unable to marshal JSON for %v", g)
	}
	if string(b) != string(want) {
		t.Errorf("String is not valid json: got %s, want %s", string(b), string(want))
	}
}

func TestGUID_String(t *testing.T) {
	type fields struct {
		UUID uuid.UUID
	}
	g := NewGUID()
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"t", fields{g.UUID}, g.UUID.String()},
		{"t2", fields{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GUID{
				UUID: tt.fields.UUID,
			}
			if got := g.String(); got != tt.want {
				t.Errorf("GUID.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
