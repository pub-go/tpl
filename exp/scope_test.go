package exp

import (
	"reflect"
	"testing"
	"time"
)

type Err struct{}

var _ error = (*Err)(nil)

func (Err) Error() string { return "err" }
func retErr() *Err        { return nil }

func TestWithDefaultScope(t *testing.T) {
	s := WithDefaultScope(NewScope(map[string]any{
		"arr": make([]int, 2, 4),
		"err": retErr(),
	}))
	tests := []struct {
		input string
		want  any
	}{
		{input: `float64(1)`, want: float64(1)},
		{input: `duration(1)`, want: time.Duration(1)},
		{input: `isNil(err)`, want: false},
		{input: `isNull(err)`, want: true},
		{input: `len(arr)`, want: 2},
		{input: `cap(arr)`, want: 4},
		{input: `print(err)`, want: "<nil>"},
		{input: `printf("%v", err)`, want: "<nil>"},
		{input: `println(err)`, want: "<nil>\n"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			tree, err := ParseCode(tt.input)
			if err != nil {
				t.Errorf("parse code err=%+v", err)
			}
			result, err := Evaluate(NewPos(1, 1), tree, s)
			if err != nil {
				t.Errorf("Evaluate err=%+v", err)
			}
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("Evaluate got=%+v, want=%+v", result, tt.want)
			}
		})
	}

}
