package html

import "testing"

func Test_extractRange(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{args: args{s: `names `}, want: []string{"", "", "names"}},
		{args: args{s: `nam es`}, want: []string{"", "", "nam es"}},
		{args: args{s: `idx:names`}, want: []string{"idx", "", "names"}},
		{args: args{s: `idx : names`}, want: []string{"idx", "", "names"}},
		{args: args{s: `_, item: names`}, want: []string{"_", "item", "names"}},
		{args: args{s: `, item: names`}, want: []string{"", "item", "names"}},
		{args: args{s: ` , item : names`}, want: []string{"", "item", "names"}},
		{args: args{s: `  , : names`}, want: []string{"", "", "names"}},
		{args: args{s: ` i , : names`}, want: []string{"i", "", "names"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIdxName, gotItemName, gotObjName, err := extractRange(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIdxName != tt.want[0] {
				t.Errorf("extractRange() gotIdxName = %v, want %v", gotIdxName, tt.want[0])
			}
			if gotItemName != tt.want[1] {
				t.Errorf("extractRange() gotItemName = %v, want %v", gotItemName, tt.want[1])
			}
			if gotObjName != tt.want[2] {
				t.Errorf("extractRange() gotObjName = %v, want %v", gotObjName, tt.want[2])
			}
		})
	}
}
