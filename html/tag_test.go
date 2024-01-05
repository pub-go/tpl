package html

import (
	"reflect"
	"testing"
)

func TestTag_SortedAttr(t *testing.T) {
	type fields struct {
		Name   string
		Attrs  []*Attr
		sorted int
		attrs  map[string]*Attr
	}
	type args struct {
		prefix string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*Attr
	}{
		{name: "empty", fields: fields{Attrs: []*Attr{}}, want: []*Attr{}},
		{name: "one", fields: fields{Attrs: []*Attr{
			{Name: "name"},
		}}, want: []*Attr{
			{Name: "name"},
		}},
		{
			name: "two-prefix",
			fields: fields{Attrs: []*Attr{
				{Name: "name"},
				{Name: ":name"},
			}},
			args: args{prefix: ":"},
			want: []*Attr{
				{Name: ":name"},
				{Name: "name"},
			},
		},
		{
			name: "two-normal",
			fields: fields{Attrs: []*Attr{
				{Name: "xyz"},
				{Name: "abc"},
			}},
			args: args{prefix: ":"},
			want: []*Attr{
				{Name: "xyz"},
				{Name: "abc"},
			},
		},
		{
			name: "two-cmd",
			fields: fields{Attrs: []*Attr{
				{Name: ":range"},
				{Name: ":if"},
			}},
			args: args{prefix: ":"},
			want: []*Attr{
				{Name: ":if"},
				{Name: ":range"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tag{
				Name:   tt.fields.Name,
				Attrs:  tt.fields.Attrs,
				sorted: tt.fields.sorted,
				attrs:  tt.fields.attrs,
			}
			if got := tr.SortedAttr(tt.args.prefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tag.SortedAttr() = %v, want %v", got, tt.want)
			}
		})
	}
}
