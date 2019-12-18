// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package api

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/sebas7dk/go-exactonline/types"
)

// Select allows for limiting the fields returned. The default is all fields (*).
type Select struct {
	v []string
}

// Set sets the select query options
func (s *Select) Set(fields []string) {
	s.v = fields
}

// Add a field to the Select query
func (s *Select) Add(field string) {
	for _, key := range s.v {
		if key == field {
			return
		}
	}
	s.v = append(s.v, field)
}

// Remove a field from the Select query
func (s *Select) Remove(field string) {
	for i, v := range s.v {
		if v == field {
			s.v = append(s.v[:i], s.v[i+1:]...)
			break
		}
	}
}

// MarshalSchema marshals the options in a query string
func (s *Select) MarshalSchema() string {
	if len(s.v) == 0 {
		return "*" // by default return all fields
	}
	return strings.Join(s.v, ",")
}

// Filter will result in a subset filtered by the parameters.
// Filters aren't really implemented by this library and should be constructed
// by the user and past as a string.
type Filter struct {
	v string
}

// Set sets the filter query string
func (f *Filter) Set(v string) {
	f.v = v
}

// MarshalSchema marshals the options in a query string
func (f *Filter) MarshalSchema() string {
	return f.v
}

// OrderBy specifies an expression for determining what values are used
// to order the collection
type OrderBy struct {
	v map[string]bool
}

// MarshalSchema marshals the options in a query string
func (o *OrderBy) MarshalSchema() string {
	keys := make([]string, 0, len(o.v))
	opts := make([]string, 0, len(o.v))

	for k := range o.v {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v, _ := o.v[k]
		order := "asc"
		if !v {
			order = "desc"
		}
		opt := fmt.Sprintf("%s %s", k, order)
		opts = append(opts, opt)
	}
	return strings.Join(opts, ",")
}

// Add a field to the OrderBy set
func (o *OrderBy) Add(field string, ascending bool) {
	if o.v == nil {
		o.v = make(map[string]bool)
	}
	o.v[field] = ascending
}

// Remove a field to the OrderBy set
func (o *OrderBy) Remove(field string) {
	if o.v == nil {
		o.v = make(map[string]bool)
	}
	delete(o.v, field)
}

// Skip skips the count of results returned
type Skip struct {
	v uint
}

// Set sets the Skip query params
func (s *Skip) Set(v uint) {
	s.v = v
}

// MarshalSchema marshals the options in a query string
func (s *Skip) MarshalSchema() string {
	if s.v == 0 {
		return ""
	}
	return strconv.FormatUint(uint64(s.v), 10)
}

// SkipToken is a token returned from the API to returned the next paginated page
type SkipToken struct {
	v types.GUID
}

// Set sets the Skiptoken
func (s *SkipToken) Set(v types.GUID) {
	s.v = v
}

// MarshalSchema marshals the options in a query string
func (s *SkipToken) MarshalSchema() string {
	return s.v.String()
}

// Top reduces the count of results returned
type Top struct {
	v uint
}

// Set sets the Top query params
func (t *Top) Set(v uint) {
	t.v = v
}

// MarshalSchema marshals the options in a query string
func (t *Top) MarshalSchema() string {
	if t.v == 0 {
		return ""
	}
	return strconv.FormatUint(uint64(t.v), 10)
}

// ListOptions holds options for fetching a list.
type ListOptions struct {
	Filter    *Filter    `param:"$filter,omitempty"`
	OrderBy   *OrderBy   `param:"$orderby,omitempty"`
	Select    *Select    `param:"$select,omitempty"`
	Skip      *Skip      `param:"$skip,omitempty"`
	SkipToken *SkipToken `param:"$skiptoken,omitempty"`
	Top       *Top       `param:"$top,omitempty"`
}

// NewListOptions returns a newly initialized ListOptions
func NewListOptions() *ListOptions {
	return &ListOptions{
		&Filter{},
		&OrderBy{},
		&Select{},
		&Skip{},
		&SkipToken{},
		&Top{},
	}
}

// schemaMarshaler allows reflection to Marshal the schema
type schemaMarshaler interface {
	MarshalSchema() string
}

// tagOptions is the string following a comma in a struct field's tag, or
// the empty string. It does not include the leading comma.
type tagOptions []string

// Contains checks whether the tagOptions contains the specified option.
func (o tagOptions) Contains(option string) bool {
	for _, s := range o {
		if s == option {
			return true
		}
	}
	return false
}

// parseTag splits a struct field's url tag into its name and comma-separated
// options.
func parseTag(tag string) (string, tagOptions) {
	s := strings.Split(tag, ",")
	return s[0], s[1:]
}

// AddListOptionsToRequest adds the ListOptions to the request
func AddListOptionsToRequest(r *http.Request, o *ListOptions) {
	if o == nil {
		return
	}
	AddListOptionsToURL(r.URL, o)
}

// AddListOptionsToURL adds the ListOptions to the url
func AddListOptionsToURL(u *url.URL, o *ListOptions) {
	if o == nil {
		return
	}
	q := u.Query()
	v := reflect.ValueOf(o).Elem()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		key, opts := parseTag(v.Type().Field(i).Tag.Get("param"))
		if f.IsNil() && opts.Contains("omitempty") {
			continue
		} else if f.IsNil() {
			q.Add(key, "")
		}
		if m, ok := f.Interface().(schemaMarshaler); ok {
			if value := m.MarshalSchema(); value != "" || !opts.Contains("omitempty") {
				q.Add(key, value)
			}
		}
	}
	u.RawQuery = strings.Replace(q.Encode(), `%24`, "$", -1)
}
