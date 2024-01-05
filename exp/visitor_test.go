package exp_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"code.gopub.tech/tpl/exp"
)

type Foo struct {
	Age int
}

func TestFoo(t *testing.T) {
	var foo = Foo{}
	foo.Age = 1
	pfa := &foo.Age
	*pfa = 2
	t.Logf("foo.Age=%v", foo.Age)
	fa := foo.Age
	pfa = &fa
	*pfa = 3
	t.Logf("foo.Age=%v, fa=%v", foo.Age, fa)

}

func TestNewVisitor(t *testing.T) {
	const cNum = 0
	var aNum = 0
	var intPtr = &aNum

	var foo = Foo{}
	type args struct {
		input string
		data  any
	}
	tests := [][]struct {
		name    string
		args    args
		want    any
		wantErr bool
		errMsg  string
	}{
		{ // 基本表达式: literal 字面量
			{
				name: "simple-literal-nil",
				args: args{
					input: `nil`,
				},
				want:    nil,
				wantErr: false,
			},
			{
				name: "simple-literal-int",
				args: args{
					input: `0o_1_0`,
				},
				want:    int64(8),
				wantErr: false,
			},
			{
				name: "simple-literal-imaginary",
				args: args{
					input: `1i`,
				},
				want:    1i,
				wantErr: false,
			},
			{
				name: "simple-literal-float",
				args: args{
					input: `1e-1`,
				},
				want:    1e-1,
				wantErr: false,
			},
			{
				name: "simple-literal-string",
				args: args{
					// :text="${'It\'s cool, ' + name}"
					input: `'It\'s cool'`,
				},
				want:    "It's cool",
				wantErr: false,
			},
			{
				name: "simple-literal-string",
				args: args{
					// :text="${`It's cool, ` + name}"
					input: "`It's cool`",
				},
				want:    "It's cool",
				wantErr: false,
			},
			{
				name: "simple-literal-string",
				args: args{
					// :text='${`He says "ok" ` + name}'
					input: "`He says \"ok\"`",
				},
				want:    `He says "ok"`,
				wantErr: false,
			},
			{
				name: "simple-literal-string",
				args: args{
					// :text='${"He says \"ok\", " + name}'
					input: `"He says \"ok\"."`,
				},
				want:    `He says "ok".`,
				wantErr: false,
			},
		},
		{ // 基本表达式：操作数、括号
			{
				name: "oprandName",
				args: args{
					input: `i`,
					data:  map[string]any{"i": 1},
				},
				want:    1,
				wantErr: false,
			},
			{
				name: "oprandName-notFound",
				args: args{
					input: `i`,
					data:  map[string]any{},
				},
				want:    nil,
				wantErr: true,
				errMsg:  "oprand name `i` not found: key not found `i`: no such value",
			},
			{
				name: "oprand-paren",
				args: args{
					input: `(1)`,
					data:  map[string]any{},
				},
				want:    int64(1),
				wantErr: false,
			},
		},
		{ // 基本表达式：后缀
			{
				name: "field",
				args: args{
					input: `a.b`,
					data: map[string]any{
						"a": map[string]any{
							"b": false,
						},
					},
				},
				want:    false,
				wantErr: false,
			},
			{
				name: "field-notFound",
				args: args{
					input: `a.c`,
					data: map[string]any{
						"a": Foo{Age: 1},
					},
				},
				want:    nil,
				wantErr: true,
				errMsg:  "field `c` not found: field not found `c`: no such value",
			},
			{
				name: "index",
				args: args{
					input: `a[0]`,
					data: map[string]any{
						"a": []string{
							"hello", "world",
						},
					},
				},
				want:    "hello",
				wantErr: false,
			},
			{
				name: "index-outof-range",
				args: args{
					input: `a[2]`,
					data: map[string]any{
						"a": []string{
							"hello", "world",
						},
					},
				},
				want:    nil,
				wantErr: true,
				errMsg:  "index `2` not found: cannot get `2` from `[]string`: reflect: slice index out of range",
			},
			{
				name: "index-field",
				args: args{
					input: `a["b"]`,
					data: map[string]any{
						"a": map[string]any{
							"b": false,
						},
					},
				},
				want:    false,
				wantErr: false,
			},
			{
				name: "index-field-err",
				args: args{
					input: `a[a]`,
					data: map[string]any{
						"a": map[string]any{
							"b": false,
						},
					},
				},
				want:    nil,
				wantErr: true,
				errMsg:  `expected index to be int or string, got map[string]interface {}(map[b:false])`,
			},
			{
				name: "slice00",
				args: args{
					input: `a[:]`,
					data: map[string]any{
						"a": []string{"a", "b"},
					},
				},
				want:    []string{"a", "b"},
				wantErr: false,
			},
			{
				name: "slice01",
				args: args{
					input: `a[:1]`,
					data: map[string]any{
						"a": []string{"a", "b"},
					},
				},
				want:    []string{"a"},
				wantErr: false,
			},
			{
				name: "slice10",
				args: args{
					input: `a[1:]`,
					data: map[string]any{
						"a": []string{"a", "b"},
					},
				},
				want:    []string{"b"},
				wantErr: false,
			},
			{
				name: "slice11",
				args: args{
					input: `a[0:1]`,
					data: map[string]any{
						"a": []string{"a", "b"},
					},
				},
				want:    []string{"a"},
				wantErr: false,
			},
			{
				name: "slice3",
				args: args{
					input: `a[0:1:2]`,
					data: map[string]any{
						"a": []string{"a", "b"},
					},
				},
				want:    []string{"a"},
				wantErr: false,
			},
			{
				name: "slice30",
				args: args{
					input: `a[:1:2]`,
					data: map[string]any{
						"a": []string{"a", "b"},
					},
				},
				want:    []string{"a"},
				wantErr: false,
			},
			{
				name: "call-invalid",
				args: args{
					input: `f()`,
					data: map[string]any{
						"f": true,
					},
				},
				want:    nil,
				wantErr: true,
				errMsg:  `function call only works on function, got bool(true)`,
			},
			{
				name: "call-invalid-out",
				args: args{
					input: `f()`,
					data: map[string]any{
						"f": func() {},
					},
				},
				want:    nil,
				wantErr: true,
				errMsg:  ` error: should return 1 or 2 values, got 0`,
			},
			{
				name: "call-too-many-in",
				args: args{
					input: `f(1)`,
					data: map[string]any{
						"f": func() {},
					},
				},
				want:    nil,
				wantErr: true,
				errMsg:  `error: call function failed: reflect: Call with too many input arguments`,
			},
			{
				name: "call-out-1",
				args: args{
					input: `f()`,
					data: map[string]any{
						"f": func() int { return 0 },
					},
				},
				want:    0,
				wantErr: false,
			},
			{
				name: "call-out-2",
				args: args{
					input: `f()`,
					data: map[string]any{
						"f": func() (int, error) { return 0, nil },
					},
				},
				want:    0,
				wantErr: false,
			},
			{
				name: "call-out-2-noterr",
				args: args{
					input: `f()`,
					data: map[string]any{
						"f": func() (int, bool) { return 0, false },
					},
				},
				want:    nil,
				wantErr: true,
				errMsg:  `: the second returned value should be error type, got bool(false)`,
			},
			{
				name: "call-err",
				args: args{
					input: `f()`,
					data: map[string]any{
						"f": func() (int, error) { return 0, fmt.Errorf("---error---") },
					},
				},
				want:    nil,
				wantErr: true,
				errMsg:  `function call returned error: ---error---`,
			},
		},
		{ // 函数调用
			{
				name: "call...",
				args: args{
					input: `f("Hello, %v, %v",args...)`,
					data: map[string]any{
						"f": func(tpl string, args ...any) string {
							return fmt.Sprintf(tpl, args...)
						},
						"args": []any{"Tom", 42},
					},
				},
				want: "Hello, Tom, 42",
			},
			{
				name: "call...err",
				args: args{
					input: `f("Hello, %v, %v",args...)`,
					data: map[string]any{
						"f": func(tpl string, args ...any) string {
							return fmt.Sprintf(tpl, args...)
						},
						"args": [...]any{"Tom", 42},
					},
				},
				want:    nil,
				wantErr: true,
				errMsg:  "only slice... can as variadic argument, got array",
			},
		},
		{ // 一元表达式
			{
				name: "unary+i",
				args: args{
					input: `+1`,
				},
				want:    int64(1),
				wantErr: false,
			},
			{
				name: "unary+f",
				args: args{
					input: `+1.0`,
				},
				want:    float64(+1.0),
				wantErr: false,
			},
			{
				name: "unary+j",
				args: args{
					input: `+1.0i`,
				},
				want:    complex128(+1.0i),
				wantErr: false,
			},
			{
				name: "unary+e",
				args: args{
					input: `+false`,
				},
				want:    nil,
				wantErr: true,
				errMsg:  "unary operator `+` only supports numbers, got bool(false)",
			},
			{
				name: "unary-i",
				args: args{
					input: `-1`,
				},
				want:    int64(-1),
				wantErr: false,
			},
			{
				name: "unary-f",
				args: args{
					input: `-1.0`,
				},
				want:    float64(-1.0),
				wantErr: false,
			},
			{
				name: "unary-j",
				args: args{
					input: `-0i`,
				},
				want:    complex128(-0i),
				wantErr: false,
			},
			{
				name: "unary-e",
				args: args{
					input: `-true`,
				},
				want:    nil,
				wantErr: true,
				errMsg:  "unary operator `-` only supports numbers, got bool(true)",
			},
			{
				name: "unary!",
				args: args{
					input: `!true`,
				},
				want:    false,
				wantErr: false,
			},
			{
				name: "unary!e",
				args: args{
					input: `!0`,
				},
				want:    nil,
				wantErr: true,
				errMsg:  "unary operator `!` only supports bool, got int64(0)",
			},
			{
				name: "unary^",
				args: args{
					input: `^2`,
				},
				want:    int64(^2),
				wantErr: false,
			},
			{
				name: "unary^e",
				args: args{
					input: `^false`,
				},
				want:    nil,
				wantErr: true,
				errMsg:  "unary operator `^` only supports int, got bool(false)",
			},
			{
				name: "unary*",
				args: args{
					input: `*a`,
					data: map[string]any{
						"a": intPtr,
					},
				},
				want:    *intPtr,
				wantErr: false,
			},
			{
				name: "unary*e",
				args: args{
					input: `*a`,
					data: map[string]any{
						"a": 0,
					},
				},
				want:    nil,
				wantErr: true,
				errMsg:  "unary operator `*` only supports pointer, got int(0)",
			},
			{
				name: "unary&", // todo
				args: args{
					input: `&foo.Age`,
					data: map[string]any{
						"foo": foo,
						"f": func(foo Foo) *int {
							return &foo.Age
						},
					},
				},
				want:    nil,
				wantErr: true, // unsupported
				errMsg:  "unary oprator `&` is unsupported yet",
			},
			{
				name: "unary&e",
				args: args{
					input: `&a`,
					data: map[string]any{
						"a": cNum,
					},
				},
				want:    nil,
				wantErr: true,
				errMsg:  "unary oprator `&` is unsupported yet",
			},
			{
				name: "unary<-",
				args: args{
					input: `<-a`,
					data: map[string]any{
						"a": func() chan int {
							ch := make(chan int)
							go func() {
								ch <- 1
							}()
							return ch
						}(),
					},
				},
				want:    1,
				wantErr: false,
			},
			{
				name: "unary<-",
				args: args{
					input: `<-a`,
					data: map[string]any{
						"a": 0,
					},
				},
				want:    nil,
				wantErr: true,
				errMsg:  "unary operator `<-` only supports channel, got int(0)",
			},
		},
		{ // mul
			{
				name: "mul-i*i",
				args: args{
					input: `1*2`,
				},
				want:    int64(2),
				wantErr: false,
			},
			{
				name: "mul-i*f",
				args: args{
					input: `1*2.0`,
				},
				want:    float64(2.0),
				wantErr: false,
			},
			{
				name: "mul-i*j",
				args: args{
					input: `1*2.0i`,
				},
				want:    complex128(1 * 2i),
				wantErr: false,
			},
			{
				name: "mul-i*e",
				args: args{
					input: `1*true`,
				},
				want:    nil,
				wantErr: true,
				errMsg:  "binary operator `*` only supports numbers, right expression is bool(true)",
			},
			{
				name: "mul-f*i",
				args: args{
					input: `1.0*2`,
				},
				want:    float64(2.0),
				wantErr: false,
			},
			{
				name: "mul-f*f",
				args: args{
					input: `1.0*2.0`,
				},
				want:    float64(2.0),
				wantErr: false,
			},
			{
				name: "mul-f*j",
				args: args{
					input: `1.0*2.0i`,
				},
				want:    complex128(1 * 2i),
				wantErr: false,
			},
			{
				name: "mul-f*e",
				args: args{
					input: `1.0*nil`,
				},
				want:    nil,
				wantErr: true,
				errMsg:  "binary operator `*` only supports numbers, right expression is <nil>(<nil>)",
			},
			{
				name: "mul-j*i",
				args: args{
					input: `2i*2`,
				},
				want:    complex128(2i * 2),
				wantErr: false,
			},
			{
				name: "mul-j*f",
				args: args{
					input: `2i*2.5`,
				},
				want:    complex128(2i * 2.5),
				wantErr: false,
			},
			{
				name: "mul-j*j",
				args: args{
					input: `2i*2.5i`,
				},
				want:    complex128(2i * 2.5i),
				wantErr: false,
			},
			{
				name: "mul-j*e",
				args: args{
					input: `2i*nil`,
				},
				want:    nil,
				wantErr: true,
				errMsg:  "binary operator `*` only supports numbers, right expression is <nil>(<nil>)",
			},
			{
				name: "mul-e*f",
				args: args{
					input: `""*2.0`,
				},
				want:    nil,
				wantErr: true,
				errMsg:  "binary operator `*` only supports numbers, left expression is string()",
			},
			{
				name: "mul-i/i",
				args: args{
					input: `4/3`,
				},
				want:    int64(4 / 3),
				wantErr: false,
			},
			{
				name: "mul-i%i",
				args: args{
					input: `4%3`,
				},
				want:    int64(4 % 3), // =1
				wantErr: false,
			},
			{
				name: "mul-i<<i",
				args: args{
					input: `0b_11<<1`,
				},
				want:    int64(0b_110),
				wantErr: false,
			},
			{
				name: "mul-i>>i",
				args: args{
					input: `0b_11>>1`,
				},
				want:    int64(0b_01),
				wantErr: false,
			},
			{
				name: "mul-i&i",
				args: args{
					input: `0b_11&0b_10`,
				},
				want:    int64(0b_10),
				wantErr: false,
			},
			{
				name: "mul-i&^i",
				args: args{
					input: `0b_111&^0b_101`, // bit clear
				},
				want: int64(0b_010), // 对第一个操作数做运算，
				// 将其对应 第二个操作数 是 1 的那些位清空为 0，
				// 是 0 的那些位保留
				wantErr: false,
			},
			{
				name: "add-s+s",
				args: args{
					input: `'4'+"2"`,
				},
				want:    "42",
				wantErr: false,
			},
			{
				name: "add-i+s",
				args: args{
					input: `4+"2"`,
				},
				want:    "42",
				wantErr: false,
			},
			{
				name: "add-i+i",
				args: args{
					input: `4+2`,
				},
				want:    int64(6),
				wantErr: false,
			},
			{
				name: "add-i-i",
				args: args{
					input: `4-2`,
				},
				want:    int64(2),
				wantErr: false,
			},
			{
				name: "add-i|i",
				args: args{
					input: `0b_11|0b_10`,
				},
				want:    int64(0b_11),
				wantErr: false,
			},
			{
				name: "add-i^i",
				args: args{
					input: `0b_11^0b_01`, // xor
				},
				want:    int64(0b_10), // 对应位不同置为 1
				wantErr: false,
			},
			{
				name: "rel-i==i",
				args: args{
					input: `1==1`,
				},
				want:    true,
				wantErr: false,
			},
			{
				name: "rel-i!=i",
				args: args{
					input: `1!=0`,
				},
				want:    true,
				wantErr: false,
			},
			{
				name: "rel-i>i",
				args: args{
					input: `1>0`,
				},
				want:    true,
				wantErr: false,
			},
			{
				name: "rel-i>=f",
				args: args{
					input: `1>=0.1`,
				},
				want:    true,
				wantErr: false,
			},
			{
				name: "rel-f<i",
				args: args{
					input: `1.5<1`,
				},
				want:    false,
				wantErr: false,
			},
			{
				name: "rel-f<=f",
				args: args{
					input: `1.5<=1.5`,
				},
				want:    true,
				wantErr: false,
			},
			{
				name: "rel-s<=s",
				args: args{
					input: `'abc'<'abcd'`,
				},
				want:    true,
				wantErr: false,
			},
			{
				name: "rel-s<i",
				args: args{
					input: `'abc'<0`,
				},
				want:    nil,
				wantErr: true,
				errMsg:  "binary operator `<` only supports int/float/string, got string(abc) < int64(0)",
			},
			{
				name: "logic-f&&n",
				args: args{
					input: `false&&nil`,
				},
				want:    false,
				wantErr: false,
			},
			{
				name: "logic-t&&n",
				args: args{
					input: `true&&nil`,
				},
				want:    nil,
				wantErr: true,
				errMsg:  "logic binary operator `&&` only supports bool, got bool(true) && <nil>(<nil>)",
			},
			{
				name: "logic-t&&t",
				args: args{
					input: `true&&1==1`,
				},
				want:    true,
				wantErr: false,
			},
			{
				name: "logic-t||n",
				args: args{
					input: `1==1||nil`,
				},
				want:    true,
				wantErr: false,
			},
			{
				name: "logic-f||f",
				args: args{
					input: `1==2||1>0`,
				},
				want:    true,
				wantErr: false,
			},
			{
				name: "cond?true:f",
				args: args{
					input: `1==1?1:0`,
				},
				want:    int64(1),
				wantErr: false,
			},
			{
				name: "cond?t:false",
				args: args{
					input: `1!=1?1:0`,
				},
				want:    int64(0),
				wantErr: false,
			},
			{
				name: "err?t:f",
				args: args{
					input: `nil?1:0`,
				},
				want:    nil,
				wantErr: true,
				errMsg:  "condition is not bool: <nil>(<nil>)",
			},
		},
	}
	for _, tts := range tests {
		for _, tt := range tts {
			t.Run(tt.name, func(t *testing.T) {
				tree, err := exp.ParseCode(tt.args.input)
				t.Logf("tree: %v, err=%v", exp.TreesToString(tree), err)
				if err != nil {
					t.Errorf("ParseCode() error = %v", err)
					return
				}
				v := exp.NewVisitor(exp.NewPos(1, 1), exp.NewScope(tt.args.data))
				result, err := v.Evaluate(tree)
				t.Logf("result=%T(%v), err=%v", result, result, err)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetResult error = %v, wantErr %v", err, tt.wantErr)
				}
				if !reflect.DeepEqual(result, tt.want) {
					t.Errorf("GetResult = %v, want %v", result, tt.want)
				}
				if err != nil && !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("GetResult err = %v, want contains %v", err, tt.errMsg)
				}
			})
		}
	}
}
