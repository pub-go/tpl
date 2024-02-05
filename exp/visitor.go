package exp

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"code.gopub.tech/errors"
	"code.gopub.tech/tpl/exp/parser"
	"github.com/antlr4-go/antlr/v4"
)

var _ parser.GoExpressionVisitor = (*visitor)(nil)
var errorInterface = reflect.TypeOf((*error)(nil)).Elem()

// NewVisitor 创建一个 Visitor 实例
// 通过 Visitor 遍历语法树来对其求值
func NewVisitor(pos Pos, s Scope) *visitor {
	return &visitor{
		Scope: WithDefaultScope(s),
		Pos:   pos,
	}
}

type visitor struct {
	*parser.BaseGoExpressionVisitor
	Scope
	Pos
	error
}

// Evaluate 对语法树求值
func (v *visitor) Evaluate(tree antlr.ParseTree) (result any, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = errors.Errorf("recovered from panic: %v", x)
		}
	}()
	return tree.Accept(v), v.GetError()
}

// GetError 返回求值过程的错误信息
func (v *visitor) GetError() error {
	return v.error
}

func (v *visitor) SetError(ctx antlr.ParserRuleContext, format string, args ...any) any {
	start := ctx.GetStart()
	stop := ctx.GetStop()
	line, column := start.GetLine(), start.GetColumn()
	endLine, endColumn := stop.GetLine(), stop.GetColumn()
	if line == endLine && column == endColumn {
		v.error = errors.Errorf("evaluate code failed: ```%v```, at position %v: %w",
			ctx.GetText(), v.Pos.Add(line, column), errors.Errorf(format, args...),
		)
	} else {
		v.error = errors.Errorf("evaluate code failed: ```%v```, from position %v to %v: %w",
			ctx.GetText(), v.Pos.Add(line, column), v.Pos.Add(endLine, endColumn),
			errors.Errorf(format, args...),
		)
	}
	return nil
}

func (v *visitor) SetErrorOnToken(token antlr.Token, format string, args ...any) any {
	v.error = errors.Errorf("evaluate code failed at position %v: %w",
		v.Pos.Add(token.GetLine(), token.GetColumn()),
		errors.Errorf(format, args...),
	)
	return nil
}

func (v *visitor) VisitExpression(ctx *parser.ExpressionContext) any {
	if v.error != nil {
		return nil
	}
	switch {
	case ctx.PrimaryExpr() != nil:
		return ctx.PrimaryExpr().Accept(v)
	case ctx.GetUnary_op() != nil:
		return v.unOp(ctx.GetUnary_op(), ctx.Expression(0).Accept(v))
	case ctx.GetMul_op() != nil:
		left := ctx.Expression(0).Accept(v)
		right := ctx.Expression(1).Accept(v)
		return v.mulOp(ctx.GetMul_op(), left, right)
	case ctx.GetAdd_op() != nil:
		left := ctx.Expression(0).Accept(v)
		right := ctx.Expression(1).Accept(v)
		return v.addOp(ctx.GetAdd_op(), left, right)
	case ctx.GetRel_op() != nil:
		left := ctx.Expression(0).Accept(v)
		right := ctx.Expression(1).Accept(v)
		return v.relOp(ctx.GetRel_op(), left, right)
	case ctx.LOGICAL_AND() != nil:
		left := ctx.Expression(0).Accept(v)
		if i, ok := IsBool(left); ok && !i {
			return false // 短路 无需计算 right
		}
		right := ctx.Expression(1).Accept(v)
		return v.logOp(ctx.LOGICAL_AND().GetSymbol(), left, right)
	case ctx.LOGICAL_OR() != nil:
		left := ctx.Expression(0).Accept(v)
		if i, ok := IsBool(left); ok && i {
			return true // 短路 无需计算 right
		}
		right := ctx.Expression(1).Accept(v)
		return v.logOp(ctx.LOGICAL_OR().GetSymbol(), left, right)
	case ctx.Question() != nil:
		cond := ctx.Expression(0).Accept(v)
		if val, ok := IsBool(cond); ok {
			if val {
				return ctx.Expression(1).Accept(v)
			}
			return ctx.Expression(2).Accept(v)
		}
		return v.SetError(ctx, "condition is not bool: %T(%v)", cond, cond)
	default:
		panic("assert error")
	}
}

func (v *visitor) VisitPrimaryExpr(ctx *parser.PrimaryExprContext) any {
	if v.error != nil {
		return nil
	}
	switch {
	case ctx.Operand() != nil:
		return ctx.Operand().Accept(v)
	case ctx.PrimaryExpr() != nil:
		primaryValue := ctx.PrimaryExpr().Accept(v)
		switch {
		case ctx.Field() != nil:
			// struct?.field
			// struct.field
			// map.key
			field := ctx.Field().IDENTIFIER()
			val, err := getValue(field.GetText(), primaryValue)
			if err != nil {
				return v.SetError(ctx, "field `%v` not found: %w", field.GetText(), err)
			}
			return val
		case ctx.Index() != nil:
			// slice[0]
			// struct["field"]
			// map["key"]
			var name string
			index := ctx.Index().Expression().Accept(v)
			switch iv := index.(type) {
			case int64:
				name = strconv.FormatInt(iv, 10)
			case string:
				name = iv
			default:
				return v.SetError(ctx, "expected index to be int or string, got %T(%v)",
					index, index)
			}
			val, err := getValue(name, primaryValue)
			if err != nil {
				return v.SetError(ctx, "index `%v` not found: %w", name, err)
			}
			return val
		case ctx.Slice() != nil:
			rv := reflect.ValueOf(primaryValue)
			rk := rv.Kind()
			if rk != reflect.Array && rk != reflect.Slice {
				return v.SetError(ctx, "slice operator only works on array or slice, got %T(%v)",
					primaryValue, primaryValue)
			}
			start := 0
			end := rv.Len()
			// arr[?:?]
			// arr[?:_:_]
			sctx := ctx.Slice()
			var getIntValue = func(e parser.IExpressionContext) (int64, error) {
				val := e.Accept(v)
				if i, ok := val.(int64); ok {
					return i, nil
				} else {
					return 0, errors.Errorf("expected integer, got %T(%v)", val, val)
				}
			}
			if len(sctx.AllCOLON()) == 1 {
				if e := sctx.GetLo(); e != nil {
					if i, err := getIntValue(e); err != nil {
						return v.SetError(ctx, "invalid start index: %w", err)
					} else {
						start = int(i)
					}
				}
				if e := sctx.GetHi(); e != nil {
					if i, err := getIntValue(e); err != nil {
						return v.SetError(ctx, "invalid end index: %w", err)
					} else {
						end = int(i)
					}
				}
				rv = rv.Slice(start, end)
				return rv.Interface()
			} else {
				if e := sctx.GetLo(); e != nil {
					if i, err := getIntValue(e); err != nil {
						return v.SetError(ctx, "invalid start index: %w", err)
					} else {
						start = int(i)
					}
				}
				if i, err := getIntValue(sctx.GetHi()); err != nil {
					return v.SetError(ctx, "invalid end index: %w", err)
				} else {
					end = int(i)
				}
				var max int
				if i, err := getIntValue(sctx.GetCap_()); err != nil {
					return v.SetError(ctx, "invalid cap index: %w", err)
				} else {
					max = int(i)
				}
				rv = rv.Slice3(start, end, max)
				return rv.Interface()
			}
		case ctx.Arguments() != nil:
			rv := reflect.ValueOf(primaryValue)
			rk := rv.Kind()
			if rk != reflect.Func {
				return v.SetError(ctx, "function call only works on function, got %T(%v)",
					primaryValue, primaryValue)
			}
			var args any = []reflect.Value{}
			argCtx := ctx.Arguments().ExpressionList()
			if argCtx != nil {
				args = argCtx.Accept(v)
				if v.error != nil {
					return nil
				}
			}
			if in, ok := args.([]reflect.Value); ok {
				if argCtx != nil {
					if ctx.Arguments().ELLIPSIS() != nil {
						last := in[len(in)-1]
						if last.Kind() != reflect.Slice {
							return v.SetError(ctx, "only slice... can as variadic argument, got %v", last.Kind())
						}
						in = in[:len(in)-1]
						count := last.Len()
						for i := 0; i < count; i++ {
							in = append(in, last.Index(i))
						}
					}
				}
				out, err := callFunc(rv, in)
				if err != nil {
					return v.SetError(ctx, "call function `%v` error: %w",
						NameOfFunction(primaryValue), err)
				}
				rt := rv.Type()
				switch rt.NumOut() {
				case 1:
					return out[0].Interface()
				case 2:
					v2 := out[1].Interface()

					if rt.Out(1).Implements(errorInterface) {
						if e2 := out[1].Interface(); e2 != nil {
							return v.SetError(ctx, "function call returned error: %w",
								e2.(error))
						}
						return out[0].Interface()
					} else {
						return v.SetError(ctx, "error to call function `%v`: "+
							"the second returned value should be error type, got %T(%v)",
							NameOfFunction(primaryValue), v2, v2)
					}
				default:
					return v.SetError(ctx, "call function `%v` error: "+
						"should return 1 or 2 values, got %d",
						NameOfFunction(primaryValue), len(out))
				}
			}
			panic("[VisitPrimaryExpr] assert error when get function arguments")
		default:
			panic("[VisitPrimaryExpr] assert error when visiting primaryExpr suffix")
		}
	default:
		panic("[VisitPrimaryExpr] assert error")
	}
}

func (v *visitor) VisitOperand(ctx *parser.OperandContext) any {
	if v.error != nil {
		return nil
	}
	switch {
	case ctx.Literal() != nil:
		return ctx.Literal().Accept(v)
	case ctx.OperandName() != nil:
		name := ctx.GetText()
		val, err := v.Get(name)
		if err != nil {
			return v.SetError(ctx, "oprand name `%v` not found: %w", name, err)
		}
		return val
	case ctx.L_PAREN() != nil:
		return ctx.Expression().Accept(v)
	default:
		panic("assert error")
	}
}

func (v *visitor) VisitLiteral(ctx *parser.LiteralContext) any {
	if v.error != nil {
		return nil
	}
	switch {
	case ctx.LiteralNil() != nil:
		return nil
	case ctx.Integer() != nil:
		text := ctx.GetText()
		// base=0=按 text 中的前缀自行判断进制
		// 0b    base=2
		// 0o, 0 base=8
		// 0x    base=16
		// else  base=10
		val, err := strconv.ParseInt(text, 0, 64)
		if err != nil {
			panic(err)
		}
		return val
	case ctx.String_() != nil:
		text := ctx.GetText()
		if text[0] == '\'' {
			// Unquote 内部假定单引号括起的都是单个字符
			// 我们允许单引号括起字符串 这里兼容一下 转为双引号
			text = strings.ReplaceAll(text, "\\'", "'")
			text = `"` + text[1:len(text)-1] + `"`
		}
		t, err := strconv.Unquote(text) // 去除引号
		if err != nil {
			panic(errors.Errorf("unquote error(%s): %w", text, err))
		}
		return t
	case ctx.LiteralFloat() != nil:
		text := ctx.GetText()
		val, err := strconv.ParseFloat(text, 64)
		if err != nil {
			panic(err)
		}
		return val
	case ctx.LiteralImag() != nil:
		text := ctx.GetText()
		text = text[:len(text)-1] // remove i
		if strings.Contains(text, ".") {
			val, err := strconv.ParseFloat(text, 64)
			if err != nil {
				panic(err)
			}
			return complex(0, val)
		}
		val, err := strconv.ParseInt(text, 0, 64)
		if err != nil {
			panic(err)
		}
		return complex(0, float64(val))
	default:
		panic("assert error")
	}
}

func (v *visitor) VisitExpressionList(ctx *parser.ExpressionListContext) any {
	if v.error != nil {
		return nil
	}
	var args []reflect.Value
	for _, e := range ctx.AllExpression() {
		args = append(args, reflect.ValueOf(e.Accept(v)))
	}
	return args
}

// unOp 一元表达式
// + - ! ^ * & <-
func (v *visitor) unOp(token antlr.Token, value any) any {
	op := token.GetText()
	switch op {
	case "+":
		if i, ok := IsInt(value); ok {
			return +i
		}
		if i, ok := IsFloat(value); ok {
			return +i
		}
		if i, ok := IsComplex(value); ok {
			return +i
		}
		return v.SetErrorOnToken(token, "unary operator `+` only supports numbers, got %T(%v)", value, value)
	case "-":
		if i, ok := IsInt(value); ok {
			return -i
		}
		if i, ok := IsFloat(value); ok {
			return -i
		}
		if i, ok := IsComplex(value); ok {
			return -i
		}
		return v.SetErrorOnToken(token, "unary operator `-` only supports numbers, got %T(%v)", value, value)
	case "!":
		if i, ok := IsBool(value); ok {
			return !i
		}
		return v.SetErrorOnToken(token, "unary operator `!` only supports bool, got %T(%v)", value, value)
	case "^":
		if i, ok := IsInt(value); ok {
			return ^i
		}
		return v.SetErrorOnToken(token, "unary operator `^` only supports int, got %T(%v)", value, value)
	case "*":
		rv := reflect.ValueOf(value)
		if rv.Kind() == reflect.Pointer {
			return rv.Elem().Interface()
		}
		return v.SetErrorOnToken(token, "unary operator `*` only supports pointer, got %T(%v)", value, value)
	case "&":
		panic("unary oprator `&` is unsupported yet")
	case "<-":
		rv := reflect.ValueOf(value)
		if rv.Kind() == reflect.Chan {
			val, _ := rv.Recv()
			return val.Interface()
		}
		return v.SetErrorOnToken(token, "unary operator `<-` only supports channel, got %T(%v)", value, value)
	default:
		panic(fmt.Sprintf("assert error: unknown unary operator: %v", op))
	}
}

func IsInt(value any) (int64, bool) {
	switch i := value.(type) {
	case int:
		return int64(i), true
	case int8:
		return int64(i), true
	case int16:
		return int64(i), true
	case int32:
		return int64(i), true
	case int64:
		return int64(i), true
	case uint:
		return int64(i), true
	case uint8:
		return int64(i), true
	case uint16:
		return int64(i), true
	case uint32:
		return int64(i), true
	case uint64:
		return int64(i), true
	}
	return 0, false
}

func IsFloat(value any) (float64, bool) {
	switch i := value.(type) {
	case float32:
		return float64(i), true
	case float64:
		return i, true
	}
	return 0, false
}

func IsComplex(value any) (complex128, bool) {
	switch i := value.(type) {
	case complex64:
		return complex128(i), true
	case complex128:
		return i, true
	}
	return 0, false
}

func IsString(input any) (string, bool) {
	switch i := input.(type) {
	case string:
		return i, true
	}
	return "", false
}

func IsBool(input any) (bool, bool) {
	switch i := input.(type) {
	case bool:
		return i, true
	}
	return false, false
}

/*
[算数运算符](https://go.dev/ref/spec#Arithmetic_operators)
+    sum                    integers, floats, complex values, strings
-    difference             integers, floats, complex values
*    product                integers, floats, complex values
/    quotient               integers, floats, complex values
%    remainder              integers

&    bitwise AND            integers
|    bitwise OR             integers
^    bitwise XOR            integers
&^   bit clear (AND NOT)    integers

<<   left shift             integer << integer >= 0
>>   right shift            integer >> integer >= 0
*/

// mulOp 二元表达式 乘法等高优运算
// * / % << >> & &^
func (v *visitor) mulOp(token antlr.Token, left, right any) any {
	op := token.GetText()
	switch op {
	case "*", "/": // int/float/complex
		return v.binaryOp3(token, left, right)
	case "%", "<<", ">>", "&", "&^": // int
		return v.binaryOpInt(token, left, right)
	default:
		panic("assert error: unknown mul operator: " + op)
	}
}

// binaryOp3  支持 int/float/complex 三种类型 的二元运算: - * /
func (v *visitor) binaryOp3(token antlr.Token, left, right any) any {
	op := token.GetText()
	if i, ok := IsInt(left); ok {
		if j, ok := IsInt(right); ok {
			return biOp3[int64](op, i, j)
		}
		if j, ok := IsFloat(right); ok {
			return biOp3[float64](op, float64(i), j)

		}
		if j, ok := IsComplex(right); ok {
			return biOp3[complex128](op, complex(float64(i), 0), j)
		}
		return v.SetErrorOnToken(token, "binary operator `%v` only supports numbers, right expression is %T(%v)",
			op, right, right)
	}
	if i, ok := IsFloat(left); ok {
		if j, ok := IsInt(right); ok {
			return biOp3[float64](op, i, float64(j))
		}
		if j, ok := IsFloat(right); ok {
			return biOp3[float64](op, i, j)
		}
		if j, ok := IsComplex(right); ok {
			return biOp3[complex128](op, complex(i, 0), j)
		}
		return v.SetErrorOnToken(token, "binary operator `%v` only supports numbers, right expression is %T(%v)",
			op, right, right)
	}
	if i, ok := IsComplex(left); ok {
		if j, ok := IsInt(right); ok {
			return biOp3[complex128](op, i, complex(float64(j), 0))
		}
		if j, ok := IsFloat(right); ok {
			return biOp3[complex128](op, i, complex(j, 0))
		}
		if j, ok := IsComplex(right); ok {
			return biOp3[complex128](op, i, j)
		}
		return v.SetErrorOnToken(token, "binary operator `%v` only supports numbers, right expression is %T(%v)",
			op, right, right)
	}
	return v.SetErrorOnToken(token, "binary operator `%v` only supports numbers, left expression is %T(%v)",
		op, left, left)
}

// biOp3 int/float/complex 都支持的运算符 + - * /
func biOp3[T int64 | float64 | complex128](op string, left, right T) T {
	switch op {
	case "*":
		return left * right
	case "/":
		return left / right
	case "+":
		return left + right
	case "-":
		return left - right
	default:
		panic("assert error: unknown binary operator: " + op)
	}
}

// binaryOpInt 只支持 int 的二元运算符号: % & | ^ &^ << >>
func (v *visitor) binaryOpInt(token antlr.Token, left, right any) any {
	op := token.GetText()
	if i, ok := IsInt(left); ok {
		if j, ok := IsInt(right); ok {
			return biOpInt(op, i, j)
		}
	}
	return v.SetErrorOnToken(token, "binary operator `%v` only supports int, got %T(%v) %v %T(%v)",
		op, left, left, op, right, right)
}

func biOpInt(op string, left, right int64) int64 {
	switch op {
	case "%":
		return left % right
	case "&":
		return left & right
	case "|":
		return left | right
	case "^":
		return left ^ right
	case "&^":
		return left &^ right
	case "<<":
		return left << right
	case ">>":
		return left >> right
	}
	return 0
}

// addOp 二元表达式 加法等中优运算
// + - | ^
func (v *visitor) addOp(token antlr.Token, left, right any) any {
	op := token.GetText()
	switch op {
	case "+":
		i, ok1 := IsString(left)
		j, ok2 := IsString(right)
		if ok1 && ok2 {
			return i + j
		}
		if ok1 || ok2 { // 和字符串拼接
			return fmt.Sprintf("%v%v", left, right)
		}
		fallthrough
	case "-":
		return v.binaryOp3(token, left, right)
	case "|", "^":
		return v.binaryOpInt(token, left, right)
	}
	return nil
}

// relOp 二元表达式 比较等低优运算
// == != < <= > >=
func (v *visitor) relOp(token antlr.Token, left, right any) any {
	op := token.GetText()
	switch op {
	case "==":
		return left == right
	case "!=":
		return left != right
	case "<", "<=", ">", ">=":
		if i, ok := IsInt(left); ok {
			if j, ok := IsInt(right); ok {
				return relOp3[int64](op, i, j)
			}
			if j, ok := IsFloat(right); ok {
				return relOp3[float64](op, float64(i), j)
			}
		}
		if i, ok := IsFloat(left); ok {
			if j, ok := IsInt(right); ok {
				return relOp3[float64](op, i, float64(j))
			}
			if j, ok := IsFloat(right); ok {
				return relOp3[float64](op, i, j)
			}
		}
		if i, ok := IsString(left); ok {
			if j, ok := IsString(right); ok {
				return relOp3[string](op, i, j)
			}
		}
		return v.SetErrorOnToken(token, "binary operator `%v` only supports int/float/string, got %T(%v) %v %T(%v)",
			op, left, left, op, right, right)
	default:
		panic("assert error: unknown relationship binary operator: " + op)
	}
}

// relOp3 支持 int/float/string 三种数据类型的 关系运算符
func relOp3[T int64 | float64 | string](op string, left, right T) bool {
	switch op {
	case ">":
		return left > right
	case ">=":
		return left >= right
	case "<":
		return left < right
	case "<=":
		return left <= right
	}
	return false
}

// logOp 二元表达式 逻辑运算
// && ||
func (v *visitor) logOp(token antlr.Token, left, right any) any {
	op := token.GetText()
	i, ok1 := IsBool(left)
	j, ok2 := IsBool(right)
	if !ok1 || !ok2 {
		return v.SetErrorOnToken(token, "logic binary operator `%v` only supports bool, got %T(%v) %v %T(%v)",
			op, left, left, op, right, right)
	}
	switch op {
	case "&&":
		return i && j
	case "||":
		return i || j
	default:
		panic("assert error: unknown logic binary operator: " + op)
	}
}
