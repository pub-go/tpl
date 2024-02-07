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

type S string

func TestWithDefaultScope(t *testing.T) {
	s := WithDefaultScope(NewScope(map[string]any{
		"arr":    make([]int, 2, 4),
		"err":    retErr(),
		"second": time.Second,
		"s":      S("sss"),
	}))
	tests := []struct {
		input string
		want  any
	}{
		{input: `false`, want: false},
		{input: `s`, want: S("sss")},
		{input: `string(s)`, want: ("sss")},
		{input: `bytes(s)`, want: []byte("sss")},
		{input: `runes(s)`, want: []rune("sss")},
		// {input: `bytes(second)`, want: nil}, // cannot convert time.Duration(1s) to []uint8
		{input: `int64(second)`, want: int64(time.Second)},
		{input: `float64(1)`, want: float64(1)},
		{input: `duration(1)`, want: time.Duration(1)},
		{input: `isNil(err)`, want: false},
		{input: `notNil(err)`, want: true},
		{input: `isNull(err)`, want: true},
		{input: `notNull(err)`, want: false},
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

func TestNil(t *testing.T) {
	s := NewScope(nil)
	s = WithDefaultScope(s)
	v, err := s.Get("true")
	t.Logf("val= %v, err=%v", v, err)
}
