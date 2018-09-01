// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package types

import (
	"net/url"
	"reflect"
	"testing"
)

func TestURL_UnmarshalJSON(t *testing.T) {
	type fields struct {
		URL *url.URL
	}
	type args struct {
		b []byte
	}

	u, _ := url.Parse("https://start.exactonline.nl/")
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		wantEqual bool
	}{
		{"URI not quoted", fields{u}, args{[]byte(`https://start.exactonline.nl/`)}, false, true},
		{"URI quoted", fields{u}, args{[]byte(`"https://start.exactonline.nl/"`)}, false, true},
		{"URI quoted false", fields{u}, args{[]byte(`"https://start.exactonlinee.nl/"`)}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u2 := &URL{}
			if err := u2.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("URL.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if (reflect.DeepEqual(u2.URL, tt.fields.URL)) != tt.wantEqual {
				t.Errorf("URL.UnmarshalJSON() got = %+v, want %+v, equal %+v", u2, tt.fields.URL, reflect.DeepEqual(u2, tt.fields.URL))
			}
		})
	}
}

func TestURL_MarshalJSON(t *testing.T) {
	type fields struct {
		URL *url.URL
	}
	u, _ := url.Parse("https://start.exactonline.nl/")
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{"URI not quoted", fields{u}, []byte(`"https://start.exactonline.nl/"`), false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &URL{
				URL: tt.fields.URL,
			}
			got, err := u.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("URL.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URL.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
