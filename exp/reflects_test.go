package exp

import (
	"reflect"
	"testing"
)

type Foo struct {
	Name     string
	unexport string
}

func (Foo) Foo() string  { return "Foo" }
func (*Foo) Bar() string { return "Bar" }

var foo = Foo{Name: "Name"}

func TestGetMethod(t *testing.T) {
	tests := []struct {
		name    string
		field   string
		from    any
		wantErr bool
		isFn    bool
		result  string
	}{
		{name: "foo.Name", field: "Name", from: foo, isFn: false, result: "Name"},
		{name: "&foo.Name", field: "Name", from: foo, isFn: false, result: "Name"},
		{name: "foo.Foo", field: "Foo", from: foo, isFn: true, result: "Foo"},
		{name: "&foo.Foo", field: "Foo", from: &foo, isFn: true, result: "Foo"},
		{name: "foo.Bar", field: "Bar", from: foo, wantErr: true},
		{name: "&foo.Bar", field: "Bar", from: &foo, isFn: true, result: "Bar"},
		{name: "&foo.No", field: "No", from: &foo, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := getValue(tt.field, tt.from)
			if err != nil {
				if tt.wantErr {
					t.Logf("%v", err)
					return
				}
				t.Fatalf("err=%v", err)
			}
			f, ok := v.(func() string)
			if !ok && tt.isFn {
				t.Fatalf("getMethodFail")
			}
			var got string
			if !ok {
				got = v.(string)
			}
			if tt.isFn {
				got = f()
			}
			if got != tt.result {
				t.Fatalf("got=%v, want=%v", got, tt.result)

			}
		})
	}
}

func Test_getValue(t *testing.T) {
	type args struct {
		name string
		from any
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{name: "struct object", args: args{
			name: "Name",
			from: Foo{Name: "张三"},
		}, want: "张三", wantErr: false},
		{name: "struct-no-field", args: args{
			name: "Bar",
			from: Foo{Name: "张三"},
		}, want: nil, wantErr: true},
		{name: "struct private", args: args{
			name: "unexport",
			from: Foo{Name: "张三", unexport: "李四"},
		}, want: nil, wantErr: true},
		{name: "struct pointer", args: args{
			name: "Name",
			from: &Foo{Name: "张三"},
		}, want: "张三", wantErr: false},
		{name: "map", args: args{
			name: "Key",
			from: map[string]string{"Key": "张三"},
		}, want: "张三", wantErr: false},
		{name: "map-no-such-key", args: args{
			name: "O",
			from: map[string]string{"Key": "张三"},
		}, want: nil, wantErr: true},
		{name: "slice-nil-element", args: args{
			name: "1",
			from: []any{"Key", nil, "张三"},
		}, want: nil, wantErr: false},
		{name: "slice", args: args{
			name: "0",
			from: []string{"Key", "张三"},
		}, want: "Key", wantErr: false},
		{name: "arr-not-idx", args: args{
			name: "a",
			from: [...]string{"Key", "张三"},
		}, want: nil, wantErr: true},
		{name: "arr-neg-idx", args: args{
			name: "-1",
			from: [...]string{"Key", "张三"},
		}, want: "张三", wantErr: false},
		{name: "other-nil", args: args{
			name: "Key",
			from: nil,
		}, want: nil, wantErr: true},
		{name: "other", args: args{
			name: "Key",
			from: 123,
		}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getValue(tt.args.name, tt.args.from)
			if err != nil {
				t.Logf("err=%v", err)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("getValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
