// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package types

import (
	"encoding/json"
	"testing"
	"time"
)

func TestDate_isSet(t *testing.T) {
	d := &Date{}

	if d.IsSet() {
		t.Error("Empty Date should return false for Date.IsSet")
	}
}

func TestDate_unMarshalJSON(t *testing.T) {
	b := []byte(`"/Date(1535718344017)/"`)

	var d Date
	json.Unmarshal(b, &d)

	if !d.IsSet() {
		t.Error("Unmarshalled Date should return true for Date.IsSet")
	}

	want := time.Date(2018, 8, 31, 12, 25, 44, 17000000, time.UTC)

	if !d.UTC().Equal(want) {
		t.Errorf("Date = %v, want %v", d, want)
	}
}

func TestDate_unMarshalJSONEmpty(t *testing.T) {
	tests := [][]byte{
		[]byte(`null`),
		[]byte(``),
		[]byte(`""`),
		[]byte(`/Date()/`),
	}

	for _, test := range tests {
		var d Date
		json.Unmarshal(test, &d)

		if d.IsSet() {
			t.Errorf("Unmarshalled Date with value: %s should return false for Date.IsSet", string(test))
		}
	}
}

func TestDate_MarshalJSON(t *testing.T) {
	d := time.Date(2018, 8, 31, 12, 25, 44, 17000000, time.UTC)

	b, err := json.Marshal(d)
	if err != nil {
		t.Errorf("Failed marshalling date: %v", d)
	}

	want := `"2018-08-31T12:25:44.017Z"`
	got := string(b)

	if got != want {
		t.Errorf("Failed marshalling date: got: %v, want %v", got, want)
	}
}
