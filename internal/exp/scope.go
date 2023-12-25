package exp

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
)

var ErrNotFoundInScope = fmt.Errorf("not found in scope")

// Scope 表示一个作用域
type Scope interface {
	// Get 从该作用域内获取 name 对应的值
	// 返回 ErrNotFoundInScope 表示无法从该作用域获取 name 对应的值
	Get(name string) (result any, err error)
}

var _ Scope = (*scope)(nil)

type scope struct {
	data any
}

// NewScope 从结构体/map/数组/切片新建作用域
func NewScope(data any) *scope {
	return &scope{data: data}
}

func (s *scope) Get(name string) (any, error) {
	return getValue(name, s.data)
}

func getValue(name string, from any) (value any, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("cannot get `%v` from `%T`: %v", name, from, x)
		}
	}()
	rv := reflect.ValueOf(from)
	if rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}
	switch rv.Kind() {
	case reflect.Struct:
		value := rv.FieldByName(name)
		if isZero(value) {
			return nil, fmt.Errorf("no such field `%v`: %w", name, ErrNotFoundInScope)
		}
		return value.Interface(), nil
	case reflect.Map:
		value := rv.MapIndex(reflect.ValueOf(name))
		if isZero(value) {
			return nil, fmt.Errorf("no such key `%v`: %w", name, ErrNotFoundInScope)
		}
		return value.Interface(), nil
	case reflect.Array, reflect.Slice:
		idx, err := strconv.ParseInt(name, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("expected array index, but got `%s`: %w", name, err)
		}
		if idx < 0 {
			idx = int64(rv.Len()) + idx
		}
		return rv.Index(int(idx)).Interface(), nil
	default:
	}
	return nil, fmt.Errorf("cannot get `%s` from `%T`: %w", name, from, ErrNotFoundInScope)
}

func isZero(value reflect.Value) bool {
	return value == reflect.Value{}
}

// NameOfFunction 获取函数的全限定名
func NameOfFunction(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func callFunc(rv reflect.Value, in []reflect.Value) (out []reflect.Value, err error) {
	defer func() {
		if x := recover(); x != nil {
			// reflect: Call with too many input arguments
			err = fmt.Errorf("call function failed: %v", x)
		}
	}()
	return rv.Call(in), nil
}

type combineScope struct {
	parent Scope
	scope  Scope
}

// Get implements Scope.
func (s *combineScope) Get(name string) (any, error) {
	v, e := s.scope.Get(name)
	if e != nil {
		if errors.Is(e, ErrNotFoundInScope) {
			return s.parent.Get(name)
		}
	}
	return v, e
}

// Combine 组合两个作用域 当子作用域中找不到时 从父作用域找
func Combine(s, p Scope) Scope {
	return &combineScope{
		parent: p,
		scope:  s,
	}
}
