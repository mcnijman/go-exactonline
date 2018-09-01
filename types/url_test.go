// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package types

import (
	"encoding/json"
	"testing"
)

func TestURL_unMarshalJSON(t *testing.T) {
	uri := "https://start.exactonline.nl/"
	b := []byte(`"` + uri + `"`)

	var u URL
	json.Unmarshal(b, &u)

	if u.URL.String() != uri {
		t.Errorf("Unmarshalled URI failed: got %s, want %s", u.URL.String(), uri)
	}
}

func TestURL_MarshalJSON(t *testing.T) {

}
