package exactonline

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/gofrs/uuid"
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

}

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
