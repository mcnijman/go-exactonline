// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package api

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/mcnijman/go-exactonline/types"
)

/* func TestAddListOptionsToRequest(t *testing.T) {
	opts := NewListOptions()

	opts.Select.Add("test")
	opts.Select.Add("test2")
	opts.Filter.Set("foobar")
	opts.Top.Set(10)

	req, _ := http.NewRequest("GET", "foo", nil)

	err := AddListOptionsToRequest(req, opts)
	if err != nil {
		t.Errorf("AddListOptionsToRequest() error is suposed to be nil, error = %v", err)
	}

	t.Errorf("%+v", req)
} */

func TestSelect_Set(t *testing.T) {
	type fields struct {
		v []string
	}
	type args struct {
		fields []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Select
	}{
		{"1", fields{[]string{"foo", "bar"}}, args{[]string{"foo", "bar"}}, &Select{[]string{"foo", "bar"}}},
		{"2", fields{[]string{"foo"}}, args{[]string{"foo"}}, &Select{[]string{"foo"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Select{
				v: tt.fields.v,
			}
			s.Set(tt.args.fields)
			if !reflect.DeepEqual(s, tt.want) {
				t.Errorf("Select.Set(%v), should result in %v, got %v", tt.args.fields, *tt.want, *s)
			}
		})
	}
}

func TestSelect_Add(t *testing.T) {
	type fields struct {
		v []string
	}
	type args struct {
		field string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Select
	}{
		{"1", fields{[]string{"foo"}}, args{"foo"}, &Select{[]string{"foo"}}},
		{"2", fields{[]string{"foo"}}, args{"bar"}, &Select{[]string{"foo", "bar"}}},
		{"3", fields{}, args{"bar"}, &Select{[]string{"bar"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Select{
				v: tt.fields.v,
			}
			s.Add(tt.args.field)
			if !reflect.DeepEqual(s, tt.want) {
				t.Errorf("Select.Add(%v), should result in %v, got %v", tt.args.field, *tt.want, *s)
			}
		})
	}
}

func TestSelect_Remove(t *testing.T) {
	type fields struct {
		v []string
	}
	type args struct {
		field string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Select
	}{
		{"1", fields{[]string{"foo", "bar"}}, args{"foo"}, &Select{[]string{"bar"}}},
		{"2", fields{[]string{"foo", "bar"}}, args{"bar"}, &Select{[]string{"foo"}}},
		{"3", fields{[]string{"foo"}}, args{"bar"}, &Select{[]string{"foo"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Select{
				v: tt.fields.v,
			}
			s.Remove(tt.args.field)
			if !reflect.DeepEqual(s, tt.want) {
				t.Errorf("Select.Remove(%v), should result in %v, got %v", tt.args.field, *tt.want, *s)
			}
		})
	}
}

func TestSelect_MarshalSchema(t *testing.T) {
	type fields struct {
		v []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"1", fields{[]string{"foo", "bar"}}, "foo,bar"},
		{"2", fields{[]string{"bar", "foo"}}, "bar,foo"},
		{"3", fields{[]string{"foo", "bar", "test"}}, "foo,bar,test"},
		{"4", fields{[]string{"foo"}}, "foo"},
		{"5", fields{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Select{
				v: tt.fields.v,
			}
			if got := s.MarshalSchema(); got != tt.want {
				t.Errorf("Select.MarshalSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter_Set(t *testing.T) {
	type fields struct {
		v string
	}
	type args struct {
		v string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Filter
	}{
		{"1", fields{""}, args{"foo"}, &Filter{"foo"}},
		{"2", fields{"foo"}, args{"bar"}, &Filter{"bar"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Filter{
				v: tt.fields.v,
			}
			f.Set(tt.args.v)
			if !reflect.DeepEqual(f, tt.want) {
				t.Errorf("Filter.Set(%v), should result in %v, got %v", tt.args.v, *tt.want, *f)
			}
		})
	}
}

func TestFilter_MarshalSchema(t *testing.T) {
	type fields struct {
		v string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"1", fields{"foo"}, "foo"},
		{"2", fields{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Filter{
				v: tt.fields.v,
			}
			if got := f.MarshalSchema(); got != tt.want {
				t.Errorf("Filter.MarshalSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderBy_MarshalSchema(t *testing.T) {
	type fields struct {
		v map[string]bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"1", fields{map[string]bool{"foo": true, "bar": false}}, "bar desc,foo asc"},
		{"2", fields{map[string]bool{"bar": true, "foo": false}}, "bar asc,foo desc"},
		{"3", fields{map[string]bool{"foo": true, "bar": false, "test": true}}, "bar desc,foo asc,test asc"},
		{"4", fields{map[string]bool{"foo": true}}, "foo asc"},
		{"5", fields{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OrderBy{
				v: tt.fields.v,
			}
			if got := o.MarshalSchema(); got != tt.want {
				t.Errorf("OrderBy.MarshalSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderBy_Add(t *testing.T) {
	type fields struct {
		v map[string]bool
	}
	type args struct {
		field     string
		ascending bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderBy
	}{
		{"1", fields{map[string]bool{"foo": true}}, args{"bar", true}, &OrderBy{map[string]bool{"foo": true, "bar": true}}},
		{"2", fields{map[string]bool{"foo": true}}, args{"bar", false}, &OrderBy{map[string]bool{"foo": true, "bar": false}}},
		{"3", fields{}, args{"bar", true}, &OrderBy{map[string]bool{"bar": true}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OrderBy{
				v: tt.fields.v,
			}
			o.Add(tt.args.field, tt.args.ascending)
			if !reflect.DeepEqual(o, tt.want) {
				t.Errorf("OrderBy.Add(%v), should result in %v, got %v", tt.args, *tt.want, *o)
			}
		})
	}
}

func TestOrderBy_Remove(t *testing.T) {
	type fields struct {
		v map[string]bool
	}
	type args struct {
		field string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderBy
	}{
		{"1", fields{map[string]bool{"foo": true}}, args{"foo"}, &OrderBy{map[string]bool{}}},
		{"2", fields{map[string]bool{"foo": true}}, args{"bar"}, &OrderBy{map[string]bool{"foo": true}}},
		{"3", fields{}, args{"bar"}, &OrderBy{map[string]bool{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OrderBy{
				v: tt.fields.v,
			}
			o.Remove(tt.args.field)
			if !reflect.DeepEqual(o, tt.want) {
				t.Errorf("OrderBy.Remove(%v), should result in %v, got %v", tt.args, *tt.want, *o)
			}
		})
	}
}

func TestSkip_Set(t *testing.T) {
	type fields struct {
		v uint
	}
	type args struct {
		v uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Skip
	}{
		{"1", fields{}, args{uint(10)}, &Skip{uint(10)}},
		{"2", fields{uint(11)}, args{uint(20)}, &Skip{uint(20)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Skip{
				v: tt.fields.v,
			}
			s.Set(tt.args.v)
			if !reflect.DeepEqual(s, tt.want) {
				t.Errorf("Skip.Set(%v), should result in %v, got %v", tt.args.v, *tt.want, *s)
			}
		})
	}
}

func TestSkip_MarshalSchema(t *testing.T) {
	type fields struct {
		v uint
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"1", fields{uint(20)}, "20"},
		{"2", fields{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Skip{
				v: tt.fields.v,
			}
			if got := s.MarshalSchema(); got != tt.want {
				t.Errorf("Skip.MarshalSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSkipToken_Set(t *testing.T) {
	type fields struct {
		v types.GUID
	}
	type args struct {
		v types.GUID
	}
	g1 := types.NewGUID()
	g2 := types.NewGUID()
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SkipToken
	}{
		{"1", fields{}, args{g1}, &SkipToken{g1}},
		{"2", fields{g1}, args{g2}, &SkipToken{g2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SkipToken{
				v: tt.fields.v,
			}
			s.Set(tt.args.v)
			if !reflect.DeepEqual(s, tt.want) {
				t.Errorf("SkipToken.Set(%v), should result in %v, got %v", tt.args.v, *tt.want, *s)
			}
		})
	}
}

func TestSkipToken_MarshalSchema(t *testing.T) {
	type fields struct {
		v types.GUID
	}
	g1 := types.NewGUID()
	g2 := types.NewGUID()
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"1", fields{g1}, g1.String()},
		{"2", fields{g2}, g2.String()},
		{"3", fields{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SkipToken{
				v: tt.fields.v,
			}
			if got := s.MarshalSchema(); got != tt.want {
				t.Errorf("SkipToken.MarshalSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTop_Set(t *testing.T) {
	type fields struct {
		v uint
	}
	type args struct {
		v uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Top
	}{
		{"1", fields{}, args{uint(10)}, &Top{uint(10)}},
		{"2", fields{uint(11)}, args{uint(20)}, &Top{uint(20)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			to := &Top{
				v: tt.fields.v,
			}
			to.Set(tt.args.v)
			if !reflect.DeepEqual(to, tt.want) {
				t.Errorf("Top.Set(%v), should result in %v, got %v", tt.args.v, *tt.want, *to)
			}
		})
	}
}

func TestTop_MarshalSchema(t *testing.T) {
	type fields struct {
		v uint
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"1", fields{uint(20)}, "20"},
		{"2", fields{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			to := &Top{
				v: tt.fields.v,
			}
			if got := to.MarshalSchema(); got != tt.want {
				t.Errorf("Top.MarshalSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewListOptions(t *testing.T) {
	tests := []struct {
		name string
		want *ListOptions
	}{
		{"1", &ListOptions{
			&Filter{},
			&OrderBy{},
			&Select{},
			&Skip{},
			&SkipToken{},
			&Top{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListOptions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tagOptions_Contains(t *testing.T) {
	type args struct {
		option string
	}
	tests := []struct {
		name string
		o    tagOptions
		args args
		want bool
	}{
		{"1", tagOptions{"1", "2"}, args{"2"}, true},
		{"2", tagOptions{"1", "2"}, args{"3"}, false},
		{"3", tagOptions{}, args{"2"}, false},
		{"4", tagOptions{}, args{""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Contains(tt.args.option); got != tt.want {
				t.Errorf("tagOptions.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseTag(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 tagOptions
	}{
		{"1", args{"name,omitempty"}, "name", tagOptions{"omitempty"}},
		{"1", args{"name"}, "name", tagOptions{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseTag(tt.args.tag)
			if got != tt.want {
				t.Errorf("parseTag() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseTag() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAddListOptionsToRequest(t *testing.T) {
	type args struct {
		r *http.Request
		o *ListOptions
	}
	r1, _ := http.NewRequest("GET", "/foo", nil)
	r2, _ := http.NewRequest("GET", "/foo", nil)
	r3, _ := http.NewRequest("GET", "/foo", nil)
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{
			r1,
			&ListOptions{},
		}, ""},
		{"2", args{
			r2,
			&ListOptions{Select: &Select{[]string{"foo", "bar"}}},
		}, "$select=foo%2Cbar"},
		{"3", args{
			r3,
			&ListOptions{Select: &Select{[]string{"foo", "bar"}}, Top: &Top{uint(10)}},
		}, "$select=foo%2Cbar&$top=10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddListOptionsToRequest(tt.args.r, tt.args.o)

			if got := tt.args.r.URL.RawQuery; got != tt.want {
				t.Errorf("AddListOptionsToRequest, should generate query: %v, got %v", tt.want, got)
			}
		})
	}
}

func TestAddListOptionsToURL(t *testing.T) {
	type args struct {
		u *url.URL
		o *ListOptions
	}
	u1, _ := url.Parse("https://start.exactonline.nl/foobar")
	u2, _ := url.Parse("https://start.exactonline.nl/foobar")
	u3, _ := url.Parse("https://start.exactonline.nl/foobar")
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{
			u1,
			&ListOptions{},
		}, ""},
		{"2", args{
			u2,
			&ListOptions{Select: &Select{[]string{"foo", "bar"}}},
		}, "$select=foo%2Cbar"},
		{"3", args{
			u3,
			&ListOptions{Select: &Select{[]string{"foo", "bar"}}, Top: &Top{uint(10)}},
		}, "$select=foo%2Cbar&$top=10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddListOptionsToURL(tt.args.u, tt.args.o)
		})

		if got := tt.args.u.RawQuery; got != tt.want {
			t.Errorf("AddListOptionsToRequest, should generate query: %v, got %v", tt.want, got)
		}
	}
}
