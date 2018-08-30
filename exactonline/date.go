package exactonline

import (
	"strconv"
	"strings"
	"time"
)

// Date allows for unmarshalling the date objects returned by Exact
type Date struct {
	time.Time
}

// IsSet returns a boolean if the Date is actually set
func (d *Date) IsSet() bool {
	return d.UnixNano() != (time.Time{}).UnixNano()
}

// UnmarshalJSON unmarshals the date format returned from the
// Exact Online API
func (d *Date) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	s := strings.TrimPrefix(strings.TrimSuffix(string(b), `)\/`), `\/Date(`)

	if s == "null" {
		d.Time = time.Time{}
		return
	}

	s = s[:10]

	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		d.Time = time.Unix(i, 0)
	}

	return
}

// MarshalJSON marshals the date to a format expected by the
// Exact Online API
func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(d.Time.Format(`"2006-01-02"`)), nil
}
