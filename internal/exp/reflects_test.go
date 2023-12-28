package exp

import (
	"reflect"
	"testing"
)

type Foo struct {
	Name     string
	unexport string
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
