// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package api

/*
// ListQueryOptions holds options for fetching a list.
type ListQueryOptions struct {
	// Select allows for limiting the fields returned, * is all fields
	Select *string

	// Filter will result in a subset filtered by the paramets
	Filter *string

	// OrderBy specifies an expression for determining what values are used
	// to order the collection
	OrderBy *string

	// Top reduces the count of results results returned
	Top *uint

	// SkipToken is a token returned from the API to returned the next paginated page
	SkipToken *types.GUID

	// Expand eagerly loads embedded entities
	Expand *string

	// Format string
	// InlineCount string
}

// GetQueryString constructs a querystring from provided parameters
func (o ListQueryOptions) GetQueryString() string {
	s := "*"
	if o.Select != nil {
		s = *o.Select
	}
	return "?$select=" + s
}

// OptionsFromQuery constructs ListQueryOptions based on a querystring in an URL
func OptionsFromQuery(u *url.URL) ListQueryOptions {
	return ListQueryOptions{}
}
*/
