package api

import "fmt"

// Error holds the error message
type Error struct {
	Service  string `json:"service"`
	Endpoint string `json:"endpoint"`
	Field    string `json:"field"`
	Code     string `json:"code"`
	Message  string `json:"Message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v error caused by %v field in %v endpoint of %v service",
		e.Code, e.Field, e.Endpoint, e.Service)
}
