package exactonline

import (
	"encoding/json"
	"net/url"
	"regexp"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Date allows for unmarshalling the date objects returned by Exact.
type Date struct {
	time.Time
}

// IsSet returns a boolean if the Date is actually set.
func (d *Date) IsSet() bool {
	return d.UnixNano() != (time.Time{}).UnixNano()
}

// UnmarshalJSON unmarshals the date format returned from the
// Exact Online API.
func (d *Date) UnmarshalJSON(b []byte) error {
	re := regexp.MustCompile(`[0-9]+`)
	s := re.FindString(string(b))
	if s == "" {
		d.Time = time.Time{}
		return nil
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}

	d.Time = time.Unix(0, i*int64(time.Millisecond))
	return nil
}

// MarshalJSON marshals the date to a format expected by the
// Exact Online API.
func (d *Date) MarshalJSON() ([]byte, error) {
	if !d.IsSet() {
		return json.Marshal(nil)
	}

	return json.Marshal(d.Time)
}

// MetaData holds the uri and type of a result object.
type MetaData struct {
	URI  *url.URL `json:"uri"`
	Type *string  `json:"type"`
}

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
	s := b[1 : len(b)-1]
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

// ListOptions holds options for fetching a list.
type ListOptions struct {
	Select  string
	Filter  string
	OrderBy string
	Top     uint
	Skip    string
	Expand  string
	Format  string
}

// ListResponse Holds the list response data.
type ListResponse struct {
	Data struct {
		Results json.RawMessage `json:"results,omitempty"`
		Next    string          `json:"__next,omitempty"`
	} `json:"d,omitempty"`
}

// InternalServerErrorResponse Holds the error 500 response data.
// Most of the time these are validation errors.
type InternalServerErrorResponse struct {
	Error struct {
		Code    string `json:"code,omitempty"`
		Message struct {
			Lang  string `json:"lang,omitempty"`
			Value string `json:"value,omitempty"`
		}
	} `json:"error,omitempty"`
}
