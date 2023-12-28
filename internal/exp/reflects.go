package exp

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
)

var ErrNoSuchValue = fmt.Errorf("no such value")

// NameOfFunction 获取函数的全限定名
func NameOfFunction(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// getValue 从 from 中获取 name 对应的值
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
			return nil, fmt.Errorf("field not found `%v`: %w", name, ErrNoSuchValue)
		}
		return value.Interface(), nil
	case reflect.Map:
		value := rv.MapIndex(reflect.ValueOf(name))
		if isZero(value) {
			return nil, fmt.Errorf("key not found `%v`: %w", name, ErrNoSuchValue)
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
	return nil, fmt.Errorf("`%s` not found in `%T`: %w", name, from, ErrNoSuchValue)
}

// callFunc 反射调用函数
func callFunc(rv reflect.Value, in []reflect.Value) (out []reflect.Value, err error) {
	defer func() {
		if x := recover(); x != nil {
			// reflect: Call with too many input arguments
			err = fmt.Errorf("call function failed: %v", x)
		}
	}()
	return rv.Call(in), nil
}

// isZero 检查值是否为零值
func isZero(value reflect.Value) bool {
	return value == reflect.Value{}
}
