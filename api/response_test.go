package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestResponse_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Response *http.Response
		Data     json.RawMessage
		NextPage *url.URL
	}
	type args struct {
		j []byte
	}
	tests := []struct {
		name    string
		args    args
		want    fields
		wantErr bool
	}{
		{
			"1",
			args{[]byte(`{ "d": { "foo": "bar" } }`)},
			fields{nil, []byte(`{ "foo": "bar" }`), nil},
			false,
		},
		{
			"2",
			args{[]byte(`{ "d": [ { "foo": "bar" } ] }`)},
			fields{nil, []byte(`[ { "foo": "bar" } ]`), nil},
			false,
		},
		{
			"3",
			args{[]byte(`{ "d": { "results": [ { "foo": "bar" } ] }}`)},
			fields{nil, []byte(`[ { "foo": "bar" } ]`), nil},
			false,
		},
		{
			"4",
			args{[]byte(`{ "d": { "results": [ { "foo": "bar" } ], "__next": "https://www.google.com"} }`)},
			fields{nil, []byte(`[ { "foo": "bar" } ]`), &url.URL{
				Scheme: "https",
				Host:   "www.google.com",
			}},
			false,
		},
		{
			"5",
			args{[]byte(`{ "d": { "__next": "https://www.google.com", "results": [ { "foo": "bar" } ] } }`)},
			fields{nil, []byte(`[ { "foo": "bar" } ]`), &url.URL{
				Scheme: "https",
				Host:   "www.google.com",
			}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{}
			if err := r.UnmarshalJSON(tt.args.j); (err != nil) != tt.wantErr {
				t.Errorf("Response.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && !reflect.DeepEqual(r.Data, tt.want.Data) {
				t.Errorf("Response.UnmarshalJSON() Data not equal, got: %v, want %v", string(r.Data), string(tt.want.Data))
			}

			if !tt.wantErr && !reflect.DeepEqual(r.NextPage, tt.want.NextPage) {
				t.Errorf("Response.UnmarshalJSON() NextPage not equal, got: %v, want %v", *r.NextPage, *tt.want.NextPage)
			}
		})
	}
}
