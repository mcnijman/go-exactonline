package exactonline

import "encoding/json"

// MetaData holds the uri and type of a result object
type MetaData struct {
	URI  string `json:"uri"`
	Type string `json:"type"`
}

// ListOptions holds options for fetching a list
type ListOptions struct {
	Select  string
	Filter  string
	OrderBy string
	Top     uint
	Skip    string
	Expand  string
	Format  string
}

// ListResponse Holds the list response data
type ListResponse struct {
	Data *struct {
		Results json.RawMessage `json:"results"`
		Next    string          `json:"__next"`
	} `json:"d"`
}
